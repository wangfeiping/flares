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
		Use:   "summon-monster [description] [solution] [reward]",
		Short: "Summon the monster",
		Args:  cobra.MinimumNArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsDescription := string(args[0])
			argsSolution := string(args[1])
			argsReward := string(args[2])

			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgMonster(clientCtx.GetFromAddress(),
				string(argsDescription), string(argsSolution), string(argsReward))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
