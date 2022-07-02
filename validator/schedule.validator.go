package validator

import (
	"time"

	"example.com/final-exam/model"
)

var dateLayout string = "2006-01-02"
var timeLayout string = "15:04:05"

func ValidateSchedule(schedule model.Schedule) (bool, string) {
	if schedule.RouteID <= 0 {
		return true, "Invalid parameter: routeID"
	}

	if schedule.Date == "" {
		return true, "Missing required parameter: date"
	} else {
		_, err := time.Parse(dateLayout, schedule.Date)

		if err != nil {
			return true, "Wrong parameter format: date"
		}
	}

	if schedule.TimeOfDeparture == "" {
		return true, "Missing required parameter: timeOfDeparture"
	} else {
		_, err := time.Parse(timeLayout, schedule.TimeOfDeparture)
		if err != nil {
			return true, "Wrong parameter format: timeOfDeparture"
		}
	}

	if schedule.Duration == "" {
		return true, "Missing required parameter: duration"
	} else {
		_, err := time.Parse(timeLayout, schedule.Duration)
		if err != nil {
			return true, "Wrong parameter format: duration"
		}
	}

	return false, ""
}

func ValidateDelay(delay model.Delay) (bool, string) {
	if delay.ScheduleID <= 0 {
		return true, "Invalid parameter: scheduleID"
	}

	if delay.Delay == "" {
		return true, "Missing required parameter: delay"
	} else {
		_, err := time.Parse(timeLayout, delay.Delay)
		if err != nil {
			return true, "Wrong parameter format: delay"
		}
	}

	return false, ""
}

func ValidateCancel(cancel model.Cancel) (bool, string) {
	if cancel.ScheduleID <= 0 {
		return true, "Invalid parameter: scheduleID"
	}

	return false, ""
}
