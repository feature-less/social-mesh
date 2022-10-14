package repository

import (
	"database/sql"
	"log"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

const (
	create string = `INSERT INTO public.threads (poster_id, title, body, links, images) VALUES($1, $2, $3, $4; $5) RETURNING id, created_at;`
	delete string = `DELETE FROM public.threads WHERE id = $1 AND poster_id = $2`
)

type Thread struct {
	ID        uuid.UUID        `json:"id" db:"id"`
	PosterID  uuid.UUID        `json:"poster_id" db:"poster_id"`
	CreatedAt sql.NullTime     `json:"created_at" db:"created_at"`
	UpdatedAt sql.NullTime     `json:"updated_at" db:"updated_at"`
	DeletedAt sql.NullTime     `json:"deleted_at" db:"deleted_at"`
	Title     sql.NullString   `json:"title" db:"title"`
	Body      sql.NullString   `json:"body_name" db:"body_name"`
	Links     []sql.NullString `json:"links" db:"links"`
	Images    []sql.NullString `json:"images" db:"images"`
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

// delete thread
func (thread *Thread) Delete(db *sqlx.DB) (*Thread, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Printf("can't start trnsaction", err)
	}
	row, err := tx.Query(delete)
	row.Scan(&thread.ID, &thread.DeletedAt)
	if err != nil {
		log.Printf("can't execute transaction query: %v", err)
	}
	err = tx.Commit()
	thread.UpdatedAt = thread.DeletedAt
	return thread, nil
}
