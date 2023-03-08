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

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.OnchainlayerKeeper(t)
	onchainlayer.InitGenesis(ctx, *k, genesisState)
	got := onchainlayer.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
