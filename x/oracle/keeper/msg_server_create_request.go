package keeper

import (
	"context"

	"github.com/0xknstntn/fluxo/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateRequest(goCtx context.Context, msg *types.MsgCreateRequest) (*types.MsgCreateRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var request = types.Request{
		Price: msg.Price,
	}
	k.AppendRequest(ctx, request)

	return &types.MsgCreateRequestResponse{}, nil
}
