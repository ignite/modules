package keeper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/ignite/modules/x/claim/types"
)

func TestGetParams(t *testing.T) {
	testSuite := createClaimKeeper(t)
	ctx := testSuite.ctx
	tk := testSuite.tk

	t.Run("should allow params get", func(t *testing.T) {
		params := types.NewParams(types.NewEnabledDecay(
			time.Unix(1000, 0),
			time.Unix(10000, 0),
		), time.Now())
		tk.SetParams(ctx, params)
		require.EqualValues(t, params.DecayInformation, tk.GetParams(ctx).DecayInformation)
		require.Equal(t, params.AirdropStart.Unix(), tk.GetParams(ctx).AirdropStart.Unix())
	})
}
