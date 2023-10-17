package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"blit/x/storage/types"
)

func CmdListStorage() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-storage --address [address] --filter-prefix [filter prefix]",
		Short: "list all storage",
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

			params := &types.QueryAllStorageRequest{
				FilterAddress:     argAddress,
				FilterIndexPrefix: argPrefix,
				Pagination:        pageReq,
			}

			res, err := queryClient.StorageAll(cmd.Context(), params)
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

func CmdShowStorage() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-storage [address] [index]",
		Short: "shows a storage",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			argAddress := args[0]
			argIndex := args[1]

			params := &types.QueryGetStorageRequest{
				Address: argAddress,
				Index:   argIndex,
			}

			res, err := queryClient.Storage(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
