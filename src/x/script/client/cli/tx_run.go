package cli

import (
	"strconv"

	"blit/x/script/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdRun() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run --caller-address [value] --script-address [value] --extra-code [value] --function-name [value] --kwargs [value] --grantee [value]",
		Short: "Broadcast message run",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCallerAddress, _ := cmd.Flags().GetString("caller-address")
			argScriptAddress, _ := cmd.Flags().GetString("script-address")
			argExtraCode, _ := cmd.Flags().GetString("extra-code")
			argFunctionName, _ := cmd.Flags().GetString("function-name")
			argKwargs, _ := cmd.Flags().GetString("kwargs")
			grantee, _ := cmd.Flags().GetString("grantee")

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRun(
				argCallerAddress,
				argScriptAddress,
				argExtraCode,
				argFunctionName,
				argKwargs,
				grantee,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String("caller-address", "", "callerAddress")
	cmd.Flags().String("script-address", "", "scriptAddress")
	cmd.Flags().String("extra-code", "", "extraCode")
	cmd.Flags().String("function-name", "", "functionName")
	cmd.Flags().String("kwargs", "", "kwargs")
	cmd.Flags().String("grantee", "", "grantee")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
