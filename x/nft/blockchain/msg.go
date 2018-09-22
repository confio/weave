package blockchain

import (
	"regexp"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/x/nft"
)

var _ weave.Msg = (*IssueTokenMsg)(nil)

const (
	pathIssueTokenMsg = "nft/blockchain/issue"
)

var (
	//todo: revisit pattern
	isValidID = regexp.MustCompile(`^[a-zA-Z0-9_.]{4,256}$`).MatchString
)

// Path returns the routing path for this message
func (*IssueTokenMsg) Path() string {
	return pathIssueTokenMsg
}

func (i *IssueTokenMsg) Validate() error {
	if i == nil {
		return errors.ErrInternal("must not be nil")
	}
	if err := weave.Address(i.Owner).Validate(); err != nil {
		return err
	}

	if !isValidID(string(i.Id)) {
		return nft.ErrInvalidID()
	}
	if err := i.Details.Validate(); err != nil {
		return err
	}
	// TODO: impl proper approval validation
	return nil
}
