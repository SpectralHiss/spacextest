package flightcontroller

import (
	"fmt"

	"github.com/SpectralHiss/spacextest/flightcontroller/persister"
	"github.com/SpectralHiss/spacextest/flightcontroller/ticketdetails"

	"github.com/SpectralHiss/spacextest/flightcontroller/scheduler"
	"github.com/SpectralHiss/spacextest/flightcontroller/spacexquerier"
)

type flightController struct {
	db    persister.Saver
	q     spacexquerier.SpaceXQuerier
	sched scheduler.FlightScheduler
}

func NewFlightController(db persister.Saver, q spacexquerier.SpaceXQuerier, sched scheduler.FlightScheduler) *flightController {

	// query spacex for launchpadIDs

	return &flightController{
		db:    db,
		q:     q,
		sched: sched,
	}
}

func (f *flightController) Reserve(t ticketdetails.TicketDetails) error {

	_ = f.sched.CheckSchedule(scheduler.LaunchPadID(t.LaunchpadID), scheduler.Day(t.LaunchDate.Weekday()),
		scheduler.DestinationID(t.DestinationID))

	// // TODO: refactor day mapping from date to our date
	// day, ok := schedule[t.LaunchpadID][scheduler.Day(t.LaunchDate.Weekday())]
	// || day != scheduler.DestinationID(t.DestinationID)

	// if !ok {
	// 	return fmt.Errorf("Bad Request, not conforming to schedule")
	// }

	if !f.q.LaunchPossible(t.LaunchpadID, t.LaunchDate) {
		return fmt.Errorf("Bad request conflicting with spaceX launch")
	}

	return f.db.Save(t)
}
