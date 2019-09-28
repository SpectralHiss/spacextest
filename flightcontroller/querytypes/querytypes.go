package querytypes

type Day int
type DestinationID int
type LaunchPadID string

type TicketDetails struct {
	FirstName     string
	LastName      string
	Gender        string
	Birthday      string // RFC.3339
	LaunchpadID   LaunchPadID
	DestinationID DestinationID
	LaunchDate    string // RFC.3339
}
