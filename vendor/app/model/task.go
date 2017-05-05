package model

import (
	"app/shared/database"
	"time"
)

type Task struct {
	Id          int64
	Priority    string
	Date        time.Time
	Time        time.Duration
	Description string
	Done        bool
	Canceled    bool
}

func TasksByUserID(id int64) ([]*Task, error) {
	rows, err := database.SQL.Query(
		"SELECT id, priority, time, description, done, canceled FROM task WHERE user_id = ?",
		id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make([]*Task, 0)
	for rows.Next() {
		var t Task
		err := rows.Scan(&t.Id, &t.Priority, &t.Time, &t.Description, &t.Done, &t.Canceled)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &t)
	}
	return tasks, nil
}

func TaskCreate(priority string, date time.Time, time time.Duration, description string, done bool, canceled bool) (*Task, error) {
	result, err := database.SQL.Exec(
		"INSERT INTO task (priority, date, time, description, done, canceled) VALUES (?, ?, ?, ?, ?, ?)",
		priority, date, time, description, done, canceled)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	task := Task{
		Id:          id,
		Priority:    priority,
		Date:        date,
		Time:        time,
		Description: description,
		Done:        done,
		Canceled:    canceled,
	}
	return &task, nil
}

func TaskUpdateDone(id int64, done bool) error {
	_, err := database.SQL.Exec("UPDATE task SET done = ? WHERE id = ?", done, id)
	return err
}

func TaskUpdateCanceled(id int64, canceled bool) error {
	_, err := database.SQL.Exec("UPDATE task SET canceled = ? WHERE id = ?", canceled, id)
	return err
}
