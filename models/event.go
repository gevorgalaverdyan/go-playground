package models

import (
	"time"

	"github.com/gevorgalaverdyan/go-playground/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

type Error struct{
	Message string
}

func (e Event) Save() Error{
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id) 
	VALUES 
	(?,?,?,?,?)
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return Error{Message: "Prepare Nuke"}
	}

	defer stmt.Close()

	res, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	if err != nil {
		return Error{Message: "exec nuke"}
	}

	id, err := res.LastInsertId()
	e.ID = id
	return Error{}
}

func GetAllEvents() ([]Event, error){
	query := `SELECT * FROM events`

	res, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer res.Close()

	var events []Event

	for res.Next() {
		var ev Event
		err := res.Scan(&ev.ID, &ev.Name, &ev.Description, &ev.Location, &ev.DateTime, &ev.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, ev)
	}

	return events, nil
}
