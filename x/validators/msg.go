package validators

import (
	"strings"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
	abci "github.com/tendermint/tendermint/abci/types"
)

// Ensure we implement the Msg interface
var _ weave.Msg = (*SetValidatorsMsg)(nil)

const pathUpdate = "validators/update"

// Path returns the routing path for this message
func (*SetValidatorsMsg) Path() string {
	return pathUpdate
}

func (m ValidatorUpdate) Validate() error {
	if len(m.Pubkey.Data) != 32 ||
		strings.ToLower(m.Pubkey.Type) != "ed25519" {
		return ErrInvalidPubKey.New(m.Pubkey.Type)
	}
	if m.Power < 0 {
		return ErrInvalidPower.Newf("%d", m.Power)
	}
	return nil
}

func (m ValidatorUpdate) AsABCI() abci.ValidatorUpdate {
	return abci.ValidatorUpdate{
		PubKey: m.Pubkey.AsABCI(),
		Power:  m.Power,
	}
}

func (m Pubkey) AsABCI() abci.PubKey {
	return abci.PubKey{
		Data: m.Data,
		Type: m.Type,
	}
}

func (m *SetValidatorsMsg) Validate() error {
	if len(m.ValidatorUpdates) == 0 {
		return errors.ErrEmpty.New("validator set")
	}
	for _, v := range m.ValidatorUpdates {
		if err := v.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (m *SetValidatorsMsg) AsABCI() []abci.ValidatorUpdate {
	validators := make([]abci.ValidatorUpdate, len(m.ValidatorUpdates))
	for k, v := range m.ValidatorUpdates {
		validators[k] = v.AsABCI()
	}

	return validators
}
