package keeper

import (
	"github.com/onomyprotocol/ochain/x/ochain/types"
)

var _ types.QueryServer = Keeper{}
