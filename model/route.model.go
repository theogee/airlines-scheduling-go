package model

type Route struct {
	ID          int64  `json: "id"`
	AirlineID   int64  `json: "airlineID"`
	FlightCode  string `json: "flightCode"`
	Origin      string `json: "origin"`
	Destination string `json: "destination"`
}
