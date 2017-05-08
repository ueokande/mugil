package model

import (
	"app/shared/database"
	"testing"
	"time"
)

func insertTasks() error {
	data := [][]interface{}{
		{0, "A", "2010-04-01", 15 * time.Minute, "apple"},
		{0, "B", "2010-04-02", 30 * time.Minute, "banana"},
		{0, "C", "2010-04-02", 60 * time.Minute, "cherry"},
		{1, "S", "2010-04-01", 90 * time.Minute, "salt"},
		{1, "A", "2010-04-01", 90 * time.Minute, "apple"},
		{1, "B", "2010-04-01", 60 * time.Minute, "banana"},
		{1, "C", "2010-04-02", 90 * time.Minute, "cherry"},
		{1, "S", "2010-04-02", 90 * time.Minute, "salt"},
		{1, "A", "2010-04-02", 90 * time.Minute, "apple"},
		{1, "B", "2010-04-02", 90 * time.Minute, "banana"},
	}
	for _, d := range data {
		_, err := database.SQL.Exec(
			"INSERT INTO task (user_id, priority, date, time, description) VALUES (?, ?, ?, ?, ?)",
			d[0], d[1], d[2], d[3], d[4],
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func TestTasksByUserIDDate(t *testing.T) {
	err := insertTasks()
	if err != nil {
		t.Fatal(err)
	}

	cases := []struct {
		uid  int64
		date time.Time
		size int
	}{
		{0, time.Date(2010, 4, 1, 0, 0, 0, 0, time.UTC), 1},
		{0, time.Date(2010, 4, 2, 0, 0, 0, 0, time.UTC), 2},
		{1, time.Date(2010, 4, 1, 0, 0, 0, 0, time.UTC), 3},
		{1, time.Date(2010, 4, 2, 0, 0, 0, 0, time.UTC), 4},
	}

	for _, c := range cases {
		tasks, err := TasksByUserIdAndDate(c.uid, c.date)
		if err != nil {
			t.Fatal(err)
		}
		if len(tasks) != c.size {
			t.Error("Unexpected tsk size: ", len(tasks))
		}
	}
}
