package dao

import (
	"example.com/final-exam/model"
	"example.com/final-exam/util"
)

func GetAirlines() ([]model.Airline, error) {
	var airlines []model.Airline

	rows, err := util.DB.Query("SELECT * FROM airlines")
	if err != nil {
		return nil, err
	}
	defer util.DB.Close()

	for rows.Next() {
		var airline model.Airline
		if err := rows.Scan(&airline.ID, &airline.Name); err != nil {
			return nil, err
		}
		airlines = append(airlines, airline)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return airlines, nil
}

func AddAirline(airline model.Airline) (int64, error) {
	result, err := util.DB.Exec("INSERT INTO airlines (name) VALUE (?)", airline.Name)
	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return id, nil
}
