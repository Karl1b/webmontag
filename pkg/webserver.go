package webmontag

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"

	"github.com/karl1b/webmontag/pkg/internal/database"
	_ "github.com/lib/pq"
)

func Startwebserver(command string) {
	fmt.Println("Startwebserver")

	Webserveroptions.isWithDb = false
	if command == "db" {
		Webserveroptions.isWithDb = true
		fmt.Println("DB enabled.")
	}

	portString := Webserveroptions.Port

	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*", "*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	if Webserveroptions.isWithDb {
		conn, cleanup := setUp()
		defer cleanup()
		queries := database.New(conn)
		r.Post("/dbscrapepages", (dbscrapepages(queries)))
		r.Get("/dbseepages", (dbseepages(queries)))
	}
	r.Get("/", (handlerReadiness))
	r.Post("/scrapepages", (scrapepages)) // post a page to scrape

	srv := &http.Server{
		Handler: r,
		Addr:    ":" + portString,
	}

	log.Printf("runs on: %s", portString)
	srv.ListenAndServe()

}

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}

func scrapepages(w http.ResponseWriter, r *http.Request) {

	type Response struct {
		Info string `json:"info"`
	}

	isKeyValid := keyCheck(r)
	if !(isKeyValid) {
		respondWithJSON(w, 400, Response{Info: "KEY INVALID"})
		return
	}

	type RequestBody struct {
		PageUrl []string `json:"urls"`
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var reqBody RequestBody
	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	go ConcurrentScraper(reqBody.PageUrl, nil)

	response := Response{Info: "Key Korrekt: Your input is beein scraped Scraped"}

	respondWithJSON(w, 200, response)
}

func dbscrapepages(queries *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		type Response struct {
			Info string `json:"info"`
		}

		isKeyValid := keyCheck(r)
		if !(isKeyValid) {
			respondWithJSON(w, 400, Response{Info: "KEY INVALID"})
			return
		}

		type RequestBody struct {
			PageUrl []string `json:"urls"`
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		var reqBody RequestBody
		err = json.Unmarshal(body, &reqBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		go ConcurrentScraper(reqBody.PageUrl, queries)

		respondWithJSON(w, 200, Response{Info: "The requested urls are beein scraped!"})

	}
}

func dbseepages(queries *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		type Response struct {
			Info string `json:"info"`
		}

		isKeyValid := keyCheck(r)
		if !(isKeyValid) {
			respondWithJSON(w, 400, Response{Info: "KEY INVALID"})
			return
		}

		results, err := queries.GetAllPages(context.Background())
		if err != nil {
			fmt.Println(err)
			respondWithJSON(w, 500, Response{Info: "Error getting results from database"})
			return
		}

		type GoodResponse struct {
			PageUrl []pageDetails `json:"pages"`
		}

		var goodResponses []pageDetails

		for _, page := range results {

			var newPageDetails pageDetails

			newPageDetails.PageDomain = page.Domain.String
			newPageDetails.SearchDeep = int(page.Searchdeep.Int32)
			newPageDetails.MaxSearchDeep = int(page.Maxsearchdeep.Int32)

			links, err := queries.GetLinksByPageID(context.Background(), sql.NullInt32{Int32: page.ID, Valid: true})
			if err != nil {
				continue
			}
			for _, link := range links {
				newPageDetails.ExternalLinks = append(newPageDetails.ExternalLinks, link.Domain.String)
			}

			goodResponses = append(goodResponses, newPageDetails)

		}

		respondWithJSON(w, 200, GoodResponse{PageUrl: goodResponses})
	}
}
