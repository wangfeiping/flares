package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/wangfeiping/flares/x/flares/types"
)

func CmdListContract() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-contract [ pending(default) / success / fail ]",
		Short: "list (pending success or fail) contracts",
		RunE: func(cmd *cobra.Command, args []string) error {
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

			params := &types.QueryAllContractRequest{
				Pagination: pageReq,
			}
			if len(args) > 0 {
				params.State = args[0]
			}

			res, err := queryClient.AllContract(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
