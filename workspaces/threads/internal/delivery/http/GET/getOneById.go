package get

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"threads/internal/repository"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

/*
import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"threads/internal/repository"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func getThreadByID(rw http.ResponseWriter, req *http.Request) {
	var db *sqlx.DB = req.Context().Value("db").(*sqlx.DB)
	var body io.ReadCloser = req.Body
	var thread = repository.Thread{}

	err := json.NewDecoder(body).Decode(&thread)
	if err != nil {
		log.Printf("an unexpected error has occured while decoding body: %v", err)
		return
	}
	if thread.ID.String() != "" {
		if _, err = uuid.Parse(thread.ID.String()); err == nil {
			thread.Create(db)
			rw.Header().Set("Content-Type", "application/json")
			rw.WriteHeader(200)
			return
		}
		rw.Header().Set("Content-Type", "text/plain")
		rw.Write([]byte("user_id was not provided"))
		rw.WriteHeader(422)
		return
	}
}*/

// getOneThreadByID
// @Summary retrieves a single thread using its id
// @Description get one thread by id
// @ID get-one-thread-by-id
// @Accept  json
// @Produce  json
// @Param   id      path   string     true  "ID"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} string "We need ID!!"
// @Failure 404 {object} string "Can not find ID"
// @Router /{id} [get]
func getOneThreadByID(rw http.ResponseWriter, req *http.Request) {
	var db *sqlx.DB = req.Context().Value("db").(*sqlx.DB)
	var body io.ReadCloser = req.Body
	var thread repository.Thread

	err := json.NewDecoder(body).Decode(&thread)

	if err != nil {
		log.Printf("an error has decoding json body: %v", err)
		return
	}

	if thread.ID.String() != "" {
		if _, err = uuid.Parse(thread.ID.String()); err == nil {
			thread.Create(db)
			rw.Header().Set("Content-Type", "application/json")
			rw.WriteHeader(200)
			return
		}
		rw.Header().Set("Content-Type", "text/plain")
		rw.Write([]byte("thread_id was not provided"))
		rw.WriteHeader(422)
		return
	}

}
