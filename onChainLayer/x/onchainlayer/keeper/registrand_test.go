package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "onChainLayer/testutil/keeper"
	"onChainLayer/testutil/nullify"
	"onChainLayer/x/onchainlayer/keeper"
	"onChainLayer/x/onchainlayer/types"
)

func createNRegistrand(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Registrand {
	items := make([]types.Registrand, n)
	for i := range items {
		items[i].Id = keeper.AppendRegistrand(ctx, items[i])
	}
	return items
}

func TestRegistrandGet(t *testing.T) {
	keeper, ctx := keepertest.OnchainlayerKeeper(t)
	items := createNRegistrand(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetRegistrand(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestRegistrandRemove(t *testing.T) {
	keeper, ctx := keepertest.OnchainlayerKeeper(t)
	items := createNRegistrand(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveRegistrand(ctx, item.Id)
		_, found := keeper.GetRegistrand(ctx, item.Id)
		require.False(t, found)
	}
}

func TestRegistrandGetAll(t *testing.T) {
	keeper, ctx := keepertest.OnchainlayerKeeper(t)
	items := createNRegistrand(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllRegistrand(ctx)),
	)
}

func TestRegistrandCount(t *testing.T) {
	keeper, ctx := keepertest.OnchainlayerKeeper(t)
	items := createNRegistrand(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetRegistrandCount(ctx))
}
