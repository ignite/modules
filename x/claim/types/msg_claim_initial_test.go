package types_test

import (
	"testing"

	"github.com/ignite/modules/testutil/sample"

	"github.com/stretchr/testify/require"

	"github.com/ignite/modules/pkg/errors"
	"github.com/ignite/modules/x/claim/types"
)

func TestMsgClaimInitial_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgClaimInitial
		err  error
	}{
		{
			name: "should validate valid claimer address",
			msg: types.MsgClaimInitial{
				Claimer: sample.Address(r),
			},
		},
		{
			name: "should prevent validate invalid claimer address",
			msg: types.MsgClaimInitial{
				Claimer: "invalid_address",
			},
			err: errors.ErrInvalidAddress,
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
