package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/bitbadges/badges-module/x/badges/testutil/keeper"

	"github.com/bitbadges/badges-module/x/badges/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.BadgesKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
