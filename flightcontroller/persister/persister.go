package persister

import (
	. "github.com/SpectralHiss/spacextest/flightcontroller/querytypes"
)

type SaverGetter interface {
	Saver
	//Getter
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . Saver
type Saver interface {
	Save(TicketDetails) error
}

type Persister struct {
	connString string
}

/*
type TicketDetails struct {
	FirstName     string
	LastName      string
	Gender        string
	Birthday      time.Time
	LaunchpadID   LaunchPadID
	DestinationID DestinationID
	LaunchDate    time.Time
}*/


func (*Persister) Save (TicketDetails) error {
	
	return nil
}

// type Getter interface {
// 	GetAll() []TicketDetails
// }

