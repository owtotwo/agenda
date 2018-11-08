package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	service "github.com/owtotwo/agenda/entity/service"
	tools "github.com/owtotwo/agenda/entity/tools"
)

// leaveCmd represents the leave command
var leaveCmd = &cobra.Command{
	Use:   "leave",
	Short: "Leave meeting",
	Long: `Use this command to leave a meeting.`,
	Run: func(cmd *cobra.Command, args []string) {
		var Service service.Service
		service.StartAgenda(&Service)
		// check whether other user logged in
		ok, name := Service.AutoUserLogin()
		if ok == true {
			fmt.Println(strings.Join([]string{name,"@:"}, ""))
		}
		if !ok {
			fmt.Fprintln(os.Stderr, "error: No current logged user.")
			tools.LogInfoOrErrorIntoFile(name, false, fmt.Sprintf("Leave meeting with no user login."))
			os.Exit(0)
		}
		
		if meetingName == "" {
			fmt.Fprintln(os.Stderr, "error: Meeting theme is required.")
			tools.LogInfoOrErrorIntoFile(name, false, fmt.Sprintf("Leave meeting with no title."))
			os.Exit(0)
		}
		
		ok = Service.QuitMeeting(name, meetingName)
		if ok {
			fmt.Printf("Finish leaving meeting %s.\n", meetingName)
			tools.LogInfoOrErrorIntoFile(name, true, fmt.Sprintf("Leave meeting %s.", meetingName))
		} else {
			fmt.Printf("Can not leave this meeting %s.\n", meetingName)
			tools.LogInfoOrErrorIntoFile(name, true, fmt.Sprintf("Can not leave meeting %s.", meetingName))
		}
	},
}

func init() {
	meetingCmd.AddCommand(leaveCmd)
	
	leaveCmd.Flags().StringVarP(&meetingName, "name", "", "", "the name of meeting to be managed")
}
