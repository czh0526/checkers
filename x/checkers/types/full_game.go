package types

import (
	"errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/czh0526/checkers/x/checkers/rules"
)

func (g *StoredGame) GetRedAddress() (red sdk.AccAddress, err error) {
	red, errRed := sdk.AccAddressFromBech32(g.Black)
	return red, sdkerrors.Wrapf(errRed, ErrInvalidRed.Error(), g.Red)
}

func (g *StoredGame) GetBlackAddress() (black sdk.AccAddress, err error) {
	black, errBlack := sdk.AccAddressFromBech32(g.Black)
	return black, sdkerrors.Wrapf(errBlack, ErrInvalidBlack.Error(), g.Black)
}

func (g *StoredGame) ParseGame() (game *rules.Game, err error) {
	game, errGame := rules.Parse(g.Game)
	if errGame != nil {
		return nil, sdkerrors.Wrapf(errGame, ErrGameNotParseable.Error())
	}
	game.Turn = rules.StringPieces[g.Turn].Player
	if game.Turn.Color == "" {
		return nil, sdkerrors.Wrapf(errors.New(fmt.Sprintf("Turn: %s", g.Turn)), ErrGameNotParseable.Error())
	}
	return game, nil
}

func (g *StoredGame) Validate() (err error) {
	_, err = g.ParseGame()
	if err != nil {
		return err
	}
	_, err = g.ParseGame()
	if err != nil {
		return err
	}
	_, err = g.GetRedAddress()
	if err != nil {
		return err
	}
	_, err = g.GetBlackAddress()
	if err != nil {
		return err
	}
	return err
}
