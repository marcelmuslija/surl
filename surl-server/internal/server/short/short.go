package short

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type ShortUrlHandler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *ShortUrlHandler {
	return &ShortUrlHandler{
		db,
	}
}

func (u *ShortUrlHandler) Shorten(w http.ResponseWriter, r *http.Request) {
	fmt.Println("shorten")
}

func (u *ShortUrlHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	token := vars["token"]
	fmt.Println(token)
}

func (u *ShortUrlHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	token := vars["token"]
	fmt.Println(token)
}
