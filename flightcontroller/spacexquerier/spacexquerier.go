package spacexquerier

import (
	"encoding/json"

	"log"
	"net/http"
	"net/url"
	

	"time"

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

	u, err := url.Parse(sq.ApiUrl)

	if err != nil {
		log.Fatal(err)
	}
	u.Path += "/v3/launches/upcoming"
	Response, err := http.Get(u.String())
	if err != nil || Response.StatusCode != 200 {
		return false
	}

	dec := json.NewDecoder(Response.Body)

	type UpcomingLaunches struct {
		TimeUTC string `json:"launch_date_utc"`
		LaunchSite struct {
			SiteID LaunchPadID `json:"site_id"`
		} `json:"launch_site"`
	}


	var launch_dets []UpcomingLaunches
	for dec.More() {
		err = dec.Decode(&launch_dets)

		if err != nil {
		return false
		}
	}

	for _, launch := range launch_dets {
		if sameDayS1(launch.TimeUTC, date) && launch.LaunchSite.SiteID == launchID {
			return false
		}
	}
	return true
}


func sameDayS1(d1 string, d2 time.Time) bool {
	d1Parsed , err := time.Parse(time.RFC3339, d1)
	if err != nil {
		return false
	}
	d1y,d1m,d1d := d1Parsed.Date()
	d2y,d2m,d2d := d2.Date()
	return d1y == d2y && d1m == d2m && d1d == d2d 
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
