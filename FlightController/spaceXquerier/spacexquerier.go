package spaceXquerier

import "time"

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . SpaceQuerier
type SpaceQuerier interface {
	LaunchPossible(launchID int, date time.Time) bool
}
