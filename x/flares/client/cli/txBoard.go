package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/wangfeiping/flares/x/flares/types"
)

func CmdCreateBoard() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-board [baseDenom] [acceptDenom] [source]",
		Short: "Creates a new board",
		Args:  cobra.MinimumNArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsBaseDenom := string(args[0])
			argsAcceptDenom := string(args[1])
			argsSource := string(args[2])

			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgBoard(clientCtx.GetFromAddress(),
				string(argsBaseDenom), string(argsAcceptDenom), string(argsSource))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
