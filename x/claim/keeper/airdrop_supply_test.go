package keeper_test

import (
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	tc "github.com/ignite/modules/testutil/constructor"
	testkeeper "github.com/ignite/modules/testutil/keeper"
	"github.com/ignite/modules/testutil/nullify"
	"github.com/ignite/modules/testutil/sample"
	claim "github.com/ignite/modules/x/claim/types"
)

func TestAirdropSupplyGet(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	t.Run("should allow get", func(t *testing.T) {
		sampleSupply := sample.Coin(r)
		supply := claim.AirdropSupply{Supply: sampleSupply}
		err := tk.ClaimKeeper.AirdropSupply.Set(ctx, supply)
		require.NoError(t, err)

		rst, err := tk.ClaimKeeper.AirdropSupply.Get(ctx)
		require.NoError(t, err)
		require.Equal(t,
			nullify.Fill(&supply),
			nullify.Fill(&rst),
		)
	})
}

func TestAirdropSupplyRemove(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	t.Run("should allow remove", func(t *testing.T) {
		want := claim.AirdropSupply{Supply: sample.Coin(r)}
		err := tk.ClaimKeeper.AirdropSupply.Set(ctx, want)
		require.NoError(t, err)
		got, err := tk.ClaimKeeper.AirdropSupply.Get(ctx)
		require.NoError(t, err)
		require.Equal(t, want, got)
		err = tk.ClaimKeeper.AirdropSupply.Remove(ctx)
		require.NoError(t, err)
		_, err = tk.ClaimKeeper.AirdropSupply.Get(ctx)
		require.Error(t, err)
	})
}

func TestKeeper_InitializeAirdropSupply(t *testing.T) {
	// TODO: use mock for bank module to test critical errors
	// https://github.com/ignite/modules/issues/13
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	tests := []struct {
		name          string
		airdropSupply sdk.Coin
	}{
		{
			name:          "should allow setting airdrop supply",
			airdropSupply: tc.Coin(t, "10000foo"),
		},
		{
			name:          "should allow specifying a new token for the supply",
			airdropSupply: tc.Coin(t, "125000bar"),
		},
		{
			name:          "should allow modifying a token for the supply",
			airdropSupply: tc.Coin(t, "525000bar"),
		},
		{
			name:          "should allow setting airdrop supply to zero",
			airdropSupply: sdk.NewCoin("foo", sdkmath.ZeroInt()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tk.ClaimKeeper.InitializeAirdropSupply(ctx, tt.airdropSupply)
			require.NoError(t, err)

			airdropSupply, err := tk.ClaimKeeper.AirdropSupply.Get(ctx)
			require.NoError(t, err)
			require.True(t, airdropSupply.Supply.Equal(tt.airdropSupply))

			moduleBalance := tk.BankKeeper.GetBalance(
				ctx,
				tk.AccountKeeper.GetModuleAddress(claim.ModuleName),
				airdropSupply.Supply.Denom,
			)
			require.True(t, moduleBalance.IsEqual(tt.airdropSupply))
		})
	}
}

func TestEndAirdrop(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	blockTime := sdk.UnwrapSDKContext(ctx).BlockTime()

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
			decayInfo:      claim.NewEnabledDecay(blockTime, blockTime),
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
			decayInfo:      claim.NewEnabledDecay(blockTime, blockTime.Add(time.Hour)),
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
			err := tk.ClaimKeeper.InitializeAirdropSupply(ctx, tt.airdropSupply)
			require.NoError(t, err)

			params, err := tk.ClaimKeeper.Params.Get(ctx)
			require.NoError(t, err)
			params.DecayInformation = tt.decayInfo
			err = tk.ClaimKeeper.Params.Set(ctx, params)
			require.NoError(t, err)

			err = tk.ClaimKeeper.EndAirdrop(ctx)
			require.NoError(t, err)
			if tt.wantDistribute {
				feePool, err := tk.DistrKeeper.FeePool.Get(ctx)
				require.NoError(t, err)
				for _, decCoin := range feePool.CommunityPool {
					coin := sdk.NewCoin(decCoin.Denom, decCoin.Amount.TruncateInt())
					require.Equal(t, tt.expectedCommunityPoolAmt, coin)
				}
			}

			airdropSupply, err := tk.ClaimKeeper.AirdropSupply.Get(ctx)
			require.NoError(t, err)
			expectedSupply := claim.AirdropSupply{Supply: tt.expectedSupply}
			require.Equal(t, expectedSupply, airdropSupply)
		})
	}
}
