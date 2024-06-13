package keeper_test

import (
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/ignite/modules/testutil/nullify"
	claim "github.com/ignite/modules/x/claim/types"
)

func TestAirdropSupplyGet(t *testing.T) {
	ctx, tk := createClaimKeeper(t)

	t.Run("should allow get", func(t *testing.T) {
		sampleSupply := sdk.NewCoin("foo", sdkmath.NewInt(1000))
		tk.SetAirdropSupply(ctx, sampleSupply)

		rst, found := tk.GetAirdropSupply(ctx)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&sampleSupply),
			nullify.Fill(&rst),
		)
	})
}

func TestAirdropSupplyRemove(t *testing.T) {
	ctx, tk := createClaimKeeper(t)

	t.Run("should allow remove", func(t *testing.T) {
		tk.SetAirdropSupply(ctx, sdk.NewCoin("foo", sdkmath.NewInt(1000)))
		_, found := tk.GetAirdropSupply(ctx)
		require.True(t, found)
		tk.RemoveAirdropSupply(ctx)
		_, found = tk.GetAirdropSupply(ctx)
		require.False(t, found)
	})
}

func TestKeeper_InitializeAirdropSupply(t *testing.T) {
	ctx, tk := createClaimKeeper(t)

	tests := []struct {
		name          string
		airdropSupply sdk.Coin
	}{
		{
			name:          "should allow setting airdrop supply",
			airdropSupply: sdk.NewCoin("foo", sdkmath.NewInt(10000)),
		},
		{
			name:          "should allow specifying a new token for the supply",
			airdropSupply: sdk.NewCoin("bar", sdkmath.NewInt(125000)),
		},
		{
			name:          "should allow modifying a token for the supply",
			airdropSupply: sdk.NewCoin("bar", sdkmath.NewInt(525000)),
		},
		{
			name:          "should allow setting airdrop supply to zero",
			airdropSupply: sdk.NewCoin("foo", sdkmath.ZeroInt()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tk.InitializeAirdropSupply(ctx, tt.airdropSupply)
			require.NoError(t, err)

			airdropSupply, found := tk.GetAirdropSupply(ctx)
			require.True(t, found)
			require.True(t, airdropSupply.IsEqual(tt.airdropSupply))

			moduleBalance := tk.BankKeeper.GetBalance(
				ctx,
				tk.AccountKeeper.GetModuleAddress(claim.ModuleName),
				airdropSupply.Denom,
			)
			require.True(t, moduleBalance.IsEqual(tt.airdropSupply))
		})
	}
}

func TestEndAirdrop(t *testing.T) {
	ctx, tk := createClaimKeeper(t)

	tests := []struct {
		name                     string
		airdropSupply            sdk.Coin
		decayInfo                claim.DecayInformation
		expectedSupply           sdk.Coin
		expectedCommunityPoolAmt sdk.Coin
		wantDistribute           bool
	}{
		{
			name:           "should do nothing if airdrop supply is zero",
			airdropSupply:  sdk.NewCoin("test", sdkmath.ZeroInt()),
			decayInfo:      claim.NewEnabledDecay(ctx.BlockTime(), ctx.BlockTime()),
			expectedSupply: sdk.NewCoin("test", sdkmath.ZeroInt()),
			wantDistribute: false,
		},
		{
			name:           "should do nothing if decay is disabled",
			airdropSupply:  sdk.NewCoin("test", sdkmath.NewInt(1000)),
			decayInfo:      claim.NewDisabledDecay(),
			expectedSupply: sdk.NewCoin("test", sdkmath.NewInt(1000)),
			wantDistribute: false,
		},
		{
			name:           "should do nothing if decayEnd is after current time",
			airdropSupply:  sdk.NewCoin("test", sdkmath.NewInt(1000)),
			decayInfo:      claim.NewEnabledDecay(ctx.BlockTime(), ctx.BlockTime().Add(time.Hour)),
			expectedSupply: sdk.NewCoin("test", sdkmath.NewInt(1000)),
			wantDistribute: false,
		},
		{
			name:                     "should distribute airdrop supply with valid case",
			airdropSupply:            sdk.NewCoin("test", sdkmath.NewInt(1000)),
			decayInfo:                claim.NewEnabledDecay(time.Unix(10000, 0o0), time.Unix(10000, 10)),
			expectedSupply:           sdk.NewCoin("test", sdkmath.ZeroInt()),
			expectedCommunityPoolAmt: sdk.NewCoin("test", sdkmath.NewInt(1000)),
			wantDistribute:           true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tk.InitializeAirdropSupply(ctx, tt.airdropSupply)
			require.NoError(t, err)

			params := tk.GetParams(ctx)
			params.DecayInformation = tt.decayInfo
			tk.SetParams(ctx, params)

			err = tk.EndAirdrop(ctx)
			require.NoError(t, err)
			if tt.wantDistribute {
				feePool := tk.DistrKeeper.GetFeePool(ctx)
				for _, decCoin := range feePool.CommunityPool {
					coin := sdk.NewCoin(decCoin.Denom, decCoin.Amount.TruncateInt())
					require.Equal(t, tt.expectedCommunityPoolAmt, coin)
				}
			}

			airdropSupply, found := tk.GetAirdropSupply(ctx)
			require.True(t, found)
			require.Equal(t, tt.expectedSupply, airdropSupply)
		})
	}
}
