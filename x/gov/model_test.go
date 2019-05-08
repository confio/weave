package gov

import (
	"math"
	"testing"
	"time"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/weavetest"
)

func TestElectorateValidation(t *testing.T) {
	specs := map[string]struct {
		Src Electorate
		Exp *errors.Error
	}{
		"All good with min electors count": {
			Src: Electorate{
				Title:                 "My Electorate",
				Admin:                 alice,
				Electors:              []Elector{{Address: alice, Weight: 1}},
				TotalElectorateWeight: 1,
				UpdateElectionRuleID:  weavetest.SequenceID(1),
			}},
		"All good with max electors count": {
			Src: Electorate{
				Title:                 "My Electorate",
				Admin:                 alice,
				Electors:              buildElectors(2000),
				TotalElectorateWeight: 2000,
				UpdateElectionRuleID:  weavetest.SequenceID(1),
			}},
		"All good with max weight count": {
			Src: Electorate{
				Title:                 "My Electorate",
				Admin:                 alice,
				Electors:              []Elector{{Address: alice, Weight: 65535}},
				TotalElectorateWeight: 65535,
				UpdateElectionRuleID:  weavetest.SequenceID(1),
			}},
		"Not enough electors": {
			Src: Electorate{
				Title:                 "My Electorate",
				Admin:                 alice,
				Electors:              []Elector{},
				TotalElectorateWeight: 1,
				UpdateElectionRuleID:  weavetest.SequenceID(1),
			},
			Exp: errors.ErrInvalidInput,
		},
		"Too many electors": {
			Src: Electorate{
				Title:                 "My Electorate",
				Admin:                 alice,
				Electors:              buildElectors(2001),
				TotalElectorateWeight: 2001,
				UpdateElectionRuleID:  weavetest.SequenceID(1),
			},
			Exp: errors.ErrInvalidInput,
		},
		"Duplicate electors": {
			Src: Electorate{
				Title:                 "My Electorate",
				Admin:                 alice,
				Electors:              []Elector{{Address: alice, Weight: 1}, {Address: alice, Weight: 1}},
				TotalElectorateWeight: 2,
				UpdateElectionRuleID:  weavetest.SequenceID(1),
			},
			Exp: errors.ErrInvalidInput,
		},
		"Empty electors weight ": {
			Src: Electorate{
				Title:                 "My Electorate",
				Admin:                 alice,
				Electors:              []Elector{{Address: bobby, Weight: 0}, {Address: alice, Weight: 1}},
				TotalElectorateWeight: 1,
				UpdateElectionRuleID:  weavetest.SequenceID(1),
			},
			Exp: errors.ErrInvalidInput,
		},
		"Electors weight exceeds max": {
			Src: Electorate{
				Title:                 "My Electorate",
				Admin:                 alice,
				Electors:              []Elector{{Address: alice, Weight: 65536}},
				TotalElectorateWeight: 65536,
				UpdateElectionRuleID:  weavetest.SequenceID(1),
			},
			Exp: errors.ErrInvalidInput,
		},
		"Electors address must not be empty": {
			Src: Electorate{
				Title:                 "My Electorate",
				Admin:                 alice,
				Electors:              []Elector{{Address: weave.Address{}, Weight: 1}},
				TotalElectorateWeight: 1,
				UpdateElectionRuleID:  weavetest.SequenceID(1),
			},
			Exp: errors.ErrEmpty,
		},
		"Total weight mismatch": {
			Src: Electorate{
				Title:                 "My Electorate",
				Admin:                 alice,
				Electors:              []Elector{{Address: alice, Weight: 1}},
				TotalElectorateWeight: 2,
				UpdateElectionRuleID:  weavetest.SequenceID(1),
			},
			Exp: errors.ErrInvalidInput,
		},
		"Title too short": {
			Src: Electorate{
				Title:                 "foo",
				Admin:                 alice,
				Electors:              []Elector{{Address: alice, Weight: 1}},
				TotalElectorateWeight: 1,
				UpdateElectionRuleID:  weavetest.SequenceID(1),
			},
			Exp: errors.ErrInvalidInput,
		},
		"Title too long": {
			Src: Electorate{
				Title:                 BigString(129),
				Admin:                 alice,
				Electors:              []Elector{{Address: alice, Weight: 1}},
				TotalElectorateWeight: 1,
				UpdateElectionRuleID:  weavetest.SequenceID(1),
			},
			Exp: errors.ErrInvalidInput,
		},
		"Admin must not be invalid": {
			Src: Electorate{
				Title:                 "My Electorate",
				Admin:                 weave.Address{0x0, 0x1, 0x2},
				Electors:              []Elector{{Address: alice, Weight: 1}},
				TotalElectorateWeight: 1,
				UpdateElectionRuleID:  weavetest.SequenceID(1),
			},
			Exp: errors.ErrInvalidInput,
		},
		"Admin must not be empty": {
			Src: Electorate{
				Title:                 "My Electorate",
				Admin:                 weave.Address{},
				Electors:              []Elector{{Address: alice, Weight: 1}},
				TotalElectorateWeight: 1,
				UpdateElectionRuleID:  weavetest.SequenceID(1),
			},
			Exp: errors.ErrEmpty,
		},
		"Update rule must not be empty": {
			Src: Electorate{
				Title:                 "My Electorate",
				Admin:                 alice,
				Electors:              []Elector{{Address: alice, Weight: 1}},
				TotalElectorateWeight: 1,
			},
			Exp: errors.ErrEmpty,
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			if exp, got := spec.Exp, spec.Src.Validate(); !exp.Is(got) {
				t.Errorf("expected %v but got %v", exp, got)
			}
		})
	}
}

