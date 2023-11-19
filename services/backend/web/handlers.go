package web

import (
	"encoding/json"
	"github/com/fcmdias/CSVAnalysis/services/backend/db/vehicles"
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

	if !isValidSortOrder(sortOrder) {
		http.Error(w, "Invalid sort order", http.StatusBadRequest)
		return
	}

	if !isValidFilter(filter) {
		http.Error(w, "Invalid filter", http.StatusBadRequest)
		return
	}

	data, err := vehicles.FetchVehicles(filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dataPopular, err := vehicles.Popularity(data, sortOrder)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dataPopular); err != nil {
		http.Error(w, "Failed to encode data", http.StatusInternalServerError)
	}
}

func ByYearHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	filter := r.URL.Query().Get("filter")

	if !isValidFilter(filter) {
		http.Error(w, "Invalid filter", http.StatusBadRequest)
		return
	}

	data, err := vehicles.FetchVehicles(filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dataByYear := vehicles.ByYear(data)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dataByYear); err != nil {
		http.Error(w, "Failed to encode data", http.StatusInternalServerError)
	}
}

// =========================================================================
// Helper functions to validate sortOrder and filter
func isValidSortOrder(sortOrder string) bool {
	if sortOrder == "asc" || sortOrder == "desc" {
		return true
	}
	return false
}

func isValidFilter(filter string) bool {
	if filter == "all" || filter == "hybrid" || filter == "electric" {
		return true
	}
	return false
}
