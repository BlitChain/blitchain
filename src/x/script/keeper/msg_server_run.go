package keeper

import (
	"context"
	"fmt"

	"blit/x/script/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/authz"
)

func (k msgServer) Run(goCtx context.Context, msg *types.MsgRun) (*types.MsgRunResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, isFound := k.GetScript(ctx, msg.ScriptAddress)

	// Allow anyone to run their own script even if it's not set yet
	if !isFound && msg.ScriptAddress != msg.CallerAddress {
		fmt.Println(fmt.Sprintf("Script at address %v not found", msg.ScriptAddress))
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("Script at address %v not set", msg.ScriptAddress))
	}

	if msg.Grantee != "" && msg.Grantee != msg.CallerAddress {

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

		var runResp types.MsgRunResponse

		data := msgExecResp.Results[0]
		k.cdc.MustUnmarshal(data, &runResp)
		return &runResp, nil

	}

	messages, err := msg.GetMessages()
	if err != nil {
		return nil, err
	}

	fmt.Println(fmt.Sprintf("Script at address %v found", msg.ScriptAddress))
	res, err := k.EvalScript(goCtx, &EvalScriptContext{
		Messages:      messages,
		CallerAddress: msg.CallerAddress,
		ScriptAddress: msg.ScriptAddress,
		ExtraCode:     msg.ExtraCode,
		FunctionName:  msg.FunctionName,
		Kwargs:        msg.Kwargs,
	}, true)
	if err != nil {
		return nil, err
	}

	k.Logger(ctx).Info(fmt.Sprintf("Run, res: %+v  err: %+v", res, err))

	var msgRunResponse types.MsgRunResponse
	msgRunResponse.Response = res.Response

	return &msgRunResponse, nil
}
