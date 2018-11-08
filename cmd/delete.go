package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	service "github.com/owtotwo/agenda/entity/service"
	tools "github.com/owtotwo/agenda/entity/tools"
)

// deleteCmd represents the delete command
var udeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete user account",
	Long: `Use this command to delete your account, meetings included.`,
	Run: func(cmd *cobra.Command, args []string) {
		var Service service.Service
		service.StartAgenda(&Service)
		// check whether other user logged in
		ok, curUsername := Service.AutoUserLogin()
		if ok == true {
			fmt.Println(strings.Join([]string{curUsername, "@:"}, ""))
		} else {
			fmt.Fprintln(os.Stderr, "Please Login in First")
			os.Exit(1)
		}

		// hints to ensure and enter password to delete User
		var password string
		fmt.Print("Plase enter password: ")
		reader := bufio.NewReader(os.Stdin)
		data, _, _ := reader.ReadLine()
		password = string(data)
		// delete user and meetings it participate
		if Service.DeleteUser(curUsername, password) == false {
			fmt.Fprintln(os.Stderr, "Some mistakes happend in DeleteUser.")
			os.Exit(1)
		}
		fmt.Println("Success : delete ", curUsername)
		ok = Service.UserLogout()
		if ok == false {
			fmt.Fprintln(os.Stderr, "some mistake happend in UserLogout")
			os.Exit(1)
		}
		Service.QuitAgenda()
		os.Exit(0)
	},
}

var mdeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete meeting",
	Long: `Use this command to delete specific meeting.`,
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
			tools.LogInfoOrErrorIntoFile(name, true, fmt.Sprintf("Delete meeting with no user login."))
			os.Exit(0)
		}

		if meetingName == "" {
			fmt.Fprintln(os.Stderr, "error: Meeting name is required.")
			tools.LogInfoOrErrorIntoFile(name, false, fmt.Sprintf("Delete meeting with no title login."))
			os.Exit(0)
		}
		ok = Service.DeleteMeeting(name, meetingName)
		if ok {
			fmt.Printf("Delete %s finished.\n", meetingName)
			tools.LogInfoOrErrorIntoFile(name, true, fmt.Sprintf("Delete %s finished.", meetingName))
		} else {
			fmt.Printf("Can not delete the meeting called %s.\n", meetingName)
		}
	},
}

func init() {
	userCmd.AddCommand(udeleteCmd)
	meetingCmd.AddCommand(mdeleteCmd)

	udeleteCmd.Flags().StringP("username", "u", "", "Delete user")
	mdeleteCmd.Flags().StringVarP(&meetingName, "name", "", "", "meeting name to be deleted")
}
