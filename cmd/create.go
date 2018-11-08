package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	service "github.com/owtotwo/agenda/entity/service"
)

var ucreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create user account",
	Long:  `Use this command to create a new user account.`,
	Run: func(cmd *cobra.Command, args []string) {
		// get service
		var Service service.Service
		service.StartAgenda(&Service)
		// check whether other user logged in
		ok, CurUsername := Service.AutoUserLogin()
		if ok == true {
			fmt.Println(strings.Join([]string{CurUsername, "@:"}, ""))
			fmt.Fprintln(os.Stderr, strings.Join([]string{CurUsername, " had logged in. Please log out first."}, ""))
			os.Exit(1)
		}
		// get createUser information
		createUsername, _ := cmd.Flags().GetString("username")
		createEmail, _ := cmd.Flags().GetString("email")
		createPhone, _ := cmd.Flags().GetString("phone")
		// check whether username, password, email or phone empty
		if createUsername == "" || createEmail == "" ||
			createPhone == "" {
			fmt.Fprintln(os.Stderr, "error : Username, Email or Phone is(are) empty")
			os.Exit(1)
		}
		// get the password
		var createPassword string
		var prePassword string
		times := 1
		reader := bufio.NewReader(os.Stdin)
		for {
			if times == 1 {
				fmt.Print("Please enter the password you want to create: ")
				data, _, _ := reader.ReadLine()
				createPassword = string(data)
			} else {
				fmt.Print("Please enter the password again: ")
				data, _, _ := reader.ReadLine()
				createPassword = string(data)
				if createPassword == prePassword {
					break
				} else {
					fmt.Println("The two passwords entered are not consistent. \nPlease restart setting password.")
				}
			}
			times *= -1
			prePassword = createPassword
		}
		// check whether User is registed
		ok = Service.UserRegister(createUsername, createPassword, createEmail, createPhone)
		if ok == false {
			fmt.Println(createUsername, " has been registered")
			os.Exit(1)
		}
		fmt.Println("Sucess : Register ", createUsername)
		Service.QuitAgenda()
		os.Exit(0)
	},
}

func init() {
	userCmd.AddCommand(ucreateCmd)
	
	ucreateCmd.Flags().StringP("username", "u", "", "Create Username")
	ucreateCmd.Flags().StringP("email", "e", "", "Create Email")
	ucreateCmd.Flags().StringP("phone", "p", "", "Create Phone")
}
