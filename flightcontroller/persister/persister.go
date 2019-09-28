package persister

import (
	. "github.com/SpectralHiss/spacextest/flightcontroller/querytypes"
	"database/sql"
	_ 	"github.com/lib/pq"
	"fmt"
	"time"
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
	query := fmt.Sprintf("INSERT INTO %s VALUES ('%s','%s','%s','%s','%s','%d','%s')",
	p.TableName, t.FirstName, t.LastName, t.Gender, t.Birthday.Format(time.RFC3339), 
	string(t.LaunchpadID), int(t.DestinationID), t.LaunchDate.Format(time.RFC3339))
	
	fmt.Printf(query)
	
	_, err = db.Exec(query)
	return err
}

// type Getter interface {
// 	GetAll() []TicketDetails
// }

