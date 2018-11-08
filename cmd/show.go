package cmd

import (
	"os"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	service "github.com/owtotwo/agenda/entity/service"
	tools "github.com/owtotwo/agenda/entity/tools"
)

var startTime string
var endTime string

// showCmd represents the show command
var ushowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show user account",
	Long: `Use this command to show every user's information.`,
	Run: func(cmd *cobra.Command, args []string) {
		var Service service.Service
		service.StartAgenda(&Service)
		// check whether user login
		ok, CurUsername := Service.AutoUserLogin()
		if ok == true {
			fmt.Println(strings.Join([]string{CurUsername,"@:"}, ""))
		} else {
			fmt.Fprintln(os.Stderr, "error : No User has Logined in")
			os.Exit(1)
		}
		// get email and phone by username
		users := Service.ListAllUsers()
		fmt.Printf("%-15s%-25s%-25s\n", "Username", "E-mail", "phone number")
		for _, user := range users {
			fmt.Printf("%-15s%-25s%-25s\n", user.GetUserName(), user.GetEmail(), user.GetPhone())
		}
		fmt.Printf("\nTotal number is %d\n", len(users))

		os.Exit(0)
	},
}

var mshowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show meeting information",
	Long: `Use this command to show meeting information.`,
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
			tools.LogInfoOrErrorIntoFile(name, false, fmt.Sprintf("Show meeting with no user login."))
			os.Exit(0)
		}
		
		if startTime == "" || endTime == "" {
			fmt.Fprintln(os.Stderr, "error: Start time and end time is required.")
			tools.LogInfoOrErrorIntoFile(name, true, fmt.Sprintf("Show meeting with no invalid time."))
			os.Exit(0)
		}
		meetingList := Service.MeetingQueryByUserAndTime(name, startTime, endTime)

		// print all meetings with the given name
		if len(meetingList) == 0 {
			fmt.Println("--·--·--·--·--·--·--·--·--·--·--")
			fmt.Println("No matching meeting")
			fmt.Println("--·--·--·--·--·--·--·--·--·--·--")
		} else {
			fmt.Println("--·--·--·--·--·--·--·--·--·--·--")
			for _, v := range meetingList {
				fmt.Printf("Theme: %s\n", v.GetTitle())
				fmt.Printf("Sponsor: %s\n", v.GetSponsor())
				fmt.Printf("Start time: %s\n", v.GetStartDate())
				fmt.Printf("End time: %s\n", v.GetEndDate())
				fmt.Printf("Participator: %s\n", strings.Join(v.GetParticipators(), ", "))
				fmt.Println("--·--·--·--·--·--·--·--·--·--·--")
			}
		}
	},
}

func init() {
	userCmd.AddCommand(ushowCmd)
	meetingCmd.AddCommand(mshowCmd)

	mshowCmd.Flags().StringVarP(&startTime, "start", "s", "", "Start time of meeting")
	mshowCmd.Flags().StringVarP(&endTime, "end", "e", "", "End time of meeting")
}
