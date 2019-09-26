package persister

import "time"

type TicketDetails struct {
	FirstName     string
	LastName      string
	Gender        string
	Birthday      time.Time
	LaunchpadID   int
	DestinationID int
	LaunchDate    time.Time
}

type SaverGetter interface {
	Saver
	Getter
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . Saver
type Saver interface {
	Save(TicketDetails) error
}

type Getter interface {
	GetAll() []TicketDetails
}
