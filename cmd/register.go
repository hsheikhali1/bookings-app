package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register for a conference",
	Long:  `Register for one of the available conferences [list conferences]`,
	Run: func(cmd *cobra.Command, args []string) {
		register(args)
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
}

func register(args []string) {

	// get number of tickets
	// get conference name
	// save this to a file or something, then run `./app purchase` to simulate a purchased ticket command

	if len(args) <= 0 {
		fmt.Println("No arguments were passed into the register command. For help run `bookingapp help`")
		os.Exit(1)
	}

	fmt.Printf("Registering for `%v` conference\n", args[0])
}