func TestElectionRuleValidation(t *testing.T) {
	specs := map[string]struct {
		Src ElectionRule
		Exp *errors.Error
	}{
		"All good": {
			Src: ElectionRule{
				Title:             "My election rule",
				Admin:             alice,
				VotingPeriodHours: 1,
				Threshold:         Fraction{Numerator: 1, Denominator: 2},
			},
		},
		"Threshold fraction allowed at 0.5 ratio": {
			Src: ElectionRule{
				Title:             "My election rule",
				Admin:             alice,
				VotingPeriodHours: 1,
				Threshold:         Fraction{Numerator: 1 << 31, Denominator: math.MaxUint32},
			},
		},
		"Title too short": {
			Src: ElectionRule{
				Title:             "foo",
				Admin:             alice,
				VotingPeriodHours: 1,
				Threshold:         Fraction{Numerator: 1, Denominator: 2},
			},
			Exp: errors.ErrInvalidInput,
		},
		"Title too long": {
			Src: ElectionRule{
				Title:             BigString(129),
				Admin:             alice,
				VotingPeriodHours: 1,
				Threshold:         Fraction{Numerator: 1, Denominator: 2},
			},
			Exp: errors.ErrInvalidInput,
		},
		"Voting period empty": {
			Src: ElectionRule{
				Title:             "My election rule",
				Admin:             alice,
				VotingPeriodHours: 0,
				Threshold:         Fraction{Numerator: 1, Denominator: 2},
			},
			Exp: errors.ErrInvalidInput,
		},
		"Voting period too long": {
			Src: ElectionRule{
				Title:             "My election rule",
				Admin:             alice,
				VotingPeriodHours: 673, // = 4 * 7 * 24 + 1
				Threshold:         Fraction{Numerator: 1, Denominator: 2},
			},
			Exp: errors.ErrInvalidInput,
		},
		"Threshold must not be lower han 0.5": {
			Src: ElectionRule{
				Title:             "My election rule",
				Admin:             alice,
				VotingPeriodHours: 1,
				Threshold:         Fraction{Numerator: 1<<31 - 1, Denominator: math.MaxUint32},
			},
			Exp: errors.ErrInvalidInput,
		},
		"Threshold fraction must not be higher than 1": {
			Src: ElectionRule{
				Title:             "My election rule",
				Admin:             alice,
				VotingPeriodHours: 1,
				Threshold:         Fraction{Numerator: math.MaxUint32, Denominator: math.MaxUint32 - 1},
			},
			Exp: errors.ErrInvalidInput,
		},
		"Threshold fraction must not contain 0 numerator": {
			Src: ElectionRule{
				Title:             "My election rule",
				Admin:             alice,
				VotingPeriodHours: 1,
				Threshold:         Fraction{Numerator: 0, Denominator: math.MaxUint32 - 1},
			},
			Exp: errors.ErrInvalidInput,
		},
		"Threshold fraction must not contain 0 denominator": {
			Src: ElectionRule{
				Title:             "My election rule",
				Admin:             alice,
				VotingPeriodHours: 1,
				Threshold:         Fraction{Numerator: 1, Denominator: 0},
			},
			Exp: errors.ErrInvalidInput,
		},
		"Quorum must not be lower han 0.5": {
			Src: ElectionRule{
				Title:             "My election rule",
				Admin:             alice,
				VotingPeriodHours: 1,
				Quorum:            &Fraction{Numerator: 1<<31 - 1, Denominator: math.MaxUint32},
				Threshold:         Fraction{Numerator: 1, Denominator: 2},
			},
			Exp: errors.ErrInvalidInput,
		},
		"Quorum fraction must not be higher than 1": {
			Src: ElectionRule{
				Title:             "My election rule",
				Admin:             alice,
				VotingPeriodHours: 1,
				Quorum:            &Fraction{Numerator: math.MaxUint32, Denominator: math.MaxUint32 - 1},
				Threshold:         Fraction{Numerator: 1, Denominator: 2},
			},
			Exp: errors.ErrInvalidInput,
		},
		"Quorum fraction must not contain 0 numerator": {
			Src: ElectionRule{
				Title:             "My election rule",
				Admin:             alice,
				VotingPeriodHours: 1,
				Quorum:            &Fraction{Numerator: 0, Denominator: math.MaxUint32 - 1},
				Threshold:         Fraction{Numerator: 1, Denominator: 2},
			},
			Exp: errors.ErrInvalidInput,
		},
		"Quorum fraction must not contain 0 denominator": {
			Src: ElectionRule{
				Title:             "My election rule",
				Admin:             alice,
				VotingPeriodHours: 1,
				Quorum:            &Fraction{Numerator: 1, Denominator: 0},
				Threshold:         Fraction{Numerator: 1, Denominator: 2},
			},
			Exp: errors.ErrInvalidInput,
		},
		"Admin must not be invalid": {
			Src: ElectionRule{
				Title:             "My election rule",
				Admin:             weave.Address{0x0, 0x1, 0x2},
				VotingPeriodHours: 1,
				Quorum:            &Fraction{Numerator: 1, Denominator: 1},
				Threshold:         Fraction{Numerator: 1, Denominator: 2},
			},
			Exp: errors.ErrInvalidInput,
		},
		"Admin must not be empty": {
			Src: ElectionRule{
				Title:             "My election rule",
				Admin:             weave.Address{},
				VotingPeriodHours: 1,
				Quorum:            &Fraction{Numerator: 1, Denominator: 1},
				Threshold:         Fraction{Numerator: 1, Denominator: 2},
			},
			Exp: errors.ErrEmpty,
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			if exp, got := spec.Exp, spec.Src.Validate(); !exp.Is(got) {
				t.Errorf("expected %v but got %v", exp, got)
			}
		})
	}
}

