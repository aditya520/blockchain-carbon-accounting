package main

// EmissionsCalcInput is the input for emissions calculation
type EmissionsCalcInput struct {
	UtilityID                 string `json:"utilityID"`
	PartyID                   string `json:"partyID"`
	FromDate                  string `json:"fromDate"`
	ThruDate                  string `json:"thruDate"`
	EnergUseAmount            string `json:"energyUseAmount"`
	EnergyUseUom              string `json:"energyUseUom"`
	CO2EquivalentEmissions    string `json:"CO2EquivalentEmissions"`
	NetGeneration             string `json:"netGeneration"`
	Usage                     string `json:"usage"`
	UsageuOM                  string `json:"UsageuOM"`
	NetGenerationUOM          string `json:"netGenerationUOM"`
	CO2equivalentEmissionsuOM string `json:"CO2EquivalentEmissionsuOM"`
	EmissionsUOM              string `json:"emissionsUOM"`
}

// UtilityEmissionsFactors is seed data for emissions factors used to calculate emissions based on audited utility data
type UtilityEmissionsFactors struct {
	UtilityID              string  `json:"utilityID"`
	UtilityName            string  `json:"utilitName"`
	Year                   string  `json:"year"`
	Country                string  `json:"country"`
	DivisionType           string  `json:"divisionType"`
	DivisionID             string  `json:"divisionId"`
	DivisionName           string  `json:"divisionName"`
	NetGeneration          float32 `json:"netGeneration"`
	NetGenerationUOM       string  `json:"netGenerationUOM"`
	CO2EquivalentEmissions float32 `json:"CO2EquivalentEmissions"`
	EmissionsUOM           string  `json:"emissionsUOM"`
}
