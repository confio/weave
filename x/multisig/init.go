package multisig

import "github.com/iov-one/weave"

// Initializer fulfils the Initializer interface to load data from the genesis
// file
type Initializer struct{}

var _ weave.Initializer = (*Initializer)(nil)

// FromGenesis will parse initial account info from genesis and save it in the
// database.
func (*Initializer) FromGenesis(opts weave.Options, db weave.KVStore) error {
	var contracts []struct {
		Sigs                []weave.Address `json:"sigs"`
		ActivationThreshold int64           `json:"activation_threshold"`
		AdminThreshold      int64           `json:"admin_threshold"`
	}
	if err := opts.ReadOptions("multisig", &contracts); err != nil {
		return err
	}

	bucket := NewContractBucket()
	for _, c := range contracts {
		sigs := make([][]byte, 0, len(c.Sigs))
		for _, s := range c.Sigs {
			sigs = append(sigs, []byte(s))
		}
		contract := Contract{
			Sigs:                sigs,
			ActivationThreshold: c.ActivationThreshold,
			AdminThreshold:      c.AdminThreshold,
		}
		obj := bucket.Build(db, &contract)
		if err := bucket.Save(db, obj); err != nil {
			return err
		}
	}
	return nil
}
