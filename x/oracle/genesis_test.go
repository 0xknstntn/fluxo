package oracle_test

import (
	"testing"

	keepertest "github.com/0xknstntn/fluxo/testutil/keeper"
	"github.com/0xknstntn/fluxo/testutil/nullify"
	"github.com/0xknstntn/fluxo/x/oracle"
	"github.com/0xknstntn/fluxo/x/oracle/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		RequestList: []types.Request{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		RequestCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.OracleKeeper(t)
	oracle.InitGenesis(ctx, *k, genesisState)
	got := oracle.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.RequestList, got.RequestList)
	require.Equal(t, genesisState.RequestCount, got.RequestCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
