package flightcontroller

import (
	"github.com/SpectralHiss/spacextest/flightcontroller/persister"
	"github.com/SpectralHiss/spacextest/flightcontroller/ticketdetails"

	"github.com/SpectralHiss/spacextest/flightcontroller/scheduler"
	"github.com/SpectralHiss/spacextest/flightcontroller/spacexquerier"
)

type flightController struct {
	db    persister.Saver
	q     spacexquerier.SpaceXQuerier
	sched scheduler.Scheduler
}

func NewFlightController(db persister.Saver, q spacexquerier.SpaceXQuerier, sched scheduler.Scheduler) *flightController {

	// query spacex for launchpadIDs

	return &flightController{
		db:    db,
		q:     q,
		sched: sched,
	}
}

func (f *flightController) Reserve(ticketdetails.TicketDetails) error {
	return nil
}
