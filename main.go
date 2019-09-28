package main
import (
	"github.com/SpectralHiss/spacextest/flightcontroller/spacexquerier"
	"github.com/SpectralHiss/spacextest/flightcontroller/persister"
	"github.com/SpectralHiss/spacextest/flightcontroller/scheduler"
	"github.com/SpectralHiss/spacextest/flightcontroller"
)

func main() {
	// here is how it ties together

	http.Handle("/ticket", func(w http.ResponseWriter, r *http.Request){
		if !r.Method() == http.MethodPost {
			// return bad request
			w.WriteHeader(http.StatusBadRequest)
			
		}
		// skipping validation!

		someTicket := TicketDetails{
			FirstName : r.PostFormValue("firstName"),
			LastName: r.PostFormValue("lastName"),
			Gender: r.PostFormValue("gender"),
			Birthday: r.PostFormValue("birthday"),
			LaunchpadID: LaunchPadID(r.PostFormValue("switzerland")),
			DestinationID: r.PostForm("destination_id"),
			LaunchDate: r.PostForm("launch_date"),
		}
		//func NewFlightController(db persister.Saver, q spacexquerier.SpaceXQuerier, sched scheduler.FlightScheduler) *flightController {
		spaceX :=  &spacexquerier.SpaceXQueryImpl{ApiUrl:"https://api.spacexdata.com/"} 
		
		flightController := NewFlightController(&persister.Persister{
			ConnString: "postgres://postgres@localhost:8001/postgres?sslmode=disable",
			TableName: "tickets"
		}, spaceX, scheduler.NewScheduler(spaceX))

		//TODO make custom error type corresponding to failure modes
		err := flightController.Reserve(someTicket)
		
		if err != nil {
			// can use different status like 409 conflict
			w.WriteHeader(400)
		}
		// success!
		w.WriteHeader(200)

	})


	// go http.ListenAndServer(":8000", nil)
}