package spacexquerier

import (
	"time"

	"github.com/SpectralHiss/spacextest/flightcontroller/querytypes"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . SpaceXQuerier
type SpaceXQuerier interface {
	LaunchPossible(launchID string, date time.Time) bool
	GetLaunchIds() []querytypes.LaunchPadID
}
