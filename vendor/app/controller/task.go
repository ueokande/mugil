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

type TaskCreateForm struct {
	Time        time.Duration `form:"time",json:"time"`
	Priority    string        `form:"priority",json:"priority"`
	Description string        `form:"description",json:"description"`
}

type IdResponse struct {
	Id int64 `json:"id"`
}

type TaskDto struct {
	Time        HumanReadableTime
	Priority    string
	Description string
	Done        bool
	Canceled    bool
}

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

func TaskIndex(c echo.Context) error {
	uid, err := CurrentUserId(c)
	if err != nil {
		return c.Redirect(http.StatusFound, "/login")
	}

	tasks, err := model.TasksByUserID(uid)
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

func TaskCreate(c echo.Context) error {
	uid, err := CurrentUserId(c)
	if err != nil {
		return c.Redirect(http.StatusFound, "/login")
	}

	var form TaskCreateForm
	err = c.Bind(&form)
	log.Info(form)
	if err != nil {
		log.Error(err)
		return err
	}

	id, err := model.TaskCreate(uid, form.Priority, time.Now(), form.Time, form.Description)
	if err != nil {
		log.Error(err)
		return err
	}
	return c.JSON(http.StatusOK, IdResponse{Id: id})
}
