package keeper

import (
	"context"
	"github.com/czh0526/checkers/x/checkers/rules"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/czh0526/checkers/x/checkers/types"
)

func (k msgServer) CreateGame(goCtx context.Context, msg *types.MsgCreateGame) (*types.MsgCreateGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	nextGame, found := k.Keeper.GetNextGame(ctx)
	if !found {
		panic("NextGame not found")
	}
	newIndex := strconv.FormatUint(nextGame.IdValue, 10)
	newGame := rules.New()
	storedGame := types.StoredGame{
		Index: newIndex,
		Game:  newGame.String(),
		Turn:  rules.PieceStrings[newGame.Turn],
		Red:   msg.Red,
		Black: msg.Black,
	}
	err := storedGame.Validate()
	if err != nil {
		return nil, err
	}
	k.Keeper.SetStoredGame(ctx, storedGame)

	nextGame.IdValue++
	k.Keeper.SetNextGame(ctx, nextGame)

	return &types.MsgCreateGameResponse{
		IdValue: newIndex,
	}, nil
}
