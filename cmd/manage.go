package cmd

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
	"github.com/spf13/cobra"

	service "github.com/owtotwo/agenda/entity/service"	
	tools "github.com/owtotwo/agenda/entity/tools"
)

// manageCmd represents the manage command
var manageCmd = &cobra.Command{
	Use:   "manage",
	Short: "Manage meeting",
	Long: `Create meeting.`,
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
			tools.LogInfoOrErrorIntoFile(name, false, fmt.Sprintf("Manage meeting with no user login."))
			os.Exit(0)
		}
		
		if meetingName == "" {
			fmt.Fprintln(os.Stderr, "error: Meeting theme is required.")
			tools.LogInfoOrErrorIntoFile(name, false, fmt.Sprintf("Manage  meeting %s with no title.", meetingName))
			os.Exit(0)
		}
		
		meetingList := Service.MeetingQueryByTitle(name, meetingName)
		if len(meetingList) == 0 {
			fmt.Println("No matching meeting with the given theme.")
			os.Exit(1)
		}

		// delete users
		if isDelete {
			var participator []string
			fmt.Println("Participators:")
			for i, v := range meetingList[0].GetParticipators() {
				participator = append(participator, v)
				fmt.Printf("%d. %s\n", i + 1, v)
			}
			fmt.Print("Please input the number you want to remove: ")
			var inputNums string
			reader := bufio.NewReader(os.Stdin)
			data, _, _ := reader.ReadLine()
			inputNums = string(data)
			chosenList := strings.Split(inputNums, " ")
			var toBeRemovedParticipators []string
			for _, v := range chosenList {
				num, err := strconv.Atoi(v)
				if err != nil || num <= 0 || num > len(participator) {
					fmt.Fprintln(os.Stderr, "error: Invalid input")
					os.Exit(1)
				}
				toBeRemovedParticipators = append(toBeRemovedParticipators, participator[num - 1])
			}
			for _, v := range toBeRemovedParticipators {
				ok := Service.DeleteParticipatorByTitle(name, meetingName, v)
				if ok {
					fmt.Printf("%s was removed.\n", v)
					tools.LogInfoOrErrorIntoFile(name, true, fmt.Sprintf("Remove %s from meeting %s.", v, meetingName))
				} else {
					fmt.Printf("%s can not be removed.\n", v)
					tools.LogInfoOrErrorIntoFile(name, false, fmt.Sprintf("Can not remove %s from meeting %s.", v, meetingName))
				}
			}
		} else {
			// add users
			fmt.Println("You can choose some of them to add to your meeting:")
			userList := Service.ListAllUsers()
			for i, v := range userList {
				fmt.Printf("%d. %s\n", i + 1, v.GetUserName())
			}
			fmt.Print("Please input the number of users you want to add(separate with blank): ")
			var userNums string
			reader := bufio.NewReader(os.Stdin)
			data, _, _ := reader.ReadLine()
			userNums = string(data)
			userNumList := strings.Split(userNums, " ")
			for _, v := range userNumList {
				if len(v) == 0 {
					continue
				}
				i, ok := strconv.Atoi(v)
				if ok != nil || i > len(userList) || i <= 0 {
					fmt.Fprintln(os.Stderr, "error: Invalid input.")
					os.Exit(0)
				}
				if Service.AddParticipatorByTitle(name, meetingName, userList[i - 1].GetUserName()) {
					fmt.Printf("%s was added.\n", userList[i - 1].GetUserName())
					tools.LogInfoOrErrorIntoFile(name, true, fmt.Sprintf("Add %s to meeting %s.", userList[i - 1].GetUserName(), meetingName))
				} else {
					fmt.Printf("%s can not be added.\n", userList[i - 1].GetUserName())
					tools.LogInfoOrErrorIntoFile(name, false, fmt.Sprintf("Can not add %s to meeting %s.", userList[i - 1].GetUserName(), meetingName))
				}
			}
		}
	},
}

func init() {
	meetingCmd.AddCommand(manageCmd)
	
	manageCmd.Flags().StringVarP(&meetingName, "name", "", "", "the name of meeting to be managed")
	manageCmd.Flags().BoolVarP(&isDelete, "", "d", false, "Delete user(s) from a meeting")
}
