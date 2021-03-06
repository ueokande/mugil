package model

import (
	"app/shared/database"
	"time"
)

type Task struct {
	Id            int64
	Priority      string
	Date          time.Time
	EstimatedTime time.Duration
	Description   string
	Done          bool
	Canceled      bool
}

func TasksByUserIdAndDate(id int64, date time.Time) ([]*Task, error) {
	rows, err := database.SQL.Query(
		"SELECT id, priority, estimated_time, description, done, canceled FROM task WHERE user_id = ? and date = ?",
		id, date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make([]*Task, 0)
	for rows.Next() {
		var t Task
		err := rows.Scan(&t.Id, &t.Priority, &t.EstimatedTime, &t.Description, &t.Done, &t.Canceled)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &t)
	}
	return tasks, nil
}

func TaskCreate(uid int64, priority string, date time.Time, estimatedTime time.Duration, description string) (int64, error) {
	result, err := database.SQL.Exec(
		"INSERT INTO task (user_id, priority, date, estimated_time, description) VALUES (?, ?, ?, ?, ?)",
		uid, priority, date, estimatedTime, description)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func TaskUpdateDone(id int64, done bool) error {
	_, err := database.SQL.Exec("UPDATE task SET done = ? WHERE id = ?", done, id)
	return err
}

func TaskUpdateCanceled(id int64, canceled bool) error {
	_, err := database.SQL.Exec("UPDATE task SET canceled = ? WHERE id = ?", canceled, id)
	return err
}
