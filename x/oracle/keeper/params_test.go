package keeper_test

import (
	"testing"

	testkeeper "github.com/0xknstntn/fluxo/testutil/keeper"
	"github.com/0xknstntn/fluxo/x/oracle/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.OracleKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
