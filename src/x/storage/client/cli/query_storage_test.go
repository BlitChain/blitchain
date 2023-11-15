package cli_test

import (
	"fmt"
	"strconv"
	"testing"

	tmcli "github.com/cometbft/cometbft/libs/cli"
	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"blit/testutil/network"
	"blit/testutil/nullify"
	"blit/x/storage/client/cli"
	"blit/x/storage/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func (s *IntegrationTestSuite) networkWithStorageObjects(n int) (*network.Network, []types.Storage) {
	s.T().Helper()
	state := types.GenesisState{}
	for i := 0; i < n; i++ {
		storage := types.Storage{
			Index: strconv.Itoa(i),
		}
		nullify.Fill(&storage)
		state.StorageList = append(state.StorageList, storage)
	}
	return s.network(&state), state.StorageList
}

func (s *IntegrationTestSuite) TestShowStorage() {
	var (
		net, objs = s.networkWithStorageObjects(2)
		ctx       = net.Validators[0].ClientCtx
		common    = []string{
			fmt.Sprintf("--%s=json", tmcli.OutputFlag),
		}
	)
	tests := []struct {
		desc    string
		idIndex string

		args []string
		err  error
		obj  types.Storage
	}{
		{
			desc:    "found",
			idIndex: objs[0].Index,

			args: common,
			obj:  objs[0],
		},
		{
			desc:    "not found",
			idIndex: strconv.Itoa(100000),

			args: common,
			err:  status.Error(codes.NotFound, "not found"),
		},
	}
	for _, tc := range tests {
		s.T().Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.idIndex,
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowStorage(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
				return
			}
			require.NoError(t, err)
			var resp types.QueryGetStorageResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.NotNil(t, resp.Storage)
			require.Equal(t,
				nullify.Fill(&tc.obj),
				nullify.Fill(&resp.Storage),
			)
		})
	}
}

func (s *IntegrationTestSuite) TestListStorage() {
	var (
		net, objs = s.networkWithStorageObjects(5)
		ctx       = net.Validators[0].ClientCtx
	)
	request := func(next []byte, offset, limit uint64, total bool) []string {
		args := []string{
			fmt.Sprintf("--%s=json", tmcli.OutputFlag),
		}
		if next == nil {
			args = append(args, fmt.Sprintf("--%s=%d", flags.FlagOffset, offset))
		} else {
			args = append(args, fmt.Sprintf("--%s=%s", flags.FlagPageKey, next))
		}
		args = append(args, fmt.Sprintf("--%s=%d", flags.FlagLimit, limit))
		if total {
			args = append(args, fmt.Sprintf("--%s", flags.FlagCountTotal))
		}
		return args
	}
	s.T().Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(objs); i += step {
			args := request(nil, uint64(i), uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListStorage(), args)
			require.NoError(t, err)
			var resp types.QueryAllStorageResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.Storage), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.Storage),
			)
		}
	})
	s.T().Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(objs); i += step {
			args := request(next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListStorage(), args)
			require.NoError(t, err)
			var resp types.QueryAllStorageResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.Storage), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.Storage),
			)
			next = resp.Pagination.NextKey
		}
	})
	s.T().Run("Total", func(t *testing.T) {
		args := request(nil, 0, uint64(len(objs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListStorage(), args)
		require.NoError(t, err)
		var resp types.QueryAllStorageResponse
		require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(objs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(objs),
			nullify.Fill(resp.Storage),
		)
	})
}
