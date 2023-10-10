package cli

import (
	"blit/x/storage/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

func CmdCreateStorage() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-storage --address [value] --index [value] --data [value] --grantee [value]",
		Short: "Create a new storage",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Retrieve flag values
			indexAddress, _ := cmd.Flags().GetString("address")
			indexIndex, _ := cmd.Flags().GetString("index")
			argData, _ := cmd.Flags().GetString("data")
			argGrantee, _ := cmd.Flags().GetString("grantee")

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateStorage(
				indexAddress,
				indexIndex,
				argData,
				argGrantee,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	// Add flags
	cmd.Flags().String("address", "", "The script address (or grantor)")
	cmd.Flags().String("index", "", "The index for the storage")
	cmd.Flags().String("data", "", "The data to store")
	cmd.Flags().String("grantee", "", "(optional) The authz grantee for this storage")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateStorage() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-storage --address [value] --index [value] --data [value] --grantee [value]",
		Short: "Update a storage",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexAddress, _ := cmd.Flags().GetString("address")
			indexIndex, _ := cmd.Flags().GetString("index")
			argData, _ := cmd.Flags().GetString("data")
			argGrantee, _ := cmd.Flags().GetString("grantee")

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateStorage(
				indexAddress,
				indexIndex,
				argData,
				argGrantee,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String("address", "", "The script address (or grantor)")
	cmd.Flags().String("index", "", "The index for the storage")
	cmd.Flags().String("data", "", "The data to update")
	cmd.Flags().String("grantee", "", "(optional) The authz grantee for this storage")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteStorage() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-storage --address [value] --index [value] --grantee [value]",
		Short: "Delete a storage",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexAddress, _ := cmd.Flags().GetString("address")
			indexIndex, _ := cmd.Flags().GetString("index")
			argGrantee, _ := cmd.Flags().GetString("grantee")

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteStorage(
				indexAddress,
				indexIndex,
				argGrantee,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String("address", "", "The script address (or grantor)")
	cmd.Flags().String("index", "", "The index for the storage")
	cmd.Flags().String("grantee", "", "(optional) The authz grantee for this storage")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
