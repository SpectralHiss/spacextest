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

func (f *flightController) Reserve(t ticketdetails.TicketDetails) error {

	schedule := f.sched.GenerateSchedule()

	// TODO: refactor day mapping from date to our date
	day, ok := schedule[t.LaunchpadID][scheduler.Day(t.LaunchDate.Weekday())]

	if !ok || day != scheduler.DestinationID(t.DestinationID) {
		return fmt.Errorf("Bad Request, not conforming to schedule")
	}

	if f.q.LaunchPossible(t.LaunchpadID, t.LaunchDate) {
		return fmt.Errorf("Bad request conflicting with spaceX launch")
	}

	_ = f.db.Save(t)

	return nil
}
