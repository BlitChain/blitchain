package cli

import (
	"blit/x/script/types"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

func CmdCreateScript() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-script --msg-types [msg types] --code [code] [--grantee [grantee address]]",
		Short: "Create a new script",
		RunE: func(cmd *cobra.Command, args []string) error {
			msgTypesStr, _ := cmd.Flags().GetString("msg-types")
			code, _ := cmd.Flags().GetString("code")
			grantee, _ := cmd.Flags().GetString("grantee")

			// Split on whitespace to get individual message types
			msgTypes := strings.Fields(msgTypesStr)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateScript(
				clientCtx.GetFromAddress().String(),
				code,
				msgTypes,
				grantee,
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String("msg-types", "", "Types of messages")
	cmd.Flags().String("code", "", "Script code")
	cmd.Flags().String("grantee", "", "(optional) The authz grantee that has permission to update the script")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateScript() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-script --address [script address] --code [code] [--grantee [grantee address]]",
		Short: "Update a script",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			code, _ := cmd.Flags().GetString("code")
			address, _ := cmd.Flags().GetString("address")
			grantee, _ := cmd.Flags().GetString("grantee")

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateScript(
				address,
				code,
				grantee,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String("address", "", "The script address to update")
	cmd.Flags().String("code", "", "Script code")
	cmd.Flags().String("grantee", "", "(optional) The authz grantee that has permission to update the script")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
