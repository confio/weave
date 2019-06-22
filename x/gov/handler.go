package gov

import (
	"fmt"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/migration"
	"github.com/iov-one/weave/orm"
	"github.com/iov-one/weave/x"
)

const (
	proposalCost           = 0
	deleteProposalCost     = 0
	voteCost               = 0
	tallyCost              = 0
	updateElectorateCost   = 0
	updateElectionRuleCost = 0
	textResolutionCost     = 0
)

const packageName = "gov"

// RegisterQuery registers governance buckets for querying.
func RegisterQuery(qr weave.QueryRouter) {
	NewElectionRulesBucket().Register("electionRules", qr)
	NewElectorateBucket().Register("electorates", qr)
	NewProposalBucket().Register("proposal", qr)
	NewVoteBucket().Register("vote", qr)
}

// RegisterRoutes registers handlers for governance message processing.
func RegisterRoutes(r weave.Registry, auth x.Authenticator, decoder OptionDecoder, executor Executor) {
	r = migration.SchemaMigratingRegistry(packageName, r)
	r.Handle(pathVoteMsg, newVoteHandler(auth))
	r.Handle(pathTallyMsg, newTallyHandler(auth, decoder, executor))
	r.Handle(pathCreateProposalMsg, newCreateProposalHandler(auth, decoder))
	r.Handle(pathDeleteProposalMsg, newDeleteProposalHandler(auth))
	r.Handle(pathUpdateElectorateMsg, newUpdateElectorateHandler(auth))
	r.Handle(pathUpdateElectionRulesMsg, newUpdateElectionRuleHandler(auth))
	// We do NOT register the TextResultionHandler here... this is only for the proposal Executor
}

// RegisterBasicProposalRouters register the routes we accept for executing governance decisions.
func RegisterBasicProposalRouters(r weave.Registry, auth x.Authenticator) {
	r = migration.SchemaMigratingRegistry(packageName, r)
	r.Handle(pathUpdateElectorateMsg, newUpdateElectorateHandler(auth))
	r.Handle(pathUpdateElectionRulesMsg, newUpdateElectionRuleHandler(auth))
	r.Handle(pathTextResolutionMsg, newCreateTextResolutionHandler(auth))
}

type VoteHandler struct {
	auth       x.Authenticator
	elecBucket *ElectorateBucket
	propBucket *ProposalBucket
	voteBucket *VoteBucket
}

func newVoteHandler(auth x.Authenticator) *VoteHandler {
	return &VoteHandler{
		auth:       auth,
		elecBucket: NewElectorateBucket(),
		propBucket: NewProposalBucket(),
		voteBucket: NewVoteBucket(),
	}
}

func (h VoteHandler) Check(ctx weave.Context, db weave.KVStore, tx weave.Tx) (*weave.CheckResult, error) {
	if _, _, _, err := h.validate(ctx, db, tx); err != nil {
		return nil, err
	}
	return &weave.CheckResult{GasAllocated: voteCost}, nil

}

func (h VoteHandler) Deliver(ctx weave.Context, db weave.KVStore, tx weave.Tx) (*weave.DeliverResult, error) {
	voteMsg, proposal, vote, err := h.validate(ctx, db, tx)
	if err != nil {
		return nil, err
	}

	oldVote, err := h.voteBucket.GetVote(db, voteMsg.ProposalID, vote.Elector.Address)
	if !errors.ErrNotFound.Is(err) { // we only need to "UndoCount" if there was a previous vote
		if err != nil {
			return nil, errors.Wrap(err, "failed to load vote")
		}
		if err := proposal.UndoCountVote(*oldVote); err != nil {
			return nil, err
		}
	}

	if err := proposal.CountVote(*vote); err != nil {
		return nil, err
	}
	if err = h.voteBucket.Save(db, h.voteBucket.Build(db, voteMsg.ProposalID, *vote)); err != nil {
		return nil, errors.Wrap(err, "failed to store vote")
	}
	if err := h.propBucket.Update(db, voteMsg.ProposalID, proposal); err != nil {
		return nil, err
	}
	return &weave.DeliverResult{}, nil
}

