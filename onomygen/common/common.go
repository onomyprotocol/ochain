package common

import (
	"github.com/onomyprotocol/ochain/app"
	"github.com/onomyprotocol/ochain/onomygen/common/types"
)

// Aliases
type (
	Keys       types.Keys
	NodeConfig types.NodeConfig
	CLIConfig  types.CLIConfig
	Genesis    types.Genesis
)

var (
	DefaultNodeHome = app.DefaultNodeHome
	DefaultCLIHome  = app.DefaultNodeHome + "cli"
	StakeTokenDenom = types.StakeTokenDenom

	P2PPort             = 26656
	MaxNumInboundPeers  = 1000
	MaxNumOutboundPeers = 1000
	AllowDuplicateIP    = true
)
