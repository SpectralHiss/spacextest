package persister

import (
	. "github.com/SpectralHiss/spacextest/flightcontroller/querytypes"
)

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