func (h VoteHandler) validate(ctx weave.Context, db weave.KVStore, tx weave.Tx) (*VoteMsg, *Proposal, *Vote, error) {
	var msg VoteMsg
	if err := weave.LoadMsg(tx, &msg); err != nil {
		return nil, nil, nil, errors.Wrap(err, "load msg")
	}
	proposal, err := h.propBucket.GetProposal(db, msg.ProposalID)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "failed to load proposal")
	}

	if proposal.Status != Proposal_Submitted {
		return nil, nil, nil, errors.Wrap(errors.ErrState, "not in voting period")
	}
	if !weave.InThePast(ctx, proposal.VotingStartTime.Time()) {
		return nil, nil, nil, errors.Wrap(errors.ErrState, "vote before proposal start time")
	}
	if !weave.InTheFuture(ctx, proposal.VotingEndTime.Time()) {
		return nil, nil, nil, errors.Wrap(errors.ErrState, "vote after proposal end time")
	}

	voter := msg.Voter
	if voter == nil {
		voter = x.MainSigner(ctx, h.auth).Address()
	}
	obj, err := h.elecBucket.GetVersion(db, proposal.ElectorateRef)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "failed to load electorate")
	}
	elect, err := asElectorate(obj)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "electorate")
	}
	elector, ok := elect.Elector(voter)
	if !ok {
		return nil, nil, nil, errors.Wrap(errors.ErrUnauthorized, "not in participants list")
	}
	if !h.auth.HasAddress(ctx, voter) {
		return nil, nil, nil, errors.Wrap(errors.ErrUnauthorized, "voter must sign msg")
	}
	vote := &Vote{
		Metadata: &weave.Metadata{Schema: 1},
		Elector:  *elector,
		Voted:    msg.Selected,
	}
	if err := vote.Validate(); err != nil {
		return nil, nil, nil, err
	}
	return &msg, proposal, vote, nil
}

type TallyHandler struct {
	auth       x.Authenticator
	propBucket *ProposalBucket
	elecBucket *ElectorateBucket
	decoder    OptionDecoder
	executor   Executor
}

func newTallyHandler(auth x.Authenticator, decoder OptionDecoder, executor Executor) *TallyHandler {
	return &TallyHandler{
		auth:       auth,
		propBucket: NewProposalBucket(),
		elecBucket: NewElectorateBucket(),
		decoder:    decoder,
		executor:   executor,
	}
}

func (h TallyHandler) Check(ctx weave.Context, db weave.KVStore, tx weave.Tx) (*weave.CheckResult, error) {
	if _, _, err := h.validate(ctx, db, tx); err != nil {
		return nil, err
	}
	return &weave.CheckResult{GasAllocated: tallyCost}, nil

}

