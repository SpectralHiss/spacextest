package ticketdetails

import "time"

type TicketDetails struct {
	FirstName     string
	LastName      string
	Gender        string
	Birthday      time.Time
	LaunchpadID   string
	DestinationID int
	LaunchDate    time.Time
}
