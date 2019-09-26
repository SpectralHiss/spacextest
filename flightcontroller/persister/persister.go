package persister

import "github.com/SpectralHiss/spacextest/flightcontroller"

type SaverGetter interface {
	Saver
	Getter
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . Saver
type Saver interface {
	Save(flightcontroller.TicketDetails) error
}

type Getter interface {
	GetAll() []flightcontroller.TicketDetails
}