func (h TallyHandler) Deliver(ctx weave.Context, db weave.KVStore, tx weave.Tx) (resOut *weave.DeliverResult, errOut error) {
	msg, proposal, err := h.validate(ctx, db, tx)
	if err != nil {
		return nil, err
	}
	common := proposal
	if common == nil {
		return nil, errors.Wrap(errors.ErrState, "missing base proposal information")
	}

	if err := common.Tally(); err != nil {
		return nil, err
	}

	// store the proposal when done processing it, via whatever path
	defer func() {
		if err := h.propBucket.Update(db, msg.ProposalID, proposal); err != nil {
			resOut = nil
			errOut = err
		}
	}()

	if common.Result != Proposal_Accepted {
		return &weave.DeliverResult{Log: "Proposal not accepted"}, nil
	}

	// we only execute the store options upon success
	// if this fails... we should still return no error, so the tally update works
	// we just return the info from the executor in logs (tags?)
	opts, err := h.decoder(proposal.RawOption)
	if err != nil {
		proposal.ExecutorResult = Proposal_Failure
		return &weave.DeliverResult{Log: "Proposal accepted: error: cannot parse raw options"}, nil
	}
	if err := opts.Validate(); err != nil {
		return &weave.DeliverResult{Log: "Proposal accepted: error: options invalid"}, nil
	}

	// we add the vote ctx here, to authenticate results in the executor
	// ensure that the gov.Authenticator is used in those Handlers
	// we also add the proposal with id that was passed that can be accessed via CtxProposal()
	voteCtx := withProposal(withElectionSuccess(ctx, common.ElectionRuleRef.ID), proposal, msg.ProposalID)
	cstore, ok := db.(weave.CacheableKVStore)
	if !ok {
		proposal.ExecutorResult = Proposal_Failure
		return &weave.DeliverResult{Log: "Proposal accepted: error: need cachable kvstore"}, nil
	}
	subDB := cstore.CacheWrap()

	res, err := h.executor(voteCtx, subDB, opts)
	if err != nil {
		subDB.Discard()
		log := fmt.Sprintf("Proposal accepted: execution error: %v", err)
		proposal.ExecutorResult = Proposal_Failure
		return &weave.DeliverResult{Log: log}, nil
	}
	if err := subDB.Write(); err != nil {
		log := fmt.Sprintf("Proposal accepted: commit error: %v", err)
		proposal.ExecutorResult = Proposal_Failure
		return &weave.DeliverResult{Log: log}, nil
	}

	proposal.ExecutorResult = Proposal_Success
	res.Log = "Proposal accepted: execution success"
	return res, nil
}

func (h TallyHandler) validate(ctx weave.Context, db weave.KVStore, tx weave.Tx) (*TallyMsg, *Proposal, error) {
	var msg TallyMsg
	if err := weave.LoadMsg(tx, &msg); err != nil {
		return nil, nil, errors.Wrap(err, "load msg")
	}
	proposal, err := h.propBucket.GetProposal(db, msg.ProposalID)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to load proposal")
	}
	common := proposal
	if common == nil {
		return nil, nil, errors.Wrap(errors.ErrState, "missing base proposal information")
	}
	if common.Status != Proposal_Submitted {
		return nil, nil, errors.Wrapf(errors.ErrState, "unexpected status: %s", common.Status.String())
	}
	if !weave.InThePast(ctx, common.VotingEndTime.Time()) {
		return nil, nil, errors.Wrap(errors.ErrState, "tally before proposal end time: block time")
	}
	return &msg, proposal, nil
}

type CreateProposalHandler struct {
	auth        x.Authenticator
	decoder     OptionDecoder
	elecBucket  *ElectorateBucket
	propBucket  *ProposalBucket
	rulesBucket *ElectionRulesBucket
}

func newCreateProposalHandler(auth x.Authenticator, decoder OptionDecoder) *CreateProposalHandler {
	return &CreateProposalHandler{
		auth:        auth,
		decoder:     decoder,
		elecBucket:  NewElectorateBucket(),
		propBucket:  NewProposalBucket(),
		rulesBucket: NewElectionRulesBucket(),
	}
}

func (h CreateProposalHandler) Check(ctx weave.Context, db weave.KVStore, tx weave.Tx) (*weave.CheckResult, error) {
	if _, _, _, err := h.validate(ctx, db, tx); err != nil {
		return nil, err
	}
	return &weave.CheckResult{GasAllocated: proposalCost}, nil

}

