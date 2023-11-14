package vehicles

import (
	"encoding/csv"
	"fmt"
	"github/com/fcmdias/CSVAnalysis/services/backend/pkg/models"
	"io"
	"log"
	"os"
	"sort"
	"time"
)

func FetchVehicles() (vehicles []models.VehicleData) {
	f, err := os.Open("Electric_Vehicle_Population_Data.csv")
	if err != nil {
		log.Fatal("Unable to read input file yourfile.csv", err)
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
			log.Fatal("Read() error:", err)
		}

		vehicle := models.VehicleData{
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

		vehicles = append(vehicles, vehicle)
	}

	return vehicles
}

// Popularity computes the popularity and sorts the slice based on the sort parameter.
func Popularity(vehicles []models.VehicleData, sortOrder string) []models.VehiclePopularity {

	start := time.Now()
	popularity := make(map[models.VehiclePopularity]int)
	for _, vehicle := range vehicles {
		popularity[models.VehiclePopularity{vehicle.Make, vehicle.Model, 0}]++
	}

	var popularityVehicles []models.VehiclePopularity
	for vehicle, total := range popularity {
		popularityVehicles = append(popularityVehicles, models.VehiclePopularity{vehicle.Make, vehicle.Model, total})
	}

	// Sorting based on the sortOrder parameter
	if sortOrder == "asc" {
		sort.Slice(popularityVehicles, func(i, j int) bool {
			return popularityVehicles[i].Total < popularityVehicles[j].Total // Ascending order
		})
	} else { // Default to descending order if no sort order is specified or if it's "desc"
		sort.Slice(popularityVehicles, func(i, j int) bool {
			return popularityVehicles[i].Total > popularityVehicles[j].Total // Descending order
		})
	}

	fmt.Println("time taken", time.Since(start))
	return popularityVehicles
}
