package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/0xknstntn/fluxo/testutil/keeper"
	"github.com/0xknstntn/fluxo/x/maker/keeper"
	"github.com/0xknstntn/fluxo/x/maker/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.MakerKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
