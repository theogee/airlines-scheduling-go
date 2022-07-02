package model

type Schedule struct {
	ID              int64  `json: "id"`
	RouteID         int64  `json: "routeID"`
	Date            string `json: "date"`
	TimeOfDeparture string `json: "timeOfDeparture"`
	Duration        string `json: "duration"`
	Status          string `json: "status"`
}

type Delay struct {
	ScheduleID int64  `json: "scheduleID"`
	Delay      string `json: "delay"`
}

type Cancel struct {
	ScheduleID int64 `json: "scheduleID"`
}
