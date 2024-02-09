package keeper

import (
	blittypes "blit/x/blit/types"
	"blit/x/services/types"
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server/config"
	gogogrpc "github.com/cosmos/gogoproto/grpc"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/spf13/viper"
)

// RegisterNodeService registers the node gRPC service on the provided gRPC router.
func RegisterNodeService(clientCtx client.Context, server gogogrpc.Server, config config.Config) {

	types.RegisterServiceServer(server, NewQueryServer())
}

var _ types.ServiceServer = queryServer{}

type queryServer struct {
	config blittypes.BlitConfig
}

func NewQueryServer() types.ServiceServer {
	return queryServer{}
}

// RegisterGRPCGatewayRoutes mounts the node gRPC service's GRPC-gateway routes
// on the given mux object.
func RegisterGRPCGatewayRoutes(clientConn gogogrpc.ClientConn, mux *runtime.ServeMux) {
	err := types.RegisterServiceHandlerClient(context.Background(), mux, types.NewServiceClient(clientConn))
	if err != nil {
		panic(err)
	}
}

func (s queryServer) Endpoints(goCtx context.Context, req *types.EndpointsRequest) (*types.EndpointsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	//ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.EndpointsResponse{
		ApiUrl: viper.GetString("blit.public_rest_api_url"),
		RpcUrl: viper.GetString("blit.public_comet_rpc_url"),
	}, nil
}
