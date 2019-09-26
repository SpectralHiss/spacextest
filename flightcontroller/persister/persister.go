package persister

import "github.com/SpectralHiss/spacextest/flightcontroller/ticketdetails"

type SaverGetter interface {
	Saver
	Getter
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . Saver
type Saver interface {
	Save(ticketdetails.TicketDetails) error
}

type Getter interface {
	GetAll() []ticketdetails.TicketDetails
}
