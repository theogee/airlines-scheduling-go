package validator

import "example.com/final-exam/model"

func ValidateAirline(airline model.Airline) (bool, string) {
	if airline.Name == "" {
		return true, "missing required parameters: name"
	}
	return false, ""
}
