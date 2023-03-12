package oracle

import (
	"github.com/0xknstntn/fluxo/x/oracle/keeper"
	"github.com/0xknstntn/fluxo/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the request
	for _, elem := range genState.RequestList {
		k.SetRequest(ctx, elem)
	}

	// Set request count
	k.SetRequestCount(ctx, genState.RequestCount)
	// Set all the validator
	for _, elem := range genState.ValidatorList {
		k.SetValidator(ctx, elem)
	}

	// Set validator count
	k.SetValidatorCount(ctx, genState.ValidatorCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.RequestList = k.GetAllRequest(ctx)
	genesis.RequestCount = k.GetRequestCount(ctx)
	genesis.ValidatorList = k.GetAllValidator(ctx)
	genesis.ValidatorCount = k.GetValidatorCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
