package cli

import (
	"github.com/Leave-Management-System/lms-cosmos/x/lms/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func NewQueryCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "LMS",
		DisableFlagParsing:         false,
		SuggestionsMinimumDistance: 4,
		RunE:                       client.ValidateCmd,
	}
	txCmd.AddCommand(
		GetStudentsCmd(),
		GetLeaveRequestsCmd(),
		GetLeaveApprovedRequestsCmd(),
		GetLeaveStatusCmd(),
	)
	return txCmd
}

func GetStudentsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "getstudent",
		Short: "To get the list of all students",
		Long:  "To get teh list of all students",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			handleError(err)
			getStudentsRequest := &types.GetStudentsRequest{}
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.GetStudentsQuery(cmd.Context(), getStudentsRequest)
			handleError(err)
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetLeaveRequestsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "getleaves",
		Short: "To get the list of all the leave requets",
		Long:  "To get the list of all the leave requests",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			handleError(err)
			getLeavesRequest := &types.GetLeaveRequestsRequest{}
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.GetLeaveRequestsQuery(cmd.Context(), getLeavesRequest)
			handleError(err)
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetLeaveStatusCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "getstatus",
		Short: "To get the leave status",
		Long:  "To get the leave status",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			handleError(err)
			getLeaveStatusRequest := &types.GetLeaveStatusRequest{
				LeaveID: args[0],
			}
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.GetLeaveStatusQuery(cmd.Context(), getLeaveStatusRequest)
			handleError(err)
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetLeaveApprovedRequestsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "getapproved",
		Short: "To get the list of approved leaves",
		Long:  "To get the list of approved leaves",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			handleError(err)
			getLeavesRequest := &types.GetLeaveApprovedRequestsRequest{}
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.GetLeaveApprovedRequestsQuery(cmd.Context(), getLeavesRequest)
			handleError(err)
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}