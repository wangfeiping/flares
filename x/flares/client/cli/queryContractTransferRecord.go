package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/wangfeiping/flares/x/flares/types"
)

func CmdListContractTransferRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-contractTransferRecord [contract receiver address]",
		Args:  cobra.MinimumNArgs(1),
		Short: "list all contractTransferRecord",
		RunE: func(cmd *cobra.Command, args []string) error {
			receiver := args[0]

			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllContractTransferRecordRequest{
				Receiver:   receiver,
				Pagination: pageReq,
			}

			res, err := queryClient.AllContractTransferRecord(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
