package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
    cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
    sdk "github.com/cosmos/cosmos-sdk/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
    // this line is used by starport scaffolding # 2
cdc.RegisterConcrete(MsgContractTransferRecord{}, "flares/CreateContractTransferRecord", nil)
cdc.RegisterConcrete(MsgBoard{}, "flares/CreateBoard", nil)
cdc.RegisterConcrete(MsgContract{}, "flares/CreateContract", nil)
} 

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
    // this line is used by starport scaffolding # 3
registry.RegisterImplementations((*sdk.Msg)(nil),
	&MsgContractTransferRecord{},
)
registry.RegisterImplementations((*sdk.Msg)(nil),
	&MsgBoard{},
)
registry.RegisterImplementations((*sdk.Msg)(nil),
	&MsgContract{},
)
}

var (
	amino = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
