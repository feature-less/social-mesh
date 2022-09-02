package repository

import (
	"database/sql"
	"log"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

const (
	create string = `INSERT INTO public.threads (poster_id, title, body) VALUES($1, $2, $3) RETURNING id, created_at;`
)

type Thread struct {
	ID        uuid.UUID      `json:"id" db:"id"`
	PosterID  uuid.UUID      `json:"poster_id" db:"poster_id"`
	CreatedAt sql.NullTime   `json:"created_at" db:"created_at"`
	UpdatedAt sql.NullTime   `json:"updated_at" db:"updated_at"`
	DeletedAt sql.NullTime   `json:"deleted_at" db:"deleted_at"`
	Title     sql.NullString `json:"title" db:"title"`
	body      sql.NullString `json:"body_name" db:"body_name"`
}

// creates new thread
func (thread *Thread) Create(db *sqlx.DB) (*Thread, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Printf("can't start transaction: %v", err)
	}
	row, err := tx.Query(create)
	row.Scan(&thread.ID, &thread.CreatedAt)
	if err != nil {
		log.Printf("can't execute transaction query: %v", err)
	}
	err = tx.Commit()
	thread.UpdatedAt = thread.CreatedAt
	return thread, nil
}

// update entire thread at once
func (thread *Thread) Update() {

}
