package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/wangfeiping/flares/x/sealedmonsters/types"
)

func CmdCreateMonster() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-monster [description] [solutionHash] [reward] [solution] [scavenger]",
		Short: "Creates a new monster",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsDescription := string(args[0])
			argsSolutionHash := string(args[1])
			argsReward := string(args[2])
			argsSolution := string(args[3])
			argsScavenger := string(args[4])

			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgMonster(clientCtx.GetFromAddress(), string(argsDescription), string(argsSolutionHash), string(argsReward), string(argsSolution), string(argsScavenger))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
