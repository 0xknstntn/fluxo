package keeper

import (
	"context"

	"github.com/0xknstntn/fluxo/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ValidatorRegistration(goCtx context.Context, msg *types.MsgValidatorRegistration) (*types.MsgValidatorRegistrationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	validator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	var val = types.Validator{
		Address: msg.Creator,
	}
	amount, err := sdk.ParseCoinsNormalized(msg.Amount)
	if err != nil {
		panic(err)
	}

	/*denom, err := sdk.GetBaseDenom()
	if err != nil {
		panic(err)
	}
	coinString := cast.ToString(30) + denom*/

	stake, err := sdk.ParseCoinsNormalized("30token")
	if err != nil {
		panic(err)
	}

	res := amount.IsEqual(stake)

	if res == true {
		sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, validator, types.ModuleName, amount)
		if sdkError != nil {
			return nil, sdkError
		}
		k.AppendValidator(ctx, val)
	}
	return &types.MsgValidatorRegistrationResponse{}, nil
}