func (h CreateProposalHandler) Deliver(ctx weave.Context, db weave.KVStore, tx weave.Tx) (*weave.DeliverResult, error) {
	msg, rule, electorate, err := h.validate(ctx, db, tx)
	if err != nil {
		return nil, err
	}
	blockTime, err := weave.BlockTime(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "block time")
	}

	proposal := &Proposal{
		Metadata:        &weave.Metadata{Schema: 1},
		Title:           msg.Title,
		RawOption:       msg.RawOption,
		Description:     msg.Description,
		ElectionRuleRef: orm.VersionedIDRef{ID: msg.ElectionRuleID, Version: rule.Version},
		ElectorateRef:   orm.VersionedIDRef{ID: rule.ElectorateID, Version: electorate.Version},
		VotingStartTime: msg.StartTime,
		VotingEndTime:   msg.StartTime.Add(rule.VotingPeriod.Duration()),
		SubmissionTime:  weave.AsUnixTime(blockTime),
		Author:          msg.Author,
		VoteState:       NewTallyResult(rule.Quorum, rule.Threshold, electorate.TotalElectorateWeight),
		Status:          Proposal_Submitted,
		Result:          Proposal_Undefined,
		ExecutorResult:  Proposal_NotRun,
	}

	obj, err := h.propBucket.Create(db, proposal)
	if err != nil {
		return nil, errors.Wrap(err, "failed to persist proposal")
	}

	return &weave.DeliverResult{Data: obj.Key()}, nil
}

func (h CreateProposalHandler) validate(ctx weave.Context, db weave.KVStore, tx weave.Tx) (*CreateProposalMsg, *ElectionRule, *Electorate, error) {
	var msg CreateProposalMsg
	if err := weave.LoadMsg(tx, &msg); err != nil {
		return nil, nil, nil, errors.Wrap(err, "load msg")
	}

	if !weave.InTheFuture(ctx, msg.StartTime.Time()) {
		return nil, nil, nil, errors.Wrap(errors.ErrInput, "start time must be in the future")
	}
	if weave.InTheFuture(ctx, msg.StartTime.Time().Add(-maxFutureStart)) {
		return nil, nil, nil, errors.Wrapf(errors.ErrInput, "start time cam not be more than %s h in the future", maxFutureStart)
	}

	_, rObj, err := h.rulesBucket.GetLatestVersion(db, msg.ElectionRuleID)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "failed to load election rule")
	}
	rule, err := asElectionRule(rObj)
	if err != nil {
		return nil, nil, nil, err
	}

	_, obj, err := h.elecBucket.GetLatestVersion(db, rule.ElectorateID)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "failed to load electorate")
	}
	elect, err := asElectorate(obj)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "electorate")
	}

	// A proposal can be created only by an entity that belongs to the
	// electorate group. At least one signature must be present in order to
	// be authorized to create a new proposal.
	authorized := false
	for _, e := range elect.Electors {
		if h.auth.HasAddress(ctx, e.Address) {
			authorized = true
			break
		}
	}
	if !authorized {
		return nil, nil, nil, errors.Wrap(errors.ErrUnauthorized, "proposal creation must be signed by at least one of the electors")
	}

	author := msg.Author
	if author != nil {
		if !h.auth.HasAddress(ctx, author) {
			return nil, nil, nil, errors.Wrap(errors.ErrUnauthorized, "author's signature required")
		}
	} else {
		author = x.MainSigner(ctx, h.auth).Address()
	}
	msg.Author = author

	opts, err := h.decoder(msg.RawOption)
	if err != nil {
		return nil, nil, nil, errors.Wrap(errors.ErrInput, "cannot parse raw options")
	}
	if err := opts.Validate(); err != nil {
		return nil, nil, nil, errors.Wrap(err, "options invalid")
	}

	return &msg, rule, elect, nil
}

type DeleteProposalHandler struct {
	auth       x.Authenticator
	propBucket *ProposalBucket
}

func newDeleteProposalHandler(auth x.Authenticator) *DeleteProposalHandler {
	return &DeleteProposalHandler{
		auth:       auth,
		propBucket: NewProposalBucket(),
	}
}

