package spacexquerier

import (
	"encoding/json"

	"log"
	"net/http"
	"net/url"
	"time"

	//"fmt"


	"github.com/SpectralHiss/spacextest/flightcontroller/querytypes"
	. "github.com/SpectralHiss/spacextest/flightcontroller/querytypes"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . SpaceXQuerier
type SpaceXQuerier interface {
	LaunchPossible(launchID LaunchPadID, date time.Time) bool
	GetLaunchIds() []querytypes.LaunchPadID
}

type SpaceXQueryImpl struct {
	ApiUrl string
}

func (sq*SpaceXQueryImpl) LaunchPossible(launchID LaunchPadID, date time.Time) bool {
	return true
}

func (sq *SpaceXQueryImpl) GetLaunchIds() []LaunchPadID {
	u, err := url.Parse(sq.ApiUrl)

	if err != nil {
		log.Fatal(err)
	}
	u.Path += "/v3/landpads"
	Response, err := http.Get(u.String())
	if err != nil || Response.StatusCode != 200 {
		return []LaunchPadID{}
	}

	dec := json.NewDecoder(Response.Body)

	type LaunchDetails struct {
		SiteID LaunchPadID `json:"site_id"`
	}

	var lids []LaunchDetails
	for dec.More() {
		err = dec.Decode(&lids)

		if err != nil {
			return []LaunchPadID{}
		}
	}
	
	collectLaunchIDs := []LaunchPadID{}
	for _, elem := range lids {
		collectLaunchIDs = append(collectLaunchIDs, elem.SiteID)
	}

	return collectLaunchIDs
}
