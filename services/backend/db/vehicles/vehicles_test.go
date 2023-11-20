package vehicles

import (
	"github/com/fcmdias/CSVAnalysis/services/backend/models"
	"strconv"
	"testing"
	"time"
)

func TestByYear(t *testing.T) {
	currentYear := time.Now().Year()

	empty := []models.VehicleByYear{
		{Year: currentYear - 9, Total: 0},
		{Year: currentYear - 8, Total: 0},
		{Year: currentYear - 7, Total: 0},
		{Year: currentYear - 6, Total: 0},
		{Year: currentYear - 5, Total: 0},
		{Year: currentYear - 4, Total: 0},
		{Year: currentYear - 3, Total: 0},
		{Year: currentYear - 2, Total: 0},
		{Year: currentYear - 1, Total: 0},
		{Year: currentYear, Total: 0},
	}

	tests := []struct {
		name     string
		data     []models.VehicleData
		expected []models.VehicleByYear
	}{
		{
			name:     "EmptyData",
			data:     []models.VehicleData{},
			expected: empty,
		},
		{
			name: "ValidData",
			data: []models.VehicleData{
				{ModelYear: strconv.Itoa(currentYear - 2)},
				{ModelYear: strconv.Itoa(currentYear - 2)},
				{ModelYear: strconv.Itoa(currentYear - 5)},
			},
			expected: []models.VehicleByYear{
				{Year: currentYear - 9, Total: 0},
				{Year: currentYear - 8, Total: 0},
				{Year: currentYear - 7, Total: 0},
				{Year: currentYear - 6, Total: 0},
				{Year: currentYear - 5, Total: 1},
				{Year: currentYear - 4, Total: 0},
				{Year: currentYear - 3, Total: 0},
				{Year: currentYear - 2, Total: 2},
				{Year: currentYear - 1, Total: 0},
				{Year: currentYear, Total: 0},
			},
		},
		{
			name: "DataWithInvalidYear",
			data: []models.VehicleData{
				{ModelYear: "InvalidYear"},
				{ModelYear: strconv.Itoa(currentYear - 1)},
			},
			expected: []models.VehicleByYear{
				{Year: currentYear - 9, Total: 0},
				{Year: currentYear - 8, Total: 0},
				{Year: currentYear - 7, Total: 0},
				{Year: currentYear - 6, Total: 0},
				{Year: currentYear - 5, Total: 0},
				{Year: currentYear - 4, Total: 0},
				{Year: currentYear - 3, Total: 0},
				{Year: currentYear - 2, Total: 0},
				{Year: currentYear - 1, Total: 1},
				{Year: currentYear, Total: 0},
			},
		},
		{
			name: "DataOutOfRange",
			data: []models.VehicleData{
				{ModelYear: strconv.Itoa(currentYear + 1)},  // Future year
				{ModelYear: strconv.Itoa(currentYear - 11)}, // More than 10 years ago
			},
			expected: empty,
		},
		{
			name: "DataWithBoundaryYears",
			data: []models.VehicleData{
				{ModelYear: strconv.Itoa(currentYear - 10)},
				{ModelYear: strconv.Itoa(currentYear - 11)},
				{ModelYear: strconv.Itoa(currentYear)},
			},
			expected: []models.VehicleByYear{
				{Year: currentYear - 9, Total: 0},
				{Year: currentYear - 8, Total: 0},
				{Year: currentYear - 7, Total: 0},
				{Year: currentYear - 6, Total: 0},
				{Year: currentYear - 5, Total: 0},
				{Year: currentYear - 4, Total: 0},
				{Year: currentYear - 3, Total: 0},
				{Year: currentYear - 2, Total: 0},
				{Year: currentYear - 1, Total: 0},
				{Year: currentYear, Total: 1},
			},
		},
		{
			name: "DataWithSameYear",
			data: []models.VehicleData{
				{ModelYear: strconv.Itoa(currentYear - 3)},
				{ModelYear: strconv.Itoa(currentYear - 3)},
				{ModelYear: strconv.Itoa(currentYear - 3)},
			},
			expected: []models.VehicleByYear{
				{Year: currentYear - 9, Total: 0},
				{Year: currentYear - 8, Total: 0},
				{Year: currentYear - 7, Total: 0},
				{Year: currentYear - 6, Total: 0},
				{Year: currentYear - 5, Total: 0},
				{Year: currentYear - 4, Total: 0},
				{Year: currentYear - 3, Total: 3},
				{Year: currentYear - 2, Total: 0},
				{Year: currentYear - 1, Total: 0},
				{Year: currentYear, Total: 0},
			},
		},
		{
			name: "DataWithSequentialYears",
			data: []models.VehicleData{
				{ModelYear: strconv.Itoa(currentYear - 1)},
				{ModelYear: strconv.Itoa(currentYear - 2)},
				{ModelYear: strconv.Itoa(currentYear - 3)},
			},
			expected: []models.VehicleByYear{
				{Year: currentYear - 9, Total: 0},
				{Year: currentYear - 8, Total: 0},
				{Year: currentYear - 7, Total: 0},
				{Year: currentYear - 6, Total: 0},
				{Year: currentYear - 5, Total: 0},
				{Year: currentYear - 4, Total: 0},
				{Year: currentYear - 3, Total: 1},
				{Year: currentYear - 2, Total: 1},
				{Year: currentYear - 1, Total: 1},
				{Year: currentYear, Total: 0},
			},
		},
		{
			name: "DataWithMixedValidInvalidYears",
			data: []models.VehicleData{
				{ModelYear: "InvalidYear1"},
				{ModelYear: strconv.Itoa(currentYear - 4)},
				{ModelYear: "InvalidYear2"},
				{ModelYear: strconv.Itoa(currentYear - 2)},
			},
			expected: []models.VehicleByYear{
				{Year: currentYear - 9, Total: 0},
				{Year: currentYear - 8, Total: 0},
				{Year: currentYear - 7, Total: 0},
				{Year: currentYear - 6, Total: 0},
				{Year: currentYear - 5, Total: 0},
				{Year: currentYear - 4, Total: 1},
				{Year: currentYear - 3, Total: 0},
				{Year: currentYear - 2, Total: 1},
				{Year: currentYear - 1, Total: 0},
				{Year: currentYear, Total: 0},
			},
		},
		{
			name: "DataWithExtremeYears",
			data: []models.VehicleData{
				{ModelYear: "1900"},
				{ModelYear: "3000"},
			},
			expected: empty,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ByYear(tt.data)

			if !compareVehicleByYearSlices(result, tt.expected) {
				t.Errorf("ByYear(%v) = %v (%T), expected %v (%T)", tt.data, result, result, tt.expected, tt.expected)
			}

		})
	}
}

func compareVehicleByYearSlices(a, b []models.VehicleByYear) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
