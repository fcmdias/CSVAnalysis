package main

import (
	"github/com/fcmdias/CSVAnalysis/services/backend/pkg/web"
	"log"
	"net/http"
)

func main() {
	http.Handle("/popular", web.EnableCORSMiddleware(http.HandlerFunc(web.PopularHandler)))
	http.Handle("/byyear", web.EnableCORSMiddleware(http.HandlerFunc(web.ByYearHandler)))
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
