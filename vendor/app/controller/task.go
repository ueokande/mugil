package controller

import (
	"app/model"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

type HumanReadableTime time.Duration

func (t HumanReadableTime) String() string {
	d := time.Duration(t)
	hour := d / time.Hour
	min := (d - hour*time.Hour) / time.Minute
	if hour == 0 {
		return fmt.Sprintf("%dmin", min)
	} else if min == 0 {
		return fmt.Sprintf("%dh", hour)
	} else {
		return fmt.Sprintf("%dh %dmin", hour, min)
	}
}

type TaskDto struct {
	Time        HumanReadableTime
	Priority    string
	Description string
}

func TaskIndex(c echo.Context) error {
	tasks, err := model.TasksByUserID(0)
	if err != nil {
		return err
	}

	dtos := make([]TaskDto, 0, len(tasks))
	for _, t := range tasks {
		dtos = append(dtos, TaskDto{
			Time:        HumanReadableTime(t.Time),
			Priority:    t.Priority,
			Description: t.Description,
		})
	}
	return c.Render(http.StatusOK, "task_index.html", dtos)
}