func TestTextProposalValidation(t *testing.T) {
	specs := map[string]struct {
		Src Proposal
		Exp *errors.Error
	}{
		"Happy path": {
			Src: textProposalFixture(),
		},
		"Title too short": {
			Src: textProposalFixture(func(p *Proposal) {
				p.Title = "foo"
			}),
			Exp: errors.ErrInvalidInput,
		},
		"Title too long": {
			Src: textProposalFixture(func(p *Proposal) {
				p.Title = BigString(129)
			}),
			Exp: errors.ErrInvalidInput,
		},
		"Description empty": {
			Src: textProposalFixture(func(p *Proposal) {
				p.Description = ""
			}),
			Exp: errors.ErrInvalidInput,
		},
		"Description too long": {
			Src: textProposalFixture(func(p *Proposal) {
				p.Description = BigString(5001)
			}),
			Exp: errors.ErrInvalidInput,
		},
		"Author missing": {
			Src: textProposalFixture(func(p *Proposal) {
				p.Author = nil
			}),
			Exp: errors.ErrInvalidInput,
		},
		"ElectorateID missing": {
			Src: textProposalFixture(func(p *Proposal) {
				p.ElectorateID = nil
			}),
			Exp: errors.ErrInvalidInput,
		},
		"ElectionRuleID missing": {
			Src: textProposalFixture(func(p *Proposal) {
				p.ElectionRuleID = nil
			}),
			Exp: errors.ErrInvalidInput,
		},
		"StartTime missing": {
			Src: textProposalFixture(func(p *Proposal) {
				var unset time.Time
				p.VotingStartTime = weave.AsUnixTime(unset)
			}),
			Exp: errors.ErrInvalidInput,
		},
		"EndTime missing": {
			Src: textProposalFixture(func(p *Proposal) {
				var unset time.Time
				p.VotingEndTime = weave.AsUnixTime(unset)
			}),
			Exp: errors.ErrInvalidInput,
		},
		"Status missing": {
			Src: textProposalFixture(func(p *Proposal) {
				p.Status = Proposal_Status(0)
			}),
			Exp: errors.ErrInvalidInput,
		},
		"Result missing": {
			Src: textProposalFixture(func(p *Proposal) {
				p.Result = Proposal_Result(0)
			}),
			Exp: errors.ErrInvalidInput,
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			if exp, got := spec.Exp, spec.Src.Validate(); !exp.Is(got) {
				t.Errorf("expected %v but got %v", exp, got)
			}
		})
	}
}

