package onchainlayer

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"onChainLayer/testutil/sample"
	onchainlayersimulation "onChainLayer/x/onchainlayer/simulation"
	"onChainLayer/x/onchainlayer/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = onchainlayersimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateRegistrand = "op_weight_msg_registrand"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateRegistrand int = 100

	opWeightMsgUpdateRegistrand = "op_weight_msg_registrand"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateRegistrand int = 100

	opWeightMsgDeleteRegistrand = "op_weight_msg_registrand"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteRegistrand int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	onchainlayerGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		RegistrandList: []types.Registrand{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		RegistrandCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&onchainlayerGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateRegistrand int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateRegistrand, &weightMsgCreateRegistrand, nil,
		func(_ *rand.Rand) {
			weightMsgCreateRegistrand = defaultWeightMsgCreateRegistrand
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateRegistrand,
		onchainlayersimulation.SimulateMsgCreateRegistrand(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateRegistrand int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateRegistrand, &weightMsgUpdateRegistrand, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateRegistrand = defaultWeightMsgUpdateRegistrand
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateRegistrand,
		onchainlayersimulation.SimulateMsgUpdateRegistrand(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteRegistrand int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteRegistrand, &weightMsgDeleteRegistrand, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteRegistrand = defaultWeightMsgDeleteRegistrand
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteRegistrand,
		onchainlayersimulation.SimulateMsgDeleteRegistrand(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
