package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ignite/modules/pkg/errors"
	"github.com/ignite/modules/testutil/sample"
	"github.com/ignite/modules/x/claim/types"
)

func TestMsgClaim_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgClaim
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgClaim{
				Claimer: "invalid_address",
			},
			err: errors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: types.MsgClaim{
				Claimer: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
