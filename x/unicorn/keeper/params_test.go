package keeper_test

import (
	"testing"

	testkeeper "Unicorn/testutil/keeper"
	"Unicorn/x/unicorn/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.UnicornKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
