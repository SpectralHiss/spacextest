package persister

import (
	. "github.com/SpectralHiss/spacextest/flightcontroller/querytypes"
	
	
	"database/sql"
	
	_ "github.com/lib/pq"

	"fmt"

)

type SaverGetter interface {
	Saver
	//Getter
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . Saver
type Saver interface {
	Save(TicketDetails) error
}

type Persister struct {
	ConnString string
	TableName string
}

/*
type TicketDetails struct {
	FirstName     string
	LastName      string
	Gender        string
	Birthday      time.Time
	LaunchpadID   LaunchPadID
	DestinationID DestinationID
	LaunchDate    time.Time
}*/


func (p *Persister) Save (t TicketDetails) error {
	db, err := sql.Open("postgres",p.ConnString)
	if err != nil {
		return err
	}
	//TODO: avoid potential sqli through prepared statement
	query := fmt.Sprintf(`INSERT INTO %s(first_name,last_name,gender,birthday, launchpad_id, destination_id, launch_date) 
	VALUES ('%s','%s','%s','%s','%s','%d','%s')`,
	p.TableName, t.FirstName, t.LastName, t.Gender, t.Birthday, 
	string(t.LaunchpadID), int(t.DestinationID), t.LaunchDate)
	
	
	_, err = db.Exec(query)
	return err
}

// type Getter interface {
// 	GetAll() []TicketDetails
// }

