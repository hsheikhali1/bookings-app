package utils

import "fmt"

type User struct {
	firstName        string
	lastName         string
	email            string
	ticketsPurchased uint
}

type Conference struct {
	ConferenceName     string
	ConferenceLocation string
	AvailableTickets   uint
}

type Conferences struct {
	conferences []Conference
}

// generates a new slice of conference by reading a json file :)
func (c *Conferences) GenerateConferences() []Conference {
	jsonFile := *ReadConfigFile()

	for _, obj := range jsonFile {
		c.conferences = append(c.conferences, Conference{
			ConferenceName:     obj.ConferenceName,
			ConferenceLocation: obj.ConferenceLocation,
			AvailableTickets:   obj.AvailableTickets,
		})
	}

	return c.conferences
}

func (c *Conferences) PurchaseTickets(name string, amount uint) {
	fmt.Println("Attempting to purchase some tickets")
	var newRecord JsonSructure
	for _, conference := range c.conferences {
		if conference.ConferenceName == name {
			//TODO: fix PurchaseTickets logic
			if conference.AvailableTickets > 0 {

				conference.AvailableTickets = conference.AvailableTickets - amount
				newRecord = JsonSructure{
					ConferenceName:     conference.ConferenceName,
					ConferenceLocation: conference.ConferenceLocation,
					AvailableTickets:   conference.AvailableTickets,
				}
			} else {
				fmt.Println("Conference has sold out! Sorry!")
			}
		}
	}

	err := WriteToFile(&newRecord)

	if err != nil {
		fmt.Println("Failed to purcahse tickets")
	}
}
