package main

import (
	"github/com/fcmdias/CSVAnalysis/services/backend/pkg/web"
	"log"
	"net/http"
)

func main() {
	http.Handle("/popular", web.EnableCORS(http.HandlerFunc(web.Popular)))
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
