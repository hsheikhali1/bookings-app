package cmd

import (
	"bookings-app/pkg/utils"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register for a conference.",
	Long:  `Register for one of the available conferences.`,
	Run:   run,
}

func init() {
	rootCmd.AddCommand(registerCmd)

	registerCmd.PersistentFlags().StringP("name", "n", "", "Pass the name flag to specify which conferece you want to attend.")

	// this flag should only return conferences that don't have 0 tickets available
	registerCmd.PersistentFlags().BoolP("list", "l", false, "List available conferences")

	registerCmd.PersistentFlags().UintP("tickets", "t", 0, "List available conferences")
}

func register(args []string, cmd *cobra.Command) {
	name, _ := cmd.Flags().GetString("name")
	list, _ := cmd.Flags().GetBool("list")
	tickets, _ := cmd.Flags().GetUint("tickets")

	conferences := new(utils.Conferences)

	// get number of tickets
	// get conference name
	// save this to a file or something, then run `./app purchase` to simulate a purchased ticket command

	if name != "" {
		conferenceName, err := getConferenceByName(name, conferences.GenerateConferences())

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("Registering for %v conference\n", conferenceName)
		conferences.PurchaseTickets(conferenceName, tickets)

	} else if list {
		fmt.Println(conferences.GenerateConferences())
	} else {

		// validate that conference is a legit conference
		// if not, return error, else you're good

		if len(args) <= 0 {
			fmt.Println("No arguments were passed into the register command. For help run `bookingapp help`")
			os.Exit(1)
		}

		if args[0] == "" {
			fmt.Println("Argument cannot be an empty string")
			os.Exit(1)
		}

		conferenceName, err := getConferenceByName(args[0], conferences.GenerateConferences())

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("Registering for %v conference\n", conferenceName)
	}
}

func getConferenceByName(name string, conferences []utils.Conference) (string, error) {
	var confName string
	var error error = nil

	for _, conference := range conferences {
		if name == conference.ConferenceName {
			confName = conference.ConferenceName
		}
	}

	if confName == "" {
		error = fmt.Errorf("could not find conference: %v in list of conferences", name)
	}

	return confName, error
}

func run(cmd *cobra.Command, args []string) {
	register(args, cmd)
}
