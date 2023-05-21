package poa_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "onChainLayer/testutil/keeper"
	"onChainLayer/testutil/nullify"
	"onChainLayer/x/poa"
	"onChainLayer/x/poa/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PoaKeeper(t)
	poa.InitGenesis(ctx, *k, genesisState)
	got := poa.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
