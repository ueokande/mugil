package model

import (
	"app/shared/database"
	"time"
)

type Task struct {
	Id          int32
	Priority    string
	Date        time.Time
	Time        time.Duration
	Description string
	Canceled    bool
}

func TasksByUserID(id int32) ([]*Task, error) {
	const query = "SELECT id, priority, time, description, canceled FROM task WHERE user_id = ?"

	rows, err := database.SQL.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make([]*Task, 0)
	for rows.Next() {
		var t Task
		err := rows.Scan(&t.Id, &t.Priority, &t.Time, &t.Description, &t.Canceled)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &t)
	}
	return tasks, nil
}
