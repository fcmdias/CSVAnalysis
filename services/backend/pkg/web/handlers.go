package web

import (
	"encoding/json"
	"github/com/fcmdias/CSVAnalysis/services/backend/pkg/db/vehicles"
	"net/http"
)

func PopularHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	// Get the sort order and filter from the query string
	sortOrder := r.URL.Query().Get("sort")
	filter := r.URL.Query().Get("filter")

	data := vehicles.FetchVehicles(filter)

	dataPopular := vehicles.Popularity(data, sortOrder)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dataPopular)
}

func ByYearHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	// Get the sort order and filter from the query string
	filter := r.URL.Query().Get("filter")

	data := vehicles.FetchVehicles(filter)

	dataByYear := vehicles.ByYear(data)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dataByYear)
}
