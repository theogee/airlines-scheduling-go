package validator

import (
	"fmt"

	"example.com/final-exam/model"
)

func ValidateRoute(route model.Route) (bool, string) {
	if route.AirlineID <= 0 {
		fmt.Println(route)
		return true, "Invalid parameter: airlineID"
	}

	if route.FlightCode == "" {
		return true, "Missing required parameter: flightCode"
	}

	if route.Origin == "" {
		return true, "Missing required parameter: origin"
	}

	if route.Destination == "" {
		return true, "Missing required parameter: destination"
	}

	return false, ""
}
