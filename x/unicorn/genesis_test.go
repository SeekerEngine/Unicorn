package unicorn_test

import (
	"testing"

	keepertest "Unicorn/testutil/keeper"
	"Unicorn/testutil/nullify"
	"Unicorn/x/unicorn"
	"Unicorn/x/unicorn/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.UnicornKeeper(t)
	unicorn.InitGenesis(ctx, *k, genesisState)
	got := unicorn.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
