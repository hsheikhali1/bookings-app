package main

import (
	"bookings-app/cmd"
	"bookings-app/pkg/utils"
)

// Rundown of the project
/*
Conference app via CLI that allows users to:

1. Select a conference they want to attend
2. Register for the conference
3. When a user registers, it updates the reference to the tickets available to that conference


== other requirements
1. User struct
2. Conference struct
3. Slice of Conferences for all the conferences available
4. func to remove conference from slice once there are no more tickets available
	a. adds the unavailable conferences to a slice of Unavailable conferences
5. func to submit application/register
	a. register will just add a user to the slice of users, and update the tickets from reference of the conference selected
6. Generate the structs for the conferences based on an external file so we can persist the data after we end CLI execution


== OVERKILL stuff
1. Add concurrency
2. Make it a TUI program AND a CLI ? :)
*/

func main() {
		utils.ReadConfigFile()
		cmd.Execute()
}
