package keeper

import (
	"context"

	"blit/x/script/types"

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
		ctx = ctx.WithGasMeter(sdk.NewGasMeter(10000000))
	}
	goCtx = sdk.WrapSDKContext(ctx)

	res, err := k.EvalScript(goCtx, &EvalScriptContext{
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
