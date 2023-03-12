package keeper

import (
	"context"

	"github.com/0xknstntn/fluxo/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateRequest(goCtx context.Context, msg *types.MsgCreateRequest) (*types.MsgCreateRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	var request = types.Request{
		Price: msg.Price,
	}
	validatorList := k.GetAllValidator(ctx)
	for i := 0; i < len(validatorList); i++ {
		validator := validatorList[i]
		valAddress, _ := sdk.AccAddressFromBech32(validator.Address)
		res := sdk.Address.Equals(valAddress, creator)
		if res == true {
			k.AppendRequest(ctx, request)
			break
		}
	}
	return &types.MsgCreateRequestResponse{}, nil
}
