package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	service "github.com/owtotwo/agenda/entity/service"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Sign out",
	Long:  `Use this command to sign out`,
	Run: func(cmd *cobra.Command, args []string) {
		// get service
		var Service service.Service
		service.StartAgenda(&Service)
		// check whether other user logged in
		ok, name := Service.AutoUserLogin()
		if ok == true {
			fmt.Println(strings.Join([]string{name, "@:"}, ""))
		}
		// check Whether CurUser exits
		ok, CurUsername := Service.AutoUserLogin()
		if ok == false {
			fmt.Fprintln(os.Stderr, "error : Current User not exits")
			os.Exit(1)
		}
		ok = Service.UserLogout()
		if ok == false {
			fmt.Fprintln(os.Stderr, "error : some mistakes happend in UserLogout")
			os.Exit(1)
		}
		fmt.Println("Success : ", CurUsername, " Logout")
		Service.QuitAgenda()
		os.Exit(0)
	},
}

func init() {
	userCmd.AddCommand(logoutCmd)
}
