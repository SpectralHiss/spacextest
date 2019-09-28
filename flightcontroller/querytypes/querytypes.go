package querytypes

import "time"

type Day int
type DestinationID int
type LaunchPadID string

type TicketDetails struct {
	FirstName     string
	LastName      string
	Gender        string
	Birthday      time.Time
	LaunchpadID   LaunchPadID
	DestinationID DestinationID
	LaunchDate    time.Time
}
