package types_test

import (
	"testing"

	"github.com/ignite/modules/x/claim/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/ignite/modules/testutil/sample"
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
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: types.MsgClaim{
				Claimer: sample.Address(sample.Rand()),
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