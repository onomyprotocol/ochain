package types

import (
	"github.com/onomyprotocol/onomy-sdk/codec"
	cdctypes "github.com/onomyprotocol/onomy-sdk/codec/types"
	// this line is used by starport scaffolding # 1
	"github.com/onomyprotocol/onomy-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
