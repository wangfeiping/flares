package cli

import (
  
	"github.com/spf13/cobra"

    "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/wangfeiping/flares/x/flares/types"
)

func CmdCreateContractTransferRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-contractTransferRecord [hash] [from] [to] [amount]",
		Short: "Creates a new contractTransferRecord",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
      argsHash := string(args[0])
      argsFrom := string(args[1])
      argsTo := string(args[2])
      argsAmount := string(args[3])
      
        	clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgContractTransferRecord(clientCtx.GetFromAddress(), string(argsHash), string(argsFrom), string(argsTo), string(argsAmount))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

    return cmd
}
