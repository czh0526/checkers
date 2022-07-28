package keeper_test

import (
	codec "github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	"testing"

	store "github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/czh0526/checkers/x/checkers/keeper"
	"github.com/czh0526/checkers/x/checkers/types"
	tmdb "github.com/tendermint/tm-db"
)

func setupKeeper(t *testing.T) (*keeper.Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	paramsSubspace := typesparams.NewSubspace(
		codec.NewProtoCodec(registry),
		types.Amino,
		storeKey,
		memStoreKey,
		"CheckersParams",
	)
	keeper := keeper.NewKeeper(
		codec.NewProtoCodec(registry),
		storeKey,
		memStoreKey,
		paramsSubspace,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())
	return keeper, ctx
}
