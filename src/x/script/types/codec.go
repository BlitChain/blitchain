package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	// this line is used by starport scaffolding # 1
)

// RegisterLegacyAminoCodec registers concrete types on the LegacyAmino codec
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateScript{}, "script/CreateScript", nil)
	cdc.RegisterConcrete(&MsgUpdateScript{}, "script/UpdateScript", nil)
	cdc.RegisterConcrete(&MsgRun{}, "script/Run", nil)
	// this line is used by starport scaffolding # 2

	cdc.RegisterConcrete(Params{}, "blit/x/script/Params", nil)
	legacy.RegisterAminoMsg(cdc, &MsgUpdateParams{}, "blit/x/script/MsgUpdateParams")
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateScript{},
		&MsgUpdateScript{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRun{},
	)
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
