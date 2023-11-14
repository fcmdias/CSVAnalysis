package web

import (
	"encoding/json"
	"github/com/fcmdias/CSVAnalysis/services/backend/pkg/db/vehicles"
	"net/http"
)

func Popular(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	// Get the sort order from the query string
	sortOrder := r.URL.Query().Get("sort")

	data := vehicles.FetchVehicles()

	dataPopular := vehicles.Popularity(data, sortOrder)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dataPopular)
}
