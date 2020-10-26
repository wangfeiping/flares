package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/wangfeiping/flares/x/sealedmonsters/types"
)

func CmdCreateReveal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-reveal [solutionHash] [solution] [scavenger]",
		Short: "Creates a new reveal",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsSolutionHash := string(args[0])
			argsSolution := string(args[1])
			argsScavenger := string(args[2])

			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgReveal(clientCtx.GetFromAddress(), string(argsSolutionHash), string(argsSolution), string(argsScavenger))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
