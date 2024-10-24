package types_test

import (
	"testing"

	"cosmossdk.io/math"
	"github.com/cometbft/cometbft/crypto"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/ignite/modules/x/fundraising/types"
)

func TestValidate_AllowedBidder(t *testing.T) {
	testBidderAddr := sdk.AccAddress(crypto.AddressHash([]byte("TestBidder")))

	testCases := []struct {
		desc          string
		allowedBidder types.AllowedBidder
		err           error
	}{
		{
			desc:          "valid allowed bidder",
			allowedBidder: types.NewAllowedBidder(1, testBidderAddr, math.NewInt(100_000_000)),
		},
		{
			desc:          "valid allowed bidder",
			allowedBidder: types.NewAllowedBidder(1, testBidderAddr, math.NewInt(0)),
			err:           types.ErrInvalidMaxBidAmount,
		},
		{
			desc:          "valid allowed bidder",
			allowedBidder: types.NewAllowedBidder(1, testBidderAddr, math.ZeroInt()),
			err:           types.ErrInvalidMaxBidAmount,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.allowedBidder.Validate()
			if tc.err != nil {
				require.Error(t, err)
				require.ErrorIs(t, tc.err, err)
				return
			}
			require.NoError(t, err)
		})
	}
}
