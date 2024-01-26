package types

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	// this line is used by starport scaffolding # 1
)

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMintCoins{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBurnCoins{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgForceTransferCoins{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetDenomMetadata{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateTask{},
		&MsgUpdateTask{},
		&MsgDeleteTask{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateTaskResult{},
		&MsgUpdateTaskResult{},
		&MsgDeleteTaskResult{},
	)
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
