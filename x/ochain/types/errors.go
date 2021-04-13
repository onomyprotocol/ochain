package types

// DONTCOVER

import (
	sdkerrors "github.com/onomyprotocol/cosmos-sdk/types/errors"
)

// x/ochain module sentinel errors
var (
	ErrSample = sdkerrors.Register(ModuleName, 1100, "sample error")
)
