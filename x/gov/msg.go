package gov

import (
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/migration"
)

const (
	pathCreateProposalMsg      = "gov/create"
	pathDeleteProposalMsg      = "gov/delete"
	pathVoteMsg                = "gov/vote"
	pathTallyMsg               = "gov/tally"
	pathUpdateElectorateMsg    = "gov/electorate/update"
	pathUpdateElectionRulesMsg = "gov/electionRules/update"
)

func init() {
	migration.MustRegister(1, &CreateProposalMsg{}, migration.NoModification)
	migration.MustRegister(1, &VoteMsg{}, migration.NoModification)
	migration.MustRegister(1, &TallyMsg{}, migration.NoModification)
	migration.MustRegister(1, &DeleteProposalMsg{}, migration.NoModification)
	migration.MustRegister(1, &UpdateElectionRuleMsg{}, migration.NoModification)
	migration.MustRegister(1, &UpdateElectorateMsg{}, migration.NoModification)
}

func (CreateProposalMsg) Path() string {
	return pathCreateTextProposalMsg
}

func (m CreateProposalMsg) Validate() error {
	if err := m.GetMetadata().Validate(); err != nil {
		return errors.Wrap(err, "invalid metadata")
	}

	// TODO: is this only for text proposal?
	if len(m.ElectionRuleID) == 0 {
		return errors.Wrap(errors.ErrInput, "empty election rules id")
	}

	// Really... why not a series of if statements??
	switch {
	case len(m.GetElectorateID()) == 0:
		return errors.Wrap(errors.ErrInput, "empty electorate id")
	case m.GetStartTime() == 0:
		return errors.Wrap(errors.ErrInput, "empty start time")
	case !validTitle(m.GetTitle()):
		return errors.Wrapf(errors.ErrInput, "title: %q", m.GetTitle())
	case len(m.GetDescription()) < minDescriptionLength:
		return errors.Wrapf(errors.ErrInput, "description length lower than minimum of: %d", minDescriptionLength)
	case len(m.GetDescription()) > maxDescriptionLength:
		return errors.Wrapf(errors.ErrInput, "description length exceeds: %d", maxDescriptionLength)
	}
	if err := m.GetStartTime().Validate(); err != nil {
		return errors.Wrap(err, "start time")
	}
	if m.GetAuthor() != nil {
		if err := m.GetAuthor().Validate(); err != nil {
			return errors.Wrap(err, "author")
		}
	}
	return nil
}

func (DeleteProposalMsg) Path() string {
	return pathDeleteTextProposalMsg
}

func (m DeleteProposalMsg) Validate() error {
	if err := m.Metadata.Validate(); err != nil {
		return errors.Wrap(err, "invalid metadata")
	}

	if len(m.ID) == 0 {
		return errors.Wrap(errors.ErrInput, "empty proposal id")
	}
	return nil
}

func (VoteMsg) Path() string {
	return pathVoteMsg
}

func (m VoteMsg) Validate() error {
	if err := m.Metadata.Validate(); err != nil {
		return errors.Wrap(err, "invalid metadata")
	}
	if m.Selected != VoteOption_Yes && m.Selected != VoteOption_No && m.Selected != VoteOption_Abstain {
		return errors.Wrap(errors.ErrInput, "invalid option")
	}
	if len(m.ProposalID) == 0 {
		return errors.Wrap(errors.ErrInput, "empty proposal id")
	}
	if err := m.Voter.Validate(); m.Voter != nil && err != nil {
		return errors.Wrap(err, "invalid voter")
	}
	return nil
}

func (TallyMsg) Path() string {
	return pathTallyMsg
}

func (m TallyMsg) Validate() error {
	if err := m.Metadata.Validate(); err != nil {
		return errors.Wrap(err, "invalid metadata")
	}

	if len(m.ProposalID) == 0 {
		return errors.Wrap(errors.ErrInput, "empty proposal id")
	}
	return nil
}

func (UpdateElectionRuleMsg) Path() string {
	return pathUpdateElectionRulesMsg
}

func (m UpdateElectionRuleMsg) Validate() error {
	if err := m.Metadata.Validate(); err != nil {
		return errors.Wrap(err, "invalid metadata")
	}
	switch {
	case len(m.ElectionRuleID) == 0:
		return errors.Wrap(errors.ErrEmpty, "id")
	case m.VotingPeriodHours < minVotingPeriodHours:
		return errors.Wrapf(errors.ErrInput, "min hours: %d", minVotingPeriodHours)
	case m.VotingPeriodHours > maxVotingPeriodHours:
		return errors.Wrapf(errors.ErrInput, "max hours: %d", maxVotingPeriodHours)
	}
	return m.Threshold.Validate()
}

func (UpdateElectorateMsg) Path() string {
	return pathUpdateElectorateMsg
}

func (m UpdateElectorateMsg) Validate() error {
	if err := m.Metadata.Validate(); err != nil {
		return errors.Wrap(err, "invalid metadata")
	}

	if len(m.ElectorateID) == 0 {
		return errors.Wrap(errors.ErrEmpty, "id")
	}
	return ElectorsDiff(m.DiffElectors).Validate()
}
