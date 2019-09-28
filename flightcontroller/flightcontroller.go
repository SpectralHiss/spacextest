package flightcontroller

import (
	"fmt"
	"time"

	"github.com/SpectralHiss/spacextest/flightcontroller/persister"

	. "github.com/SpectralHiss/spacextest/flightcontroller/querytypes"
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

func (f *flightController) Reserve(t TicketDetails) error {
	
	goDateLaunch,_ := time.Parse(time.RFC3339, t.LaunchDate)
	ok := f.sched.CheckSchedule(LaunchPadID(t.LaunchpadID), Day(goDateLaunch.Weekday()),
		DestinationID(t.DestinationID))

	if !ok {
		return fmt.Errorf("Bad request, not conforming to schedule")
	}

	if !f.q.LaunchPossible(t.LaunchpadID,goDateLaunch) {
		return fmt.Errorf("Bad request conflicting with spaceX launch")
	}

	return f.db.Save(t)
}
