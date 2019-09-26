package scheduler

//import "github.com/SpectralHiss/spacextest/flightcontroller/spacexquerier"

type Day int
type DestinationID int

type Schedule map[int]map[Day]DestinationID

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . Scheduler
type Scheduler interface {
	GenerateSchedule() map[int]map[Day]DestinationID
}
