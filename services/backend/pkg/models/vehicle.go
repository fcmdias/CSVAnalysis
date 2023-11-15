package models

// Define a struct that matches the structure of your CSV file.
type VehicleData struct {
	VIN             string
	County          string
	City            string
	State           string
	PostalCode      string
	ModelYear       string
	Make            string
	Model           string
	EVType          string
	CAFVEligibility string
	ElectricRange   string
	BaseMSRP        string
	LegislativeDist string
	DOLVehicleID    string
	VehicleLocation string
	ElectricUtility string
	CensusTract     string
}

type VehiclePopularity struct {
	Make  string `json:"make"`
	Model string `json:"model"`
	Total int    `json:"total"`
}

type VehicleByYear struct {
	Total int `json:"total"`
	Year  int `json:"year"`
}
