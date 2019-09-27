package spacexquerier

import "time"

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . SpaceXQuerier
type SpaceXQuerier interface {
	LaunchPossible(launchID int, date time.Time) bool
	GetLaunchIds() []int
}
