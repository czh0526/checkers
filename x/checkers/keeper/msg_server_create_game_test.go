package keeper_test

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/czh0526/checkers/x/checkers"
	"github.com/czh0526/checkers/x/checkers/keeper"
	"github.com/czh0526/checkers/x/checkers/types"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	alice = "cosmos19whugf4xmj2cm3zzegsf63kvkkanfwyf90nd0y"
	bob   = "cosmos1cqhq9cxs7y4rhms3em0cdrpkqkvfkg6y90n8lv"
)

func setupMsgServerCreateGame(t *testing.T) (types.MsgServer, keeper.Keeper, context.Context) {
	k, ctx := setupKeeper(t)
	checkers.InitGenesis(ctx, *k, *types.DefaultGenesis())
	return keeper.NewMsgServerImpl(*k), *k, sdk.WrapSDKContext(ctx)

}

func TestCreate1GameHasSaved(t *testing.T) {
	msgServer, keeper, context := setupMsgServerCreateGame(t)
	msgServer.CreateGame(context, &types.MsgCreateGame{
		Red:   bob,
		Black: alice,
	})
	nextGame, found := keeper.GetNextGame(sdk.UnwrapSDKContext(context))
	require.True(t, found)
	require.EqualValues(t, types.NextGame{
		IdValue: 2,
	}, nextGame)

	game1, found1 := keeper.GetStoredGame(sdk.UnwrapSDKContext(context), "1")
	require.True(t, found1)
	require.EqualValues(t, types.StoredGame{
		Index: "1",
		Game:  "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:  "b",
		Red:   bob,
		Black: alice,
	}, game1)
}
