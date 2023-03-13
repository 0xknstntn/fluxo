package keeper

import (
	"context"

	"github.com/0xknstntn/fluxo/x/maker/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) BurnUSFX(goCtx context.Context, msg *types.MsgBurnUSFX) (*types.MsgBurnUSFXResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	amount, err := sdk.ParseCoinsNormalized(msg.Amount)
	if err != nil {
		panic(err)
	}

	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, amount)
	if sdkError != nil {
		return nil, sdkError
	}
	price := k.oracleKeeper.GetPrice(ctx)
	fluxoAmount := amount.QuoInt(sdk.NewIntFromUint64(price / 1000000))

	fluxo := sdk.NewCoin("fluxo", fluxoAmount.AmountOf("usfx"))
	sdkError2 := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creator, sdk.NewCoins(fluxo))
	if sdkError2 != nil {
		return nil, sdkError
	}
	return &types.MsgBurnUSFXResponse{}, nil
}
