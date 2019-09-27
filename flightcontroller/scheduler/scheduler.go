package scheduler

import (
	"github.com/SpectralHiss/spacextest/flightcontroller/spacexquerier"
	"github.com/SpectralHiss/spacextest/flightcontroller/spacexquerier/spacexquerierfakes"
)

//import "github.com/SpectralHiss/spacextest/flightcontroller/spacexquerier"

type Day int
type DestinationID int
type LaunchPadID int

type Schedule map[LaunchPadID]map[Day]DestinationID

type scheduler struct {
	spcxq spacexquerier.SpaceXQuerier
	s     Schedule
}

func GenerateMapping() Schedule {
	return make(Schedule)
}

func NewScheduler(spsxq spacexquerier.SpaceXQuerier) *scheduler {
	scheduler := &scheduler{
		spcxq: &spacexquerierfakes.FakeSpaceXQuerier{},
		s:     GenerateMapping(),
	}
	return scheduler
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . FlightScheduler
type FlightScheduler interface {
	CheckSchedule(launchID LaunchPadID, d Day, dest DestinationID) bool
}

// (sched *Schedule) func CheckSchedule(launchID , Day, DestinationID) {
// }
