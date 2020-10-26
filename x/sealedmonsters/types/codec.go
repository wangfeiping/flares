package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
    cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
    sdk "github.com/cosmos/cosmos-sdk/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
    // this line is used by starport scaffolding # 2
cdc.RegisterConcrete(MsgReveal{}, "sealedmonsters/CreateReveal", nil)
cdc.RegisterConcrete(MsgSeal{}, "sealedmonsters/CreateSeal", nil)
cdc.RegisterConcrete(MsgMonster{}, "sealedmonsters/CreateMonster", nil)
} 

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
    // this line is used by starport scaffolding # 3
registry.RegisterImplementations((*sdk.Msg)(nil),
	&MsgReveal{},
)
registry.RegisterImplementations((*sdk.Msg)(nil),
	&MsgSeal{},
)
registry.RegisterImplementations((*sdk.Msg)(nil),
	&MsgMonster{},
)
}

var (
	amino = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
