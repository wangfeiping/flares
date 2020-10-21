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
		Use:   "create-board [base] [baseDenom] [accept] [acceptDenom] [source]",
		Short: "Creates a new board",
		Args:  cobra.MinimumNArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsBase := string(args[0])
			argsBaseDenom := string(args[1])
			argsAccept := string(args[2])
			argsAcceptDenom := string(args[3])
			argsSource := string(args[4])

			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgBoard(clientCtx.GetFromAddress(),
				string(argsBase), string(argsBaseDenom), "",
				string(argsAccept), string(argsAcceptDenom), "", string(argsSource))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
