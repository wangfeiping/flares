package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/wangfeiping/flares/x/sealedmonsters/types"
)

func CmdCreateSeal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-seal [solutionHash] [solutionScavengerHash] [scavenger]",
		Short: "Creates a new seal",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsSolutionHash := string(args[0])
			argsSolutionScavengerHash := string(args[1])
			argsScavenger := string(args[2])

			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgSeal(clientCtx.GetFromAddress(), string(argsSolutionHash), string(argsSolutionScavengerHash), string(argsScavenger))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
