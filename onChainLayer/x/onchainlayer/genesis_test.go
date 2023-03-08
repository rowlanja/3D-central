package onchainlayer_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "onChainLayer/testutil/keeper"
	"onChainLayer/testutil/nullify"
	"onChainLayer/x/onchainlayer"
	"onChainLayer/x/onchainlayer/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		RegistrandList: []types.Registrand{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		RegistrandCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.OnchainlayerKeeper(t)
	onchainlayer.InitGenesis(ctx, *k, genesisState)
	got := onchainlayer.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.RegistrandList, got.RegistrandList)
	require.Equal(t, genesisState.RegistrandCount, got.RegistrandCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
