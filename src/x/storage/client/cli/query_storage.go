package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"blit/x/storage/types"
)

func CmdFilterStorage() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "filter-storage --address [address] --filter-prefix [filter prefix]",
		Short: "Filter storage by address and/or prefix",
		RunE: func(cmd *cobra.Command, args []string) error {
			argAddress, _ := cmd.Flags().GetString("filter-address")
			argPrefix, _ := cmd.Flags().GetString("filter-prefix")
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryFilterStorageRequest{
				FilterAddress:     argAddress,
				FilterIndexPrefix: argPrefix,
				Pagination:        pageReq,
			}

			res, err := queryClient.FilterStorage(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	cmd.Flags().String("filter-address", "", "Required address to filter on")
	cmd.Flags().String("filter-prefix", "", "Optional filter prefix to filter index on")
	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdGetStorage() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-storage [address] [index]",
		Short: "get a storage by address and index",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			argAddress := args[0]
			argIndex := args[1]

			params := &types.QueryStorageDetailRequest{
				Address: argAddress,
				Index:   argIndex,
			}

			res, err := queryClient.StorageDetail(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
