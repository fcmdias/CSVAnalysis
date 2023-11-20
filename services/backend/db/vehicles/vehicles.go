package vehicles

import (
	"encoding/csv"
	"github/com/fcmdias/CSVAnalysis/services/backend/models"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"time"
)

// FetchVehicles reads and filters vehicle data from a CSV file.
func FetchVehicles(filter string) (vehicles []models.VehicleData, err error) {
	start := time.Now()

	f, err := os.Open("Electric_Vehicle_Population_Data.csv")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = ','
	r.Comment = '#'
	r.LazyQuotes = true

	for {
		record, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		vehicle := parseRecordToVehicle(record)

		if shouldIncludeVehicle(vehicle, filter) {
			vehicles = append(vehicles, vehicle)
		}
	}

	log.Printf("Time to read and process CSV file: %v\n", time.Since(start))
	return vehicles, nil
}

// parseRecordToVehicle converts a CSV record to a VehicleData object.
func parseRecordToVehicle(record []string) models.VehicleData {
	return models.VehicleData{
		VIN:             record[0],
		County:          record[1],
		City:            record[2],
		State:           record[3],
		PostalCode:      record[4],
		ModelYear:       record[5],
		Make:            record[6],
		Model:           record[7],
		EVType:          record[8],
		CAFVEligibility: record[9],
		ElectricRange:   record[10],
		BaseMSRP:        record[11],
		LegislativeDist: record[12],
		DOLVehicleID:    record[13],
		VehicleLocation: record[14],
		ElectricUtility: record[15],
		CensusTract:     record[16],
	}
}

// shouldIncludeVehicle checks if a vehicle matches the given filter.
func shouldIncludeVehicle(vehicle models.VehicleData, filter string) bool {
	switch filter {
	case "electric":
		return vehicle.EVType == "Battery Electric Vehicle (BEV)" || vehicle.EVType == "Electric Vehicle Type"
	case "hybrid":
		return vehicle.EVType == "Plug-in Hybrid Electric Vehicle (PHEV)"
	case "all":
		return true
	default:
		return false
	}
}

// Popularity computes the popularity and sorts the slice based on the sort parameter.
func Popularity(vehicles []models.VehicleData, sortOrder string) ([]models.VehiclePopularity, error) {
	start := time.Now()

	popularity := make(map[models.VehiclePopularity]int)
	for _, vehicle := range vehicles {
		popularity[models.VehiclePopularity{
			Make:  vehicle.Make,
			Model: vehicle.Model,
			Total: 0,
		}]++
	}

	var popularityVehicles []models.VehiclePopularity
	for vehicle, total := range popularity {
		popularityVehicles = append(popularityVehicles, models.VehiclePopularity{
			Make:  vehicle.Make,
			Model: vehicle.Model,
			Total: total,
		})
	}

	popularityVehicles = topVehiclePopularity20(popularityVehicles)
	popularityVehicles = sortByTotalVehiclePopularity(popularityVehicles, sortOrder)

	log.Printf("Time taken to compute popularity: %v\n", time.Since(start))
	return popularityVehicles, nil
}

func sortByTotalVehiclePopularity(data []models.VehiclePopularity, sortOrder string) []models.VehiclePopularity {
	// Sorting based on the sortOrder parameter
	if sortOrder == "asc" {
		sort.Slice(data, func(i, j int) bool {
			return data[i].Total < data[j].Total // Ascending order
		})
	} else { // Default to descending order if no sort order is specified or if it's "desc"
		// already sorted
	}

	return data
}

func topVehiclePopularity20(data []models.VehiclePopularity) []models.VehiclePopularity {

	if len(data) <= 20 {
		return data
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].Total > data[j].Total // Descending order
	})

	return data[:20]

}

func ByYear(data []models.VehicleData) []models.VehicleByYear {

	thisYear := time.Now().Year()
	res := make([]models.VehicleByYear, 10)

	for i := 0; i < 10; i++ {
		res[i].Year = thisYear - 9 + i
	}

	for _, vehicle := range data {
		year, err := strconv.Atoi(vehicle.ModelYear)
		if err != nil {
			continue
		}
		if year < thisYear-9 || year > thisYear {
			continue
		}
		res[9+year-thisYear].Total++
	}

	return res
}
