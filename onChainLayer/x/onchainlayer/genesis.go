package onchainlayer

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"onChainLayer/x/onchainlayer/keeper"
	"onChainLayer/x/onchainlayer/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the registrand
	for _, elem := range genState.RegistrandList {
		k.SetRegistrand(ctx, elem)
	}

	// Set registrand count
	k.SetRegistrandCount(ctx, genState.RegistrandCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.RegistrandList = k.GetAllRegistrand(ctx)
	genesis.RegistrandCount = k.GetRegistrandCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
