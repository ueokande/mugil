package controller

import (
	"app/model"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
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
	Done        bool
	Canceled    bool
}

func TaskIndex(c echo.Context) error {
	tasks, err := model.TasksByUserID(0)
	if err != nil {
		log.Error(err)
		return err
	}

	dtos := make([]TaskDto, 0, len(tasks))
	for _, t := range tasks {
		dtos = append(dtos, TaskDto{
			Time:        HumanReadableTime(t.Time),
			Priority:    t.Priority,
			Description: t.Description,
			Done:        t.Done,
			Canceled:    t.Canceled,
		})
	}
	err = c.Render(http.StatusOK, "task_index.html", dtos)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

type TaskCreateForm struct {
	Time        time.Duration `form:"time",json:"time"`
	Priority    string        `form:"priority",json:"priority"`
	Description string        `form:"description",json:"description"`
}

func TaskCreate(c echo.Context) error {
	var form TaskCreateForm
	err := c.Bind(&form)
	log.Info(form)
	if err != nil {
		log.Error(err)
		return err
	}

	task, err := model.TaskCreate(form.Priority, time.Now(), form.Time, form.Description, false, false)
	if err != nil {
		log.Error(err)
		return err
	}
	return c.JSON(http.StatusOK, task)
}
