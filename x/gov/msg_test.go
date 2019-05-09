package gov

import (
	"testing"
	"time"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/weavetest"
)

func TestVoteMsg(t *testing.T) {
	specs := map[string]struct {
		Msg VoteMsg
		Exp *errors.Error
	}{

		"Happy path": {
			Msg: VoteMsg{ProposalID: weavetest.SequenceID(1), Selected: VoteOption_Yes, Voter: alice, Metadata: &weave.Metadata{Schema: 1}},
		},
		"Voter optional": {
			Msg: VoteMsg{ProposalID: weavetest.SequenceID(1), Selected: VoteOption_Yes, Metadata: &weave.Metadata{Schema: 1}},
		},
		"Proposal id missing": {
			Msg: VoteMsg{Selected: VoteOption_Yes, Voter: alice, Metadata: &weave.Metadata{Schema: 1}},
			Exp: errors.ErrInput,
		},
		"Vote option missing": {
			Msg: VoteMsg{ProposalID: weavetest.SequenceID(1), Voter: alice, Metadata: &weave.Metadata{Schema: 1}},
			Exp: errors.ErrInput,
		},
		"Invalid vote option": {
			Msg: VoteMsg{ProposalID: weavetest.SequenceID(1), Selected: VoteOption(100), Voter: alice, Metadata: &weave.Metadata{Schema: 1}},
			Exp: errors.ErrInput,
		},
		"Invalid voter address": {
			Msg: VoteMsg{ProposalID: weavetest.SequenceID(1), Selected: VoteOption_Yes, Voter: weave.Address([]byte{0}), Metadata: &weave.Metadata{Schema: 1}},
			Exp: errors.ErrInput,
		},
		"Metadata missing": {
			Msg: VoteMsg{ProposalID: weavetest.SequenceID(1), Selected: VoteOption_Yes, Voter: alice},
			Exp: errors.ErrMetadata,
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			err := spec.Msg.Validate()
			if !spec.Exp.Is(err) {
				t.Fatalf("check expected: %v  but got %+v", spec.Exp, err)
			}
		})
	}
}

func TestTallyMsg(t *testing.T) {
	specs := map[string]struct {
		Msg TallyMsg
		Exp *errors.Error
	}{
		"Happy path": {
			Msg: TallyMsg{ProposalID: weavetest.SequenceID(1), Metadata: &weave.Metadata{Schema: 1}},
		},
		"ID missing": {
			Msg: TallyMsg{Metadata: &weave.Metadata{Schema: 1}},
			Exp: errors.ErrInput,
		},
		"Metadata missing": {
			Msg: TallyMsg{ProposalID: weavetest.SequenceID(1)},
			Exp: errors.ErrMetadata,
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			err := spec.Msg.Validate()
			if !spec.Exp.Is(err) {
				t.Fatalf("check expected: %v  but got %+v", spec.Exp, err)
			}
		})
	}
}

func TestCrateTextProposalMsg(t *testing.T) {
	buildMsg := func(mods ...func(*CreateTextProposalMsg)) CreateTextProposalMsg {
		m := CreateTextProposalMsg{
			Metadata:       &weave.Metadata{Schema: 1},
			Title:          "any title _.-",
			Description:    "any description",
			ElectorateID:   weavetest.SequenceID(1),
			ElectionRuleID: weavetest.SequenceID(1),
			StartTime:      weave.AsUnixTime(time.Now()),
			Author:         alice,
		}
		for _, mod := range mods {
			mod(&m)
		}
		return m
	}

	specs := map[string]struct {
		Msg CreateTextProposalMsg
		Exp *errors.Error
	}{
		"Happy path": {
			Msg: buildMsg(),
		},
		"Author is optional": {
			Msg: buildMsg(func(p *CreateTextProposalMsg) {
				p.Author = nil
			}),
		},
		"Short title within range": {
			Msg: buildMsg(func(p *CreateTextProposalMsg) {
				p.Title = "fooo"
			}),
		},
		"Long title within range": {
			Msg: buildMsg(func(p *CreateTextProposalMsg) {
				p.Title = BigString(128)
			}),
		},
		"Title too short": {
			Msg: buildMsg(func(p *CreateTextProposalMsg) {
				p.Title = "foo"
			}),
			Exp: errors.ErrInput,
		},
		"Title too long": {
			Msg: buildMsg(func(p *CreateTextProposalMsg) {
				p.Title = BigString(129)
			}),
			Exp: errors.ErrInput,
		},
		"Title with invalid chars": {
			Msg: buildMsg(func(p *CreateTextProposalMsg) {
				p.Title = "title with invalid char <"
			}),
			Exp: errors.ErrInput,
		},
		"Description too short": {
			Msg: buildMsg(func(p *CreateTextProposalMsg) {
				p.Title = "foo"
			}),
			Exp: errors.ErrInput,
		},
		"Description too long": {
			Msg: buildMsg(func(p *CreateTextProposalMsg) {
				p.Title = BigString(5001)
			}),
			Exp: errors.ErrInput,
		},
		"ElectorateID missing": {
			Msg: buildMsg(func(p *CreateTextProposalMsg) {
				p.ElectorateID = nil
			}),
			Exp: errors.ErrInput,
		},
		"ElectionRuleID missing": {
			Msg: buildMsg(func(p *CreateTextProposalMsg) {
				p.ElectionRuleID = nil
			}),
			Exp: errors.ErrInput,
		},
		"StartTime zero": {
			Msg: buildMsg(func(p *CreateTextProposalMsg) {
				p.StartTime = 0
			}),
			Exp: errors.ErrInput,
		},
		"Invalid author address": {
			Msg: buildMsg(func(p *CreateTextProposalMsg) {
				p.Author = []byte{0, 0, 0, 0}
			}),
			Exp: errors.ErrInput,
		},
		"Metadata missing": {
			Msg: buildMsg(func(p *CreateTextProposalMsg) {
				p.Metadata = nil
			}),
			Exp: errors.ErrMetadata,
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			err := spec.Msg.Validate()
			if !spec.Exp.Is(err) {
				t.Fatalf("check expected: %v  but got %+v", spec.Exp, err)
			}
		})
	}
}

func TestDeleteTestProposalMsg(t *testing.T) {
	specs := map[string]struct {
		Msg DeleteProposalMsg
		Exp *errors.Error
	}{
		"Happy path": {
			Msg: DeleteProposalMsg{ID: weavetest.SequenceID(1), Metadata: &weave.Metadata{Schema: 1}},
		},
		"Empty ID": {
			Msg: DeleteProposalMsg{Metadata: &weave.Metadata{Schema: 1}},
			Exp: errors.ErrInput,
		},
		"Metadata missing": {
			Msg: DeleteProposalMsg{ID: weavetest.SequenceID(1)},
			Exp: errors.ErrMetadata,
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			err := spec.Msg.Validate()
			if !spec.Exp.Is(err) {
				t.Fatalf("check expected: %v  but got %+v", spec.Exp, err)
			}
		})
	}
}

func BigString(n int) string {
	const randomChar = "a"
	var r string
	for i := 0; i < n; i++ {
		r += randomChar
	}
	return r
}
