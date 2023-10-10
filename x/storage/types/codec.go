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
	cdc.RegisterConcrete(&MsgCreateStorage{}, "storage/CreateStorage", nil)
	cdc.RegisterConcrete(&MsgUpdateStorage{}, "storage/UpdateStorage", nil)
	cdc.RegisterConcrete(&MsgDeleteStorage{}, "storage/DeleteStorage", nil)
	// this line is used by starport scaffolding # 2

	cdc.RegisterConcrete(Params{}, "blit/x/storage/Params", nil)
	legacy.RegisterAminoMsg(cdc, &MsgUpdateParams{}, "blit/x/storage/MsgUpdateParams")
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateStorage{},
		&MsgUpdateStorage{},
		&MsgDeleteStorage{},
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