func (h DeleteProposalHandler) validate(ctx weave.Context, db weave.KVStore, tx weave.Tx) (*DeleteProposalMsg, *Proposal, error) {
	var msg DeleteProposalMsg
	if err := weave.LoadMsg(tx, &msg); err != nil {
		return nil, nil, errors.Wrap(err, "load msg")
	}
	prop, err := h.propBucket.GetProposal(db, msg.ProposalID)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to load a proposal with id %s", msg.ProposalID)
	}

	if prop.Status == Proposal_Withdrawn {
		return nil, nil, errors.Wrap(errors.ErrState, "this proposal is already withdrawn")
	}

	if weave.InThePast(ctx, prop.VotingStartTime.Time()) {
		return nil, nil, errors.Wrap(errors.ErrImmutable, "voting has already started")
	}
	if !h.auth.HasAddress(ctx, prop.Author) {
		return nil, nil, errors.Wrap(errors.ErrUnauthorized, "only the author can delete a proposal")
	}
	return &msg, prop, nil
}

func (h DeleteProposalHandler) Check(ctx weave.Context, db weave.KVStore, tx weave.Tx) (*weave.CheckResult, error) {
	if _, _, err := h.validate(ctx, db, tx); err != nil {
		return nil, err
	}
	return &weave.CheckResult{GasAllocated: deleteProposalCost}, nil
}

func (h DeleteProposalHandler) Deliver(ctx weave.Context, db weave.KVStore, tx weave.Tx) (*weave.DeliverResult, error) {
	msg, prop, err := h.validate(ctx, db, tx)
	if err != nil {
		return nil, err
	}

	prop.Status = Proposal_Withdrawn

	if err := h.propBucket.Update(db, msg.ProposalID, prop); err != nil {
		return nil, errors.Wrap(err, "failed to persist proposal")
	}

	return &weave.DeliverResult{}, nil
}

type UpdateElectorateHandler struct {
	auth       x.Authenticator
	propBucket *ProposalBucket
	elecBucket *ElectorateBucket
}

func newUpdateElectorateHandler(auth x.Authenticator) *UpdateElectorateHandler {
	return &UpdateElectorateHandler{
		auth:       auth,
		propBucket: NewProposalBucket(),
		elecBucket: NewElectorateBucket(),
	}
}

func (h UpdateElectorateHandler) Check(ctx weave.Context, db weave.KVStore, tx weave.Tx) (*weave.CheckResult, error) {
	_, _, err := h.validate(ctx, db, tx)
	if err != nil {
		return nil, err
	}
	return &weave.CheckResult{GasAllocated: updateElectorateCost}, nil
}

func (h UpdateElectorateHandler) Deliver(ctx weave.Context, db weave.KVStore, tx weave.Tx) (*weave.DeliverResult, error) {
	msg, elect, err := h.validate(ctx, db, tx)
	if err != nil {
		return nil, err
	}
	// all good, let's update
	merger := newMerger(elect.Electors)
	_ = merger.merge(msg.DiffElectors)
	elect.Electors, elect.TotalElectorateWeight = merger.serialize()

	if _, err := h.elecBucket.Update(db, msg.ElectorateID, elect); err != nil {
		return nil, errors.Wrap(err, "failed to store update")
	}
	return &weave.DeliverResult{Data: msg.ElectorateID}, nil
}

func (h UpdateElectorateHandler) validate(ctx weave.Context, db weave.KVStore, tx weave.Tx) (*UpdateElectorateMsg, *Electorate, error) {
	var msg UpdateElectorateMsg
	if err := weave.LoadMsg(tx, &msg); err != nil {
		return nil, nil, errors.Wrap(err, "load msg")
	}
	// get latest electorate version
	_, obj, err := h.elecBucket.GetLatestVersion(db, msg.ElectorateID)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to load electorate")
	}
	elect, err := asElectorate(obj)
	if err != nil {
		return nil, nil, errors.Wrap(err, "electorate")
	}

	if !h.auth.HasAddress(ctx, elect.Admin) {
		return nil, nil, errors.ErrUnauthorized
	}
	if err := newMerger(elect.Electors).merge(msg.DiffElectors); err != nil {
		return nil, nil, err
	}
	return &msg, elect, nil
}

type UpdateElectionRuleHandler struct {
	auth       x.Authenticator
	propBucket *ProposalBucket
	ruleBucket *ElectionRulesBucket
}

