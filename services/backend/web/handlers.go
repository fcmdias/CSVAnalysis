package web

import (
	_ "github/com/fcmdias/CSVAnalysis/services/backend/docs"
	"log/slog"

	"github/com/fcmdias/CSVAnalysis/services/backend/db/vehicles"
	"net/http"
)

// PopularHandler godoc
// @Summary Get popular vehicles
// @Description Fetches a list of popular vehicles, sorted and filtered as per query parameters.
// @Tags vehicles
// @Accept  json
// @Produce  json
// @Param sort query string false "Sort order (asc or desc)"
// @Param filter query string false "Filter (all, hybrid, or electric)"
// @Success 200 {array} models.VehiclePopularity "A list of vehicles"
// @Failure 400 {object} string "Invalid sort order or filter"
// @Failure 404 {object} string "Method not supported"
// @Failure 500 {object} string "Internal Server Error"
// @Router /popular [get]
func PopularHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	// Get the sort order and filter from the query string
	sortOrder := r.URL.Query().Get("sort")
	filter := r.URL.Query().Get("filter")

	if !isValidSortOrder(sortOrder) {
		slog.Error("Error invalid filter", "filter", filter)
		errorResponse := ErrorResponse{Error: "invalid filter"}
		sendResponse(w, http.StatusInternalServerError, nil, &errorResponse)
		return
	}

	if !isValidFilter(filter) {
		http.Error(w, "Invalid filter", http.StatusBadRequest)
		return
	}

	data, err := vehicles.FetchVehicles(filter)
	if err != nil {
		errorResponse := ErrorResponse{Error: "Internal Server Error"}
		sendResponse(w, http.StatusInternalServerError, nil, &errorResponse)
		return
	}

	dataPopular, err := vehicles.Popularity(data, sortOrder)
	if err != nil {
		errorResponse := ErrorResponse{Error: "Internal Server Error"}
		sendResponse(w, http.StatusInternalServerError, nil, &errorResponse)
		return
	}

	sendResponse(w, http.StatusOK, dataPopular, nil)
}

// ByYearHandler godoc
// @Summary Get vehicles by year
// @Description Fetches a list of vehicles filtered as per the query parameter.
// @Tags vehicles
// @Accept  json
// @Produce  json
// @Param filter query string false "Filter (all, hybrid, or electric)"
// @Success 200 {array} models.VehicleByYear "A list of vehicles sorted by year"
// @Failure 400 {object} ErrorResponse "Invalid filter - The filter query parameter is required and must be one of 'all', 'hybrid', or 'electric'. Response body will contain: {'error': 'Invalid filter'}"
// @Failure 404 {object} string "Method not supported"
// @Failure 500 {object} string "Internal Server Error"
// @Router /byyear [get]
func ByYearHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	filter := r.URL.Query().Get("filter")

	if !isValidFilter(filter) {
		slog.Error("Error invalid filter", "filter", filter)
		errorResponse := ErrorResponse{Error: "invalid filter"}
		sendResponse(w, http.StatusInternalServerError, nil, &errorResponse)
		return
	}

	data, err := vehicles.FetchVehicles(filter)
	if err != nil {
		slog.Error("Error fetching vehicles", "error", err.Error())
		errorResponse := ErrorResponse{Error: "Internal Server Error"}
		sendResponse(w, http.StatusInternalServerError, nil, &errorResponse)
		return
	}

	dataByYear := vehicles.ByYear(data)

	sendResponse(w, http.StatusOK, dataByYear, nil)
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
