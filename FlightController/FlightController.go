package FlightController

//TODO raise domain knoweledgfe?
type map[int]map[int]int DestinationMapping

type flightController struct {
	db                 *persister
	q                  *spacexquerier
	dmap DestinationMapping
}

var DefaultMapping = DestinationMapping{

}
}

var DefaultDestinationMapping DestinationMapping{
	0: //Monday DestinationIDs fixed, LocationIds is not.. this needs to be augmented with spaceX data.. needs refactor
}

func NewFlightController(db *persister, q *spacexquerier, destinationMapping destinationMapping) FlightController {
	return &FlightController{
		db: db,
		q:  q,
		dmap: destinationMapping
	}
}

func Reserve()
