package ochain

import (
	"fmt"

	"github.com/onomyprotocol/ochain/x/ochain/keeper"
	"github.com/onomyprotocol/ochain/x/ochain/types"
	sdk "github.com/onomyprotocol/onomy-sdk/types"
	sdkerrors "github.com/onomyprotocol/onomy-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	// this line is used by starport scaffolding # handler/msgServer

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		// this line is used by starport scaffolding # 1
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
