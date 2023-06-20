package cli

import (
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"gitlab.qredo.com/qrdochain/fusionchain/x/treasury/types"
)

var _ = strconv.Itoa(0)

func CmdSignatureRequests() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "signature-requests [pending|fulfilled|rejected|all]",
		Short: "Query SignatureRequests, optionally filtering by their current status",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			params := &types.QuerySignatureRequestsRequest{
				Pagination: pageReq,
				XStatus:    nil,
			}
			switch strings.ToLower(args[0]) {
			case "pending":
				params.XStatus = &types.QuerySignatureRequestsRequest_Status{
					Status: types.SignRequestStatus_SIGN_REQUEST_STATUS_PENDING,
				}
			case "fulfilled":
				params.XStatus = &types.QuerySignatureRequestsRequest_Status{
					Status: types.SignRequestStatus_SIGN_REQUEST_STATUS_FULFILLED,
				}
			case "rejected":
				params.XStatus = &types.QuerySignatureRequestsRequest_Status{
					Status: types.SignRequestStatus_SIGN_REQUEST_STATUS_REJECTED,
				}
			}

			res, err := queryClient.SignatureRequests(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}