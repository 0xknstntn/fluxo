package keeper

import (
	"context"

	"github.com/0xknstntn/fluxo/x/maker/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) MintUSFX(goCtx context.Context, msg *types.MsgMintUSFX) (*types.MsgMintUSFXResponse, error) {
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
	//usfxAmount := amount.QuoInt(sdk.NewIntFromUint64(price))
	usfxAmount := amount.MulInt(sdk.NewIntFromUint64(price / 1000000))
	/*denom, err := sdk.GetBaseDenom()
	if err != nil {
		panic(err)
	}*/

	usfx := sdk.NewCoin("usfx", usfxAmount.AmountOf("fluxo"))
	sdkError1 := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(usfx))
	if sdkError1 != nil {
		return nil, sdkError1
	}

	sdkError2 := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creator, sdk.NewCoins(usfx))
	if sdkError2 != nil {
		return nil, sdkError
	}
	return &types.MsgMintUSFXResponse{}, nil
}
