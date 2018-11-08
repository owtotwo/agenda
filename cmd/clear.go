package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	service "github.com/owtotwo/agenda/entity/service"
	tools "github.com/owtotwo/agenda/entity/tools"
)

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all meetings",
	Long: `Remove all meetings you created.`,
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
			tools.LogInfoOrErrorIntoFile(name, true, fmt.Sprintf("Clear meeting with no user login."))
			os.Exit(0)
		}

		fmt.Print("Are you sure you want to clear all of your meetings? (y/n) ")
		var confirm string
		fmt.Scanf("%s", &confirm)
		if confirm == "y" {
			ok = Service.DeleteAllMeetings(name)
			if ok {
				fmt.Println("All of the meeting have been deleted.")
				tools.LogInfoOrErrorIntoFile(name, true, fmt.Sprintf("%s clear all meetings.", name))
			} else {
				fmt.Println("Some problems occured when clear your meetings.")
				tools.LogInfoOrErrorIntoFile(name, true, fmt.Sprintf("%s can not clear all the meetings.", name))
			}
		} else {
			fmt.Println("You canceled the process.")
		}
	},
}

func init() {
	meetingCmd.AddCommand(clearCmd)
}
