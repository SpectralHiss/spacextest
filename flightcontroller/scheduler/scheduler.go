package scheduler

//import "github.com/SpectralHiss/spacextest/flightcontroller/spacexquerier"

type Day int
type DestinationID int
type LaunchPadID int

type Schedule map[int]map[Day]DestinationID

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . FlightScheduler
type FlightScheduler interface {
	CheckSchedule(launchID LaunchPadID, d Day, dest DestinationID) bool
}
