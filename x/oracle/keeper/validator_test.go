package keeper_test

import (
	"testing"

	keepertest "github.com/0xknstntn/fluxo/testutil/keeper"
	"github.com/0xknstntn/fluxo/testutil/nullify"
	"github.com/0xknstntn/fluxo/x/oracle/keeper"
	"github.com/0xknstntn/fluxo/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNValidator(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Validator {
	items := make([]types.Validator, n)
	for i := range items {
		items[i].Id = keeper.AppendValidator(ctx, items[i])
	}
	return items
}

func TestValidatorGet(t *testing.T) {
	keeper, ctx := keepertest.OracleKeeper(t)
	items := createNValidator(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetValidator(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestValidatorRemove(t *testing.T) {
	keeper, ctx := keepertest.OracleKeeper(t)
	items := createNValidator(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveValidator(ctx, item.Id)
		_, found := keeper.GetValidator(ctx, item.Id)
		require.False(t, found)
	}
}

func TestValidatorGetAll(t *testing.T) {
	keeper, ctx := keepertest.OracleKeeper(t)
	items := createNValidator(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllValidator(ctx)),
	)
}

func TestValidatorCount(t *testing.T) {
	keeper, ctx := keepertest.OracleKeeper(t)
	items := createNValidator(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetValidatorCount(ctx))
}
