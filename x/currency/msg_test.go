package currency

import (
	"testing"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
)

func TestValidateCreateTokenInfoMsg(t *testing.T) {
	cases := map[string]struct {
		Msg     weave.Msg
		WantErr *errors.Error
	}{
		"valid message": {
			Msg: &CreateTokenInfoMsg{
				Metadata: &weave.Metadata{Schema: 1},
				Ticker:   "IOV",
				Name:     "mytoken",
			},
			WantErr: nil,
		},
		"missing metadata": {
			Msg: &CreateTokenInfoMsg{
				Ticker: "IOV",
				Name:   "mytoken",
			},
			WantErr: errors.ErrMetadata,
		},
	}

	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			if err := tc.Msg.Validate(); !tc.WantErr.Is(err) {
				t.Fatalf("unexpected validation error: %s", err)
			}
		})
	}

}
