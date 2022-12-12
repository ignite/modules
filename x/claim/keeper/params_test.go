package keeper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/ignite/modules/testutil/keeper"
	"github.com/ignite/modules/x/claim/types"
)

func TestGetParams(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	t.Run("should allow params get", func(t *testing.T) {
		params := types.NewParams(types.NewEnabledDecay(
			time.Unix(1000, 0),
			time.Unix(10000, 0),
		), 1000)
		tk.ClaimKeeper.SetParams(ctx, params)
		require.EqualValues(t, params, tk.ClaimKeeper.GetParams(ctx))
	})
}
