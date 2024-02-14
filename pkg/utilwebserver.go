package webmontag

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// setUp the db connection
func setUp() (*sql.DB, func()) {
	godotenv.Load(".env")
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is empty")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can not connect to DB", err)
	}

	// Return connection and cleanup function
	return conn, func() {
		err := conn.Close()
		if err != nil {
			log.Fatal("Can not close DB", err)
		}
	}
}

// json responsewriter
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to Marshal JSON %v \n", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}

// secretkey checker
func keyCheck(r *http.Request) bool {

	authorization := r.Header.Get("Authorization")
	key := strings.TrimPrefix(authorization, "key ")
	return key == Webserveroptions.Key

}
