package models

import (
	"database/sql"
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

type Error struct {
	Message string `json:"message"`
}

func (e *Event) Save() Error {
	query := `
	INSERT INTO events (name, description, location, dateTime, user_id) 
	VALUES (?, ?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return Error{Message: "Failed to prepare statement: " + err.Error()}
	}
	defer stmt.Close()

	res, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return Error{Message: "Failed to execute statement: " + err.Error()}
	}

	id, err := res.LastInsertId()
	if err != nil {
		return Error{Message: "Failed to retrieve last insert ID: " + err.Error()}
	}
	e.ID = id

	return Error{}
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT id, name, description, location, dateTime, user_id FROM events`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var ev Event
		err := rows.Scan(&ev.ID, &ev.Name, &ev.Description, &ev.Location, &ev.DateTime, &ev.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, ev)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}

func GetById(ID int64) (Event, Error) {
	query := `SELECT id, name, description, location, dateTime, user_id FROM events WHERE id = ?`
	var e Event

	err := db.DB.QueryRow(query, ID).Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return Event{}, Error{Message: "Event not found"}
		}
		return Event{}, Error{Message: "Query error: " + err.Error()}
	}

	return e, Error{}
}

func (e Event) Update() (Event, Error) {
	query := `
	UPDATE events
	SET name = ?, description = ? , location = ? , dateTime = ? , user_id = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return Event{}, Error{Message: "Failed to prepare statement: " + err.Error()}
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID, e.ID)
	if err != nil {
		return Event{}, Error{Message: "Failed to execute statement: " + err.Error()}
	}

	return e, Error{}
}

func (e Event) Delete() Error {
	query := `
		DELETE FROM events
		WHERE events.id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return Error{Message: "Failed to prepare statement: " + err.Error()}
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ID)
	if err != nil {
		return Error{Message: "Failed to execute statement: " + err.Error()}
	}

	return Error{}
}