func newUpdateElectionRuleHandler(auth x.Authenticator) *UpdateElectionRuleHandler {
	return &UpdateElectionRuleHandler{
		auth:       auth,
		propBucket: NewProposalBucket(),
		ruleBucket: NewElectionRulesBucket(),
	}
}

func (h UpdateElectionRuleHandler) Check(ctx weave.Context, db weave.KVStore, tx weave.Tx) (*weave.CheckResult, error) {
	_, _, err := h.validate(ctx, db, tx)
	if err != nil {
		return nil, err
	}
	return &weave.CheckResult{GasAllocated: updateElectionRuleCost}, nil
}

func (h UpdateElectionRuleHandler) Deliver(ctx weave.Context, db weave.KVStore, tx weave.Tx) (*weave.DeliverResult, error) {
	msg, rule, err := h.validate(ctx, db, tx)
	if err != nil {
		return nil, err
	}
	rule.Threshold = msg.Threshold
	rule.VotingPeriod = msg.VotingPeriod
	if _, err := h.ruleBucket.Update(db, msg.ElectionRuleID, rule); err != nil {
		return nil, errors.Wrap(err, "failed to store update")
	}
	return &weave.DeliverResult{Data: msg.ElectionRuleID}, nil
}

func (h UpdateElectionRuleHandler) validate(ctx weave.Context, db weave.KVStore, tx weave.Tx) (*UpdateElectionRuleMsg, *ElectionRule, error) {
	var msg UpdateElectionRuleMsg
	if err := weave.LoadMsg(tx, &msg); err != nil {
		return nil, nil, errors.Wrap(err, "load msg")
	}
	_, obj, err := h.ruleBucket.GetLatestVersion(db, msg.ElectionRuleID)
	if err != nil {
		return nil, nil, err
	}
	rule, err := asElectionRule(obj)
	if err != nil {
		return nil, nil, err
	}
	if !h.auth.HasAddress(ctx, rule.Admin) {
		return nil, nil, errors.ErrUnauthorized
	}
	return &msg, rule, nil
}

type createTextResolutionHandler struct {
	auth   x.Authenticator
	bucket *ResolutionBucket
}

func newCreateTextResolutionHandler(auth x.Authenticator) *createTextResolutionHandler {
	return &createTextResolutionHandler{
		auth:   auth,
		bucket: NewResolutionBucket(),
	}
}

func (h createTextResolutionHandler) Check(ctx weave.Context, db weave.KVStore, tx weave.Tx) (*weave.CheckResult, error) {
	_, err := h.validate(ctx, db, tx)
	if err != nil {
		return nil, err
	}
	return &weave.CheckResult{GasAllocated: textResolutionCost}, nil
}

func (h createTextResolutionHandler) Deliver(ctx weave.Context, db weave.KVStore, tx weave.Tx) (*weave.DeliverResult, error) {
	msg, err := h.validate(ctx, db, tx)
	if err != nil {
		return nil, err
	}

	proposal, proposalID := CtxProposal(ctx)
	if proposal == nil {
		return nil, errors.Wrap(errors.ErrNotFound, "no proposal set for passed resolution")
	}
	resolution := &Resolution{
		Metadata:      &weave.Metadata{},
		ProposalID:    proposalID,
		ElectorateRef: proposal.ElectorateRef,
		Resolution:    msg.Resolution,
	}

	// Use IDGenBucket auto-id
	obj, err := h.bucket.Create(db, resolution)
	if err != nil {
		return nil, errors.Wrap(err, "failed to persist proposal")
	}
	return &weave.DeliverResult{Data: obj.Key()}, nil
}

func (h createTextResolutionHandler) validate(ctx weave.Context, db weave.KVStore, tx weave.Tx) (*CreateTextResolutionMsg, error) {
	var msg CreateTextResolutionMsg
	if err := weave.LoadMsg(tx, &msg); err != nil {
		return nil, errors.Wrap(err, "load msg")
	}
	// No auth, this can only be executed by gov proposal, and that info is stored alongside the resolution
	return &msg, nil
}
