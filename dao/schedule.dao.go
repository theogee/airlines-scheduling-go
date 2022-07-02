package dao

import (
	"example.com/final-exam/model"
	"example.com/final-exam/util"
)

func GetScheduleByDate(date string) ([]model.Schedule, error) {
	var schedules []model.Schedule

	rows, err := util.DB.Query("SELECT * FROM schedule WHERE date = ?", date)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var schedule model.Schedule
		if err := rows.Scan(&schedule.ID, &schedule.RouteID, &schedule.Date, &schedule.TimeOfDeparture, &schedule.Duration, &schedule.Status); err != nil {
			return nil, err
		}
		schedules = append(schedules, schedule)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return schedules, nil
}

func AddSchedule(schedule model.Schedule) (int64, error) {
	result, err := util.DB.Exec("INSERT INTO schedule (route_id, date, time_of_departure, duration, status) VALUE (?, ?, ?, ?, 'on going')", schedule.RouteID, schedule.Date, schedule.TimeOfDeparture, schedule.Duration)
	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return id, nil
}

func DelaySchedule(delay model.Delay) (int64, error) {
	result, err := util.DB.Exec("UPDATE schedule SET time_of_departure = ADDTIME(time_of_departure, ?), status = 'delayed' WHERE id = ?", delay.Delay, delay.ScheduleID)
	if err != nil {
		return -1, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return -1, err
	}

	return rowsAffected, nil
}

func CancelSchedule(cancel model.Cancel) (int64, error) {
	result, err := util.DB.Exec("UPDATE schedule SET status = 'canceled' WHERE id = ?", cancel.ScheduleID)
	if err != nil {
		return -1, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return -1, err
	}

	return rowsAffected, nil
}
