package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/onomyprotocol/ochain/x/ochain/types"
	"github.com/onomyprotocol/onomy-sdk/codec"
	sdk "github.com/onomyprotocol/onomy-sdk/types"
)

type (
	Keeper struct {
		cdc      codec.Marshaler
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey
	}
)

func NewKeeper(cdc codec.Marshaler, storeKey, memKey sdk.StoreKey) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
