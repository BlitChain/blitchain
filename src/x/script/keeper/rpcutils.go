package keeper

import (
	"fmt"
	"math"
	"net/http"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type ConsumeGasRequest struct {
	Amount int `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
}

type ConsumeGasResponse struct {
	GasConsumed int64 `protobuf:"bytes,1,opt,name=GasConsumed,proto3" json:"GasConsumed,omitempty"`
	GasLimit    int64 `protobuf:"bytes,1,opt,name=GasLimit,proto3" json:"GasLimit,omitempty"`
}

type QueryRequest struct {
	Method   string `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	JsonArgs string `protobuf:"bytes,2,opt,name=jsonArgs,proto3" json:"json_args,omitempty"`
}

type MsgRequest struct {
	JsonMsg string `protobuf:"bytes,1,opt,name=jsonMsg,proto3" json:"json_msg,omitempty"`
}

func (rpcservice *RpcService) Msg(_ *http.Request, req *MsgRequest, response *string) (err error) {
	ctx := sdk.UnwrapSDKContext(rpcservice.goCtx)
	r, err := rpcservice.k.HandleMsg(ctx, req.JsonMsg, rpcservice.ScriptAddress)

	rpcservice.k.Logger(ctx).Info("Msg", "JsonMsg", req.JsonMsg, "response", r, "err", err)
	if err != nil {
		return err
	}
	*response = r
	return nil
}

// method Query calls the HandleQuery method of the keeper
func (rpcservice *RpcService) Query(_ *http.Request, req *QueryRequest, response *string) (err error) {
	ctx := sdk.UnwrapSDKContext(rpcservice.goCtx)

	rpcservice.k.Logger(ctx).Info("Query", "method", req.Method, "args", req.JsonArgs)
	r, err := rpcservice.k.HandleQuery(ctx, req.Method, req.JsonArgs)
	//rpcservice.k.Logger(ctx).Info("Query", "response", r, "error", err)
	if err != nil {
		return err
	}
	*response = r
	return nil
}

// method Consumegas consumes gas
func (rpcservice *RpcService) Consumegas(_ *http.Request, msg *ConsumeGasRequest, response *ConsumeGasResponse) (err error) {
	ctx := sdk.UnwrapSDKContext(rpcservice.goCtx)
	defer func() {
		if r := recover(); r != nil {

			fmt.Printf("Consumegas recovered: %v\n", r)
			switch r := r.(type) {
			case sdk.ErrorOutOfGas:
				err = sdkerrors.Wrapf(sdkerrors.ErrOutOfGas,
					"Consumegas script out of gas, gasWanted: %d, gasUsed: %d",
					ctx.GasMeter().Limit(), ctx.GasMeter().GasConsumed(),
				)

				response = nil
			default:
				err = r.(error)
			}
		} else {
			if ctx.GasMeter().IsPastLimit() {
				err = sdkerrors.Wrapf(sdkerrors.ErrOutOfGas,
					"PastLimit script out of gas, gasWanted: %d, gasUsed: %d",
					ctx.GasMeter().Limit(), ctx.GasMeter().GasConsumed(),
				)
				response = nil
			}
		}

		// fmt.Printf("ConsumeGas response: %+v  err: %+v\n", response, err)
	}()

	// Recursive chain calls are exponentially more expensive
	gasUsed := uint64(float64(msg.Amount) * math.Pow(2, float64(rpcservice.k.currentDepth-1)))
	ctx.GasMeter().ConsumeGas(gasUsed, "gasUsed")
	// fmt.Printf("gasUsed: %v\n", gasUsed)
	*response = ConsumeGasResponse{
		GasConsumed: int64(ctx.GasMeter().GasConsumed()),
		GasLimit:    int64(ctx.GasMeter().Limit()),
	}
	return nil
}

type GasLimitRequest struct{}

type GasLimitResponse struct {
	GasConsumed int64 `protobuf:"bytes,1,opt,name=GasConsumed,proto3" json:"GasConsumed,omitempty"`
	GasLimit    int64 `protobuf:"bytes,1,opt,name=GasLimit,proto3" json:"GasLimit,omitempty"`
}

func (rpcservice *RpcService) Gaslimit(_ *http.Request, msg *GasLimitRequest, response *GasLimitResponse) (err error) {
	ctx := sdk.UnwrapSDKContext(rpcservice.goCtx)
	gm := ctx.GasMeter()
	*response = GasLimitResponse{
		GasConsumed: int64(gm.GasConsumed()),
		GasLimit:    int64(gm.Limit()),
	}
	return nil
}
