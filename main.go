package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rhodriguerrier/fun_size_url/cassandrasetup"
	"github.com/rhodriguerrier/fun_size_url/encoding"
)

func idHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	queryString := fmt.Sprintf(
		"SELECT original_url FROM test_keyspace.url_redirect WHERE url_hash='%s'",
		vars["url_id"],
	)
	log.Println(queryString)
	var origURL string
	iter := cassandrasetup.Session.Query(queryString).Iter()
	iter.Scan(&origURL)
	http.Redirect(w, r, origURL, http.StatusSeeOther)
}

func getNextId() uint64 {
	var maxId uint64
	iter := cassandrasetup.Session.Query("SELECT MAX(url_id) FROM test_keyspace.url_redirect").Iter()
	iter.Scan(&maxId)
	return maxId
}

func newUrlHandlerString(w http.ResponseWriter, r *http.Request) {
	var urlToShorten string
	err := json.NewDecoder(r.Body).Decode(&urlToShorten)
	if err != nil {
		log.Println(err)
		return
	}

	if !strings.HasPrefix(urlToShorten, "https://") && !strings.HasPrefix(urlToShorten, "http://") {
		urlToShorten = "https://" + urlToShorten
	}

	newUrlId := getNextId() + 1
	urlHash := encoding.Base62Encode(newUrlId)
	queryString := fmt.Sprintf(
		"INSERT INTO test_keyspace.url_redirect (url_hash, original_url, url_id) VALUES ('%s', '%s', %v)",
		urlHash,
		urlToShorten,
		newUrlId,
	)
	err = cassandrasetup.Session.Query(queryString).Exec()
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("URL shortened: ", urlToShorten)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(urlHash)
}

func handleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/{url_id}", idHandler)
	r.HandleFunc("/test/new-url-string", newUrlHandlerString).Methods("POST")
	staticFileDirectory := http.Dir("./frontend/")
	staticFileHandler := http.FileServer(staticFileDirectory)
	r.PathPrefix("/").Handler(staticFileHandler).Methods("GET")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), nil))
}

func main() {
	handleRequests()
}
