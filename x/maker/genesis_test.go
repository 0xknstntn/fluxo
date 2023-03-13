package maker_test

import (
	"testing"

	keepertest "github.com/0xknstntn/fluxo/testutil/keeper"
	"github.com/0xknstntn/fluxo/testutil/nullify"
	"github.com/0xknstntn/fluxo/x/maker"
	"github.com/0xknstntn/fluxo/x/maker/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MakerKeeper(t)
	maker.InitGenesis(ctx, *k, genesisState)
	got := maker.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	// this line is used by starport scaffolding # genesis/test/assert
}
