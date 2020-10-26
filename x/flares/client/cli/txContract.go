package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/wangfeiping/flares/x/flares/types"
)

func CmdCreateContract() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-contract [module] [key] [accept] [durationHeight] [bottom]",
		Short: "Creates a new contract",
		Args:  cobra.MinimumNArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			module := args[0]
			argsKey := string(args[1])
			argsAccept := string(args[2])
			argsDurationHeight, err := strconv.Atoi(args[3])
			if err != nil {
				return err
			}
			argsBottom := string(args[4])

			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err = client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgContract(clientCtx.GetFromAddress(),
				module, string(argsKey), string(argsAccept),
				int32(argsDurationHeight), string(argsBottom))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
