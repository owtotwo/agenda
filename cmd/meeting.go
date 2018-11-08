package cmd

import (
	"github.com/spf13/cobra"
)

var meetingName string
var isDelete bool

var meetingCmd = &cobra.Command{
	Use:   "meeting",
	Short: "Manage meetings",
	Long: `You can use this command to manage your meetings, 
but before you use it, please make sure you have already signed up an account 
and signed in the system.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	RootCmd.AddCommand(meetingCmd)
}
