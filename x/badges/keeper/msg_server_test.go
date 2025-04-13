package keeper_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/bitbadges/badges-module/x/badges/keeper"
	keepertest "github.com/bitbadges/badges-module/x/badges/testutil/keeper"
	"github.com/bitbadges/badges-module/x/badges/types"
)

func setupMsgServer(t testing.TB) (keeper.Keeper, types.MsgServer, context.Context) {
	k, ctx := keepertest.BadgesKeeper(t)
	return k, keeper.NewMsgServerImpl(k), ctx
}

func TestMsgServer(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)
}
