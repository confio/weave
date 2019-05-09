package gov

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"testing"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/migration"
	"github.com/iov-one/weave/store"
	"github.com/iov-one/weave/weavetest"
	"github.com/stretchr/testify/require"
)

func TestInitFromGenesis(t *testing.T) {
	const genesisSnippet = `
{
  "governance": {
    "electorate": [
      {
        "admin": "0000000000000000000000000000000000000000",
        "title": "first",
		"update_rule_id": 1,
        "electors": [
          {
            "weight": 10,
            "address": "1111111111111111111111111111111111111111"
          },
          {
            "weight": 11,
            "address": "2222222222222222222222222222222222222222"
          }
        ]
      },
      {
        "title": "second",
        "admin": "cond:foo/bar/0000000000000001",
		"update_rule_id": 2,
        "electors": [
          {
            "weight": 1,
            "address": "3333333333333333333333333333333333333333"
          }
        ]
      }
    ],
    "rules": [
      {
		"admin":  "cond:foo/bar/0000000000000002",
        "title": "fooo",
        "voting_period_hours": 1,
        "threshold": {
          "numerator": 2,
          "denominator": 3
        }
      },
      {
		"admin":  "4444444444444444444444444444444444444444",
        "title": "barr",
        "voting_period_hours": 2,
        "threshold": {
          "numerator": 1,
          "denominator": 2
        },
        "quorum": {
          "numerator": 2,
          "denominator": 3
        }
      }
    ]
  }
}`
	var opts weave.Options
	require.NoError(t, json.Unmarshal([]byte(genesisSnippet), &opts))

	db := store.MemStore()
	migration.MustInitPkg(db, packageName)

	// when
	var ini Initializer
	if err := ini.FromGenesis(opts, db); err != nil {
		t.Fatalf("cannot load genesis: %s", err)
	}
	// then
	// first electorate ok
	_, obj, err := NewElectorateBucket().GetLatestVersion(db, weavetest.SequenceID(1))
	if err != nil {
		t.Fatalf("unexpected result: error: %s", err)
	}
	elect, _ := asElectorate(obj)
	if exp, got := "first", elect.Title; exp != got {
		t.Errorf("expected %v but got %v", exp, got)
	}
	if exp, got := uint32(1), elect.Metadata.Schema; exp != got {
		t.Errorf("expected %v but got %v", exp, got)
	}
	if exp, got := weavetest.SequenceID(1), elect.UpdateElectionRuleID; !bytes.Equal(exp, got) {
		t.Errorf("expected %v but got %v", exp, got)
	}
	if exp, got := 2, len(elect.Electors); exp != got {
		t.Errorf("expected %v but got %v", exp, got)
	}
	if exp, got := addr("0000000000000000000000000000000000000000"), elect.Admin; !exp.Equals(got) {
		t.Errorf("expected %X but got %X", exp, got)
	}
	if exp, got := addr("1111111111111111111111111111111111111111"), elect.Electors[0].Address; !exp.Equals(got) {
		t.Errorf("expected %X but got %X", exp, got)
	}
	if exp, got := uint32(10), elect.Electors[0].Weight; exp != got {
		t.Errorf("expected %v but got %v", exp, got)
	}
	if exp, got := addr("2222222222222222222222222222222222222222"), elect.Electors[1].Address; !exp.Equals(got) {
		t.Errorf("expected %X but got %X", exp, got)
	}
	if exp, got := uint32(11), elect.Electors[1].Weight; exp != got {
		t.Errorf("expected %v but got %v", exp, got)
	}
	// second electorate ok
	_, obj, err = NewElectorateBucket().GetLatestVersion(db, weavetest.SequenceID(2))
	if err != nil || elect == nil {
		t.Fatalf("unexpected result: error: %s", err)
	}
	elect, _ = asElectorate(obj)
	if exp, got := "second", elect.Title; exp != got {
		t.Errorf("expected %v but got %v", exp, got)
	}
	if exp, got := weavetest.SequenceID(2), elect.UpdateElectionRuleID; !bytes.Equal(exp, got) {
		t.Errorf("expected %v but got %v", exp, got)
	}
	cond := weave.NewCondition("foo", "bar", weavetest.SequenceID(1)).Address()
	if exp, got := cond, elect.Admin; !exp.Equals(got) {
		t.Errorf("expected %v but got %v", exp, got)
	}

	if exp, got := 1, len(elect.Electors); exp != got {
		t.Errorf("expected %v but got %v", exp, got)
	}
	if exp, got := addr("3333333333333333333333333333333333333333"), elect.Electors[0].Address; !exp.Equals(got) {
		t.Errorf("expected %X but got %X", exp, got)
	}
	if exp, got := uint32(1), elect.Electors[0].Weight; exp != got {
		t.Errorf("expected %v but got %v", exp, got)
	}

	// and then
	// first election rule ok
	r, err := NewElectionRulesBucket().GetElectionRule(db, weavetest.SequenceID(1))
	if err != nil || r == nil {
		t.Fatalf("unexpected result: error: %s", err)
	}
	if exp, got := uint32(1), r.Metadata.Schema; exp != got {
		t.Errorf("expected %v but got %v", exp, got)
	}

	if got, exp := "fooo", r.Title; exp != got {
		t.Errorf("expected %v but got %v", exp, got)
	}
	cond = weave.NewCondition("foo", "bar", weavetest.SequenceID(2)).Address()
	if exp, got := cond, r.Admin; !exp.Equals(got) {
		t.Errorf("expected %X but got %X", exp, got)
	}

	if exp, got := uint32(1), r.VotingPeriodHours; exp != got {
		t.Errorf("expected %v but got %v", exp, got)
	}
	if exp, got := (Fraction{Numerator: 2, Denominator: 3}), r.Threshold; exp != got {
		t.Errorf("expected %v but got %v", exp, got)
	}
	if r.Quorum != nil {
		t.Errorf("expected nil but got %v", r.Quorum)
	}

	// second election rule ok
	r, err = NewElectionRulesBucket().GetElectionRule(db, weavetest.SequenceID(2))
	if err != nil || r == nil {
		t.Fatalf("unexpected result: error: %s", err)
	}
	if got, exp := "barr", r.Title; exp != got {
		t.Errorf("expected %v but got %v", exp, got)
	}
	if exp, got := addr("4444444444444444444444444444444444444444"), r.Admin; !exp.Equals(got) {
		t.Errorf("expected %X but got %X", exp, got)
	}
	if exp, got := uint32(2), r.VotingPeriodHours; exp != got {
		t.Errorf("expected %v but got %v", exp, got)
	}
	if exp, got := (Fraction{Numerator: 1, Denominator: 2}), r.Threshold; exp != got {
		t.Errorf("expected %v but got %v", exp, got)
	}
	if exp, got := (Fraction{Numerator: 2, Denominator: 3}), *r.Quorum; exp != got {
		t.Errorf("expected %#v but got %#v", exp, got)
	}
}

func addr(s string) weave.Address {
	a, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return a
}
