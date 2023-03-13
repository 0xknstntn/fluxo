package simulation

import (
	"math/rand"

	"github.com/0xknstntn/fluxo/x/maker/keeper"
	"github.com/0xknstntn/fluxo/x/maker/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgBurnUSFX(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgBurnUSFX{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the BurnUSFX simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "BurnUSFX simulation not implemented"), nil, nil
	}
}
