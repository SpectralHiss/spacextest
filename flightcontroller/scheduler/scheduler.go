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

type Scheduler struct {
	Spcxq spacexquerier.SpaceXQuerier
	SMap  Schedule
}

func GenerateMapping() Schedule {
	return make(Schedule)
}

func NewScheduler(spsxq spacexquerier.SpaceXQuerier) *Scheduler {
	scheduler := &Scheduler{
		Spcxq: &spacexquerierfakes.FakeSpaceXQuerier{},
		SMap:  GenerateMapping(),
	}
	return scheduler
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . FlightScheduler
type FlightScheduler interface {
	CheckSchedule(launchID LaunchPadID, d Day, dest DestinationID) bool
}

func (sched *Scheduler) CheckSchedule(lid LaunchPadID, dw Day, dest DestinationID) bool {
	return true
}
