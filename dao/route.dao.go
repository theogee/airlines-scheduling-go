package dao

import (
	"example.com/final-exam/model"
	"example.com/final-exam/util"
)

func GetRoutes() ([]model.Route, error) {
	var routes []model.Route

	rows, err := util.DB.Query("SELECT * FROM route")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var route model.Route
		if err := rows.Scan(&route.ID, &route.AirlineID, &route.FlightCode, &route.Origin, &route.Destination); err != nil {
			return nil, err
		}
		routes = append(routes, route)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return routes, nil
}

func AddRoute(route model.Route) (int64, error) {
	result, err := util.DB.Exec("INSERT INTO route (airline_id, flight_code, origin, destination) VALUE (?, ?, ?, ?)", route.AirlineID, route.FlightCode, route.Origin, route.Destination)
	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return id, nil
}
