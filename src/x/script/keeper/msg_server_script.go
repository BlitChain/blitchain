package keeper

import (
	"context"
	"crypto/sha256"

	"github.com/cosmos/cosmos-sdk/x/authz"

	"blit/x/script/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateScript(goCtx context.Context, msg *types.MsgCreateScript) (*types.MsgCreateScriptResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	hash := sha256.Sum256([]byte(msg.Creator + msg.Code))
	scriptAddress := sdk.AccAddress(hash[:])

	// Check if the value already exists
	_, isFound := k.GetScript(
		ctx,
		scriptAddress.String(),
	)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "This specific code has already been deployed")
	}

	// loop over all the msg_url_permissions and create new GrantAuthorization
	// objects for each of them granting the grantee the
	// permissions to the script address

	for _, msgType := range msg.MsgTypePermissions {
		authorization := authz.NewGenericAuthorization(msgType)

		if authorization == nil {
			return nil, errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized message type: %s", msgType)
		}

		grantMsg, err := authz.NewMsgGrant(scriptAddress, sdk.MustAccAddressFromBech32(msg.Grantee), authorization, nil)
		if err != nil {
			return nil, err
		}

		if k.Router == nil {
			return nil, errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "router is nil")
		}

		handler := k.Router.Handler(grantMsg)
		if handler == nil {
			return nil, sdkerrors.ErrUnknownRequest.Wrapf("unrecognized message route: %s", sdk.MsgTypeURL(msg))
		}

		_, err = handler(ctx, grantMsg)
		if err != nil {
			return nil, errorsmod.Wrapf(err, "failed to execute message; message %v", msg)
		}

	}

	var script = types.Script{
		Address: scriptAddress.String(),
		Code:    msg.Code,
		Version: 1,
	}

	k.SetScript(
		ctx,
		script,
	)
	return &types.MsgCreateScriptResponse{Address: script.Address}, nil
}

func (k msgServer) UpdateScript(goCtx context.Context, msg *types.MsgUpdateScript) (*types.MsgUpdateScriptResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	script, isFound := k.GetScript(
		ctx,
		msg.Address,
	)

	if msg.Grantee != "" && msg.Grantee != msg.Address {

		// Prevent looping
		originalGrantee := sdk.MustAccAddressFromBech32(msg.Grantee)
		msg.Grantee = ""
		execMsg := authz.NewMsgExec(
			originalGrantee,
			[]sdk.Msg{msg},
		)

		msgExecResp, err := k.AuthzKeeper.Exec(ctx, &execMsg)
		if err != nil {
			return nil, err
		}

		var execResp types.MsgUpdateScriptResponse

		data := msgExecResp.Results[0]
		k.cdc.MustUnmarshal(data, &execResp)
		return &execResp, nil

	}

	if !isFound {
		script = types.Script{
			Address: msg.Address,
			Code:    msg.Code,
			Version: 1,
		}
	} else {
		// Update fields
		script.Code = msg.Code
		script.Version++

	}

	k.SetScript(
		ctx,
		script,
	)

	return &types.MsgUpdateScriptResponse{Version: script.Version}, nil
}