func TestVoteValidate(t *testing.T) {
	specs := map[string]struct {
		Src Vote
		Exp *errors.Error
	}{
		"All good": {
			Src: Vote{
				Voted:   VoteOption_Yes,
				Elector: Elector{Address: bobby, Weight: 10},
			},
		},
		"Voted option missing": {
			Src: Vote{Elector: Elector{Address: bobby, Weight: 10}},
			Exp: errors.ErrInvalidInput,
		},
		"Elector missing": {
			Src: Vote{Voted: VoteOption_Yes},
			Exp: errors.ErrInvalidInput,
		},
		"Elector's weight missing": {
			Src: Vote{Voted: VoteOption_Yes, Elector: Elector{Address: bobby}},
			Exp: errors.ErrInvalidInput,
		},
		"Elector's Address missing": {
			Src: Vote{Voted: VoteOption_Yes, Elector: Elector{Weight: 1}},
			Exp: errors.ErrEmpty,
		},
		"Invalid option": {
			Src: Vote{Voted: VoteOption_Invalid, Elector: Elector{Address: bobby, Weight: 1}},
			Exp: errors.ErrInvalidInput,
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			if exp, got := spec.Exp, spec.Src.Validate(); !exp.Is(got) {
				t.Errorf("expected %v but got %v", exp, got)
			}
		})
	}
}

func textProposalFixture(mods ...func(*Proposal)) Proposal {
	now := weave.AsUnixTime(time.Now())
	proposal := Proposal{
		Type:            Proposal_Text,
		Title:           "My proposal",
		Description:     "My description",
		ElectionRuleID:  weavetest.SequenceID(1),
		ElectorateID:    weavetest.SequenceID(1),
		VotingStartTime: now.Add(-1 * time.Minute),
		VotingEndTime:   now.Add(time.Minute),
		SubmissionTime:  now.Add(-1 * time.Hour),
		Status:          Proposal_Submitted,
		Result:          Proposal_Undefined,
		Author:          alice,
		VoteState:       NewTallyResult(nil, Fraction{1, 2}, 11),
		Details:         &Proposal_TextDetails{&TextProposalPayload{}},
	}
	for _, mod := range mods {
		if mod != nil {
			mod(&proposal)
		}
	}
	return proposal
}
func updateElectoreateProposalFixture(mods ...func(*Proposal)) Proposal {
	now := weave.AsUnixTime(time.Now())
	proposal := Proposal{
		Type:            Proposal_UpdateElectorate,
		Title:           "My proposal",
		Description:     "My description",
		ElectionRuleID:  weavetest.SequenceID(1),
		ElectorateID:    weavetest.SequenceID(1),
		VotingStartTime: now.Add(-1 * time.Minute),
		VotingEndTime:   now.Add(time.Minute),
		SubmissionTime:  now.Add(-1 * time.Hour),
		Status:          Proposal_Submitted,
		Result:          Proposal_Undefined,
		Author:          alice,
		VoteState:       NewTallyResult(nil, Fraction{1, 2}, 11),
		Details: &Proposal_ElectorateUpdateDetails{&ElectorateUpdatePayload{
			[]Elector{{Address: alice, Weight: 10}},
		}},
	}
	for _, mod := range mods {
		if mod != nil {
			mod(&proposal)
		}
	}
	return proposal
}

func buildElectors(n int) []Elector {
	r := make([]Elector, n)
	for i := 0; i < n; i++ {
		r[i] = Elector{Weight: 1, Address: weavetest.NewCondition().Address()}
	}
	return r
}
