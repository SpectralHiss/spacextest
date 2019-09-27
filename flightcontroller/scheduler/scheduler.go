package scheduler

import (
	"math/rand"

	"github.com/SpectralHiss/spacextest/flightcontroller/spacexquerier"
)

//import "github.com/SpectralHiss/spacextest/flightcontroller/spacexquerier"

type Day int
type DestinationID int
type LaunchPadID int

type Schedule map[LaunchPadID]map[Day]DestinationID

type Scheduler struct {
	SMap Schedule
}

func GenerateMapping(querier spacexquerier.SpaceXQuerier) Schedule {

	LaunchPadIDs := querier.GetLaunchIds()

	scheduleMap := make(map[LaunchPadID]map[Day]DestinationID)

	for id := range LaunchPadIDs {
		launchPadSched := make(map[Day]DestinationID)
		scheduleMap[LaunchPadID(id)] = launchPadSched
		take7from8 := rand.Perm(8)[:7]
		for idx, val := range take7from8 {
			launchPadSched[Day(idx)] = DestinationID(val)
		}

	}

	return scheduleMap

}

func NewScheduler(spsxq spacexquerier.SpaceXQuerier) *Scheduler {
	scheduler := &Scheduler{
		SMap: GenerateMapping(spsxq),
	}
	return scheduler
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . FlightScheduler
type FlightScheduler interface {
	CheckSchedule(launchID LaunchPadID, d Day, dest DestinationID) bool
}

func (sched *Scheduler) CheckSchedule(lid LaunchPadID, dw Day, dest DestinationID) bool {

	launchPadSchedule, ok := sched.SMap[lid]
	if !ok {
		return false
	}

	return launchPadSchedule[dw] == dest
}
