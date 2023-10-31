package cli

import (
	"strconv"

	"blit/x/script/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdEval() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "eval --caller-address [value] --script-address [value] --extra-code [value] --function-name [value] --kwargs [value]",
		Short: "Query eval",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCallerAddress, _ := cmd.Flags().GetString("caller-address")
			argScriptAddress, _ := cmd.Flags().GetString("script-address")
			argExtraCode, _ := cmd.Flags().GetString("extra-code")
			argFunctionName, _ := cmd.Flags().GetString("function-name")
			argKwargs, _ := cmd.Flags().GetString("kwargs")

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryEval{

				CallerAddress: argCallerAddress,
				ScriptAddress: argScriptAddress,
				ExtraCode:     argExtraCode,
				FunctionName:  argFunctionName,
				Kwargs:        argKwargs,
			}

			res, err := queryClient.Eval(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	cmd.Flags().String("caller-address", "", "The address of the CALLER (or grantor)")
	cmd.Flags().String("script-address", "", "The address of the script")
	cmd.Flags().String("extra-code", "", "Optional extra code to run if the caller address is the script address")
	cmd.Flags().String("function-name", "", "Function name to call")
	cmd.Flags().String("kwargs", "", "Keyword arguments to pass to the function")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
