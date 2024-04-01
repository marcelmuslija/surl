package short

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"surl-server/pkg/crypto"

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
	url := struct {
		Url string
	}{}

	if err := json.NewDecoder(r.Body).Decode(&url); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var shortCode string
	for {
		shortCode, _ = crypto.GenerateShortCode()

		var codeExists bool
		existsQuery := "SELECT EXISTS(SELECT 1 FROM urls WHERE code = $1)"
		err := u.db.QueryRow(existsQuery, shortCode).Scan(&codeExists)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if !codeExists {
			insertQuery := "INSERT INTO urls(code, url) VALUES($1, $2)"
			_, err := u.db.Exec(insertQuery, shortCode, url.Url)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			fmt.Fprintf(w, "http://%s/api/short/%s\n", r.Host, shortCode)
			return
		}
	}
}

func (u *ShortUrlHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortCode := vars["code"]

	var url string
	selectQuery := "SELECT url FROM urls WHERE code = $1"
	err := u.db.QueryRow(selectQuery, shortCode).Scan(&url)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Short URL is not valid", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	http.Redirect(w, r, url, http.StatusSeeOther)
}
