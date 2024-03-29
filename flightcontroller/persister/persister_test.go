package persister_test

import (

	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"database/sql"

	_ "github.com/lib/pq"


	. "github.com/SpectralHiss/spacextest/flightcontroller/persister"
	. "github.com/SpectralHiss/spacextest/flightcontroller/querytypes"

)

func TestPersister(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controller suite")
}

var _ = Describe("Persister",func(){
	// testing using localhost and alternate table
	var db *sql.DB
	var connString = "postgres://postgres@localhost:8001/postgres?sslmode=disable"
	
	AfterSuite(func(){db.Close()})
	BeforeSuite(func(){
		var err error
		db, err = sql.Open("postgres", connString)
		Expect(err).To(BeNil())
	})
	

	var saver *Persister
	
	Context("when save persister is invoked with a valid ticketrequest",func(){

		BeforeEach(func(){
			//TODO: make birthday a date
			_, err := db.Exec(`CREATE TABLE test_tickets (
				first_name varchar (20),
				last_name varchar (20),
				gender varchar (10),
				birthday timestamptz,
				launchpad_id varchar (50),
				destination_id SMALLINT,
				launch_date timestamptz)`)
				Expect(err).To(BeNil())

			saver = &Persister{ConnString: connString, TableName: "test_tickets"}
		})

		AfterEach(func(){
			db.Exec(`DROP TABLE test_tickets`)
		})

		It("The ticket should be added to the database",func(){
			someTicket := TicketDetails{
				FirstName : "Joe",
				LastName: "Goodman",
				Gender: "Male",
				Birthday: time.Now().UTC().Format(time.RFC3339),
				LaunchpadID: LaunchPadID("switzerland"),
				DestinationID: DestinationID(4),
				LaunchDate: time.Now().UTC().Format(time.RFC3339),
			}


			err := saver.Save(someTicket)

			Expect(err).To(BeNil())

			rows, err := db.Query(`SELECT * from test_tickets WHERE first_name ='Joe' AND 
			launchpad_id = 'switzerland' AND destination_id = 4`)

			Expect(err).To(BeNil())

			defer rows.Close()
			
			var ticket TicketDetails
			for rows.Next() {
				// should only be 1
				err = rows.Scan(&ticket.FirstName, &ticket.LastName, &ticket.Gender,
					 &ticket.Birthday, &ticket.LaunchpadID, &ticket.DestinationID, &ticket.LaunchDate)
			}
			Expect(err).To(BeNil())
			Expect(ticket).To(Equal(someTicket))
			
		})
	})
})
