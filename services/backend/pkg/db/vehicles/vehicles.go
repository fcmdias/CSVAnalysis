package vehicles

import (
	"encoding/csv"
	"fmt"
	"github/com/fcmdias/CSVAnalysis/services/backend/pkg/models"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"time"
)

func FetchVehicles(filter string) (vehicles []models.VehicleData) {

	start := time.Now()
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

		switch filter {
		case "electric":
			if vehicle.EVType != "Battery Electric Vehicle (BEV)" && vehicle.EVType != "Electric Vehicle Type" {
				continue
			}
		case "hybrid":
			if vehicle.EVType != "Plug-in Hybrid Electric Vehicle (PHEV)" {
				continue
			}
		default:
			// accept all
		}

		vehicles = append(vehicles, vehicle)
	}

	log.Println("time to read csv file", time.Since(start))
	return vehicles
}

// Popularity computes the popularity and sorts the slice based on the sort parameter.
func Popularity(vehicles []models.VehicleData, sortOrder string) []models.VehiclePopularity {

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

	popularityVehicles = top20(popularityVehicles)

	// Sorting based on the sortOrder parameter
	if sortOrder == "asc" {
		sort.Slice(popularityVehicles, func(i, j int) bool {
			return popularityVehicles[i].Total < popularityVehicles[j].Total // Ascending order
		})
	} else { // Default to descending order if no sort order is specified or if it's "desc"
		// already sorted
	}

	fmt.Println("time taken", time.Since(start))
	return popularityVehicles
}

func top20(data []models.VehiclePopularity) []models.VehiclePopularity {

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

	mapByYear := make(map[models.VehicleByYear]int)
	for _, vehicle := range data {

		year, err := strconv.Atoi(vehicle.ModelYear)
		if err != nil {
			continue
		}
		if year < thisYear-10 || year > thisYear {
			continue
		}
		mapByYear[models.VehicleByYear{
			Year:  year,
			Total: 0,
		}]++
	}

	var dataByYear []models.VehicleByYear

	for vehicle, total := range mapByYear {
		dataByYear = append(dataByYear, models.VehicleByYear{
			Total: total,
			Year:  vehicle.Year,
		})
	}

	sort.Slice(dataByYear, func(i, j int) bool {
		return dataByYear[i].Year < dataByYear[j].Year
	})

	return dataByYear
}
