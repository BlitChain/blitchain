package keeper

import (
	"context"

	"blit/x/storage/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/authz"
)

func (k msgServer) CreateStorage(goCtx context.Context, msg *types.MsgCreateStorage) (*types.MsgCreateStorageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Grantee != "" && msg.Grantee != msg.Address {

		// Prevent looping
		originalGrantee := sdk.MustAccAddressFromBech32(msg.Grantee)
		msg.Grantee = msg.Address
		execMsg := authz.NewMsgExec(
			originalGrantee,
			[]sdk.Msg{msg},
		)

		msgExecResp, err := k.AuthzKeeper.Exec(ctx, &execMsg)
		if err != nil {
			return nil, err
		}

		var resp types.MsgCreateStorageResponse

		data := msgExecResp.Results[0]
		k.cdc.MustUnmarshal(data, &resp)
		return &resp, nil

	}
	// Check if the value already exists
	_, isFound := k.GetStorage(
		ctx,
		msg.Address,
		msg.Index,
	)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var storage = types.Storage{
		Address: msg.Address,
		Index:   msg.Index,
		Data:    msg.Data,
	}

	k.SetStorage(
		ctx,
		storage,
	)
	return &types.MsgCreateStorageResponse{}, nil
}

func (k msgServer) UpdateStorage(goCtx context.Context, msg *types.MsgUpdateStorage) (*types.MsgUpdateStorageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Grantee != "" && msg.Grantee != msg.Address {

		// Prevent looping
		originalGrantee := sdk.MustAccAddressFromBech32(msg.Grantee)
		msg.Grantee = msg.Address
		execMsg := authz.NewMsgExec(
			originalGrantee,
			[]sdk.Msg{msg},
		)

		msgExecResp, err := k.AuthzKeeper.Exec(ctx, &execMsg)
		if err != nil {
			return nil, err
		}

		var resp types.MsgUpdateStorageResponse

		data := msgExecResp.Results[0]
		k.cdc.MustUnmarshal(data, &resp)
		return &resp, nil

	}
	// Check if the value exists
	storage, isFound := k.GetStorage(
		ctx,
		msg.Address,
		msg.Index,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	storage.Data = msg.Data

	k.SetStorage(ctx, storage)

	return &types.MsgUpdateStorageResponse{}, nil
}

func (k msgServer) DeleteStorage(goCtx context.Context, msg *types.MsgDeleteStorage) (*types.MsgDeleteStorageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Grantee != "" && msg.Grantee != msg.Address {

		// Prevent looping
		originalGrantee := sdk.MustAccAddressFromBech32(msg.Grantee)
		msg.Grantee = msg.Address
		execMsg := authz.NewMsgExec(
			originalGrantee,
			[]sdk.Msg{msg},
		)

		msgExecResp, err := k.AuthzKeeper.Exec(ctx, &execMsg)
		if err != nil {
			return nil, err
		}

		var resp types.MsgDeleteStorageResponse

		data := msgExecResp.Results[0]
		k.cdc.MustUnmarshal(data, &resp)
		return &resp, nil

	}
	// Check if the value exists
	_, isFound := k.GetStorage(
		ctx,
		msg.Address,
		msg.Index,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	k.RemoveStorage(
		ctx,
		msg.Address,
		msg.Index,
	)

	return &types.MsgDeleteStorageResponse{}, nil
}
