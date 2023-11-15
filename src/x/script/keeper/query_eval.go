package keeper

import (
	"context"
	"encoding/json"
	"fmt"

	"blit/x/script/types"

	errorsmod "cosmossdk.io/errors"
	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Eval(goCtx context.Context, msg *types.QueryEvalRequest) (*types.QueryEvalResponse, error) {
	if msg == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx, _ := sdk.UnwrapSDKContext(goCtx).CacheContext()
	if ctx.GasMeter().Limit() < 1 || ctx.GasMeter().Limit() > 10000000 { // TODO make param
		ctx = ctx.WithGasMeter(storetypes.NewGasMeter(10000000))
	}
	goCtx = sdk.WrapSDKContext(ctx)

	attachedMsgs := []sdk.Msg{}

	if msg.AttachedMessages != "" {
		var objects []map[string]interface{}
		err := json.Unmarshal([]byte(msg.AttachedMessages), &objects)
		if err != nil {
			fmt.Println("error:", err)
			return nil, errorsmod.Wrapf(err, "failed to unmarshal attached messages: %+v", msg.AttachedMessages)
		}

		// Step 2: Iterate and marshal each object back to JSON string
		var stringifiedObjects []string
		for i, obj := range objects {
			jsonStr, err := json.Marshal(obj)
			if err != nil {
				return nil, errorsmod.Wrapf(err, "failed to marshal attached message at index %d: %s", i, obj)
			}
			stringifiedObjects = append(stringifiedObjects, string(jsonStr))
		}

		for i, anyJSON := range stringifiedObjects {
			var attachedMsg sdk.Msg
			err = k.cdc.UnmarshalInterfaceJSON([]byte(anyJSON), &attachedMsg)
			if err != nil {
				return nil, errorsmod.Wrapf(err, "failed to unmarshal attached message at index %d: %s", i, anyJSON)
			}

			attachedMsgs = append(attachedMsgs, attachedMsg)
		}
	}

	res, err := k.EvalScript(goCtx, &EvalScriptContext{
		Messages:      attachedMsgs,
		CallerAddress: msg.CallerAddress,
		ScriptAddress: msg.ScriptAddress,
		ExtraCode:     msg.ExtraCode,
		FunctionName:  msg.FunctionName,
		Kwargs:        msg.Kwargs,
	}, k.currentDepth != 0)
	if err != nil {
		return nil, err
	}
	return &types.QueryEvalResponse{Response: res.Response}, nil
}
