package controller

import (
	"app/model"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type TaskIndexJsonForm struct {
	Date string `query:"date"`
}

type TaskCreateForm struct {
	EstimatedTime time.Duration `json:"estimated_time"`
	Priority      string        `json:"priority"`
	Description   string        `json:"description"`
}

type IdResponse struct {
	Id int64 `json:"id"`
}

type TaskDto struct {
	Id            int64         `json:"id"`
	EstimatedTime time.Duration `json:"estimated_time"`
	Priority      string        `json:"priority"`
	Description   string        `json:"description"`
	Done          bool          `json:"done"`
	Canceled      bool          `json:"canceled"`
}

type MessageJsonDto struct {
	Message string `json:"message"`
}

func TaskIndex(c echo.Context) error {
	_, err := CurrentUserId(c)
	if err != nil {
		return c.Redirect(http.StatusFound, "/login")
	}
	err = c.Render(http.StatusOK, "task_index.html", nil)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func TaskIndexJson(c echo.Context) error {
	uid, err := CurrentUserId(c)
	if err != nil {
		return c.JSON(http.StatusForbidden, MessageJsonDto{"not authenticated"})
	}

	var form TaskIndexJsonForm
	err = c.Bind(&form)
	date, err := time.Parse("2006-01-02", form.Date)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusForbidden, MessageJsonDto{"invalid date"})
	}

	tasks, err := model.TasksByUserIdAndDate(uid, date)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusForbidden, MessageJsonDto{"invalid request"})
	}

	dtos := make([]TaskDto, 0, len(tasks))
	for _, t := range tasks {
		dtos = append(dtos, TaskDto{
			Id:            t.Id,
			EstimatedTime: t.EstimatedTime,
			Priority:      t.Priority,
			Description:   t.Description,
			Done:          t.Done,
			Canceled:      t.Canceled,
		})
	}
	return c.JSON(http.StatusOK, dtos)
}

func TaskCreateJson(c echo.Context) error {
	uid, err := CurrentUserId(c)
	if err != nil {
		return c.Redirect(http.StatusFound, "/login")
	}

	var form TaskCreateForm
	err = c.Bind(&form)
	if err != nil {
		log.Error(err)
		return err
	}

	id, err := model.TaskCreate(uid, form.Priority, time.Now(), form.EstimatedTime, form.Description)
	if err != nil {
		log.Error(err)
		return err
	}
	return c.JSON(http.StatusOK, IdResponse{Id: id})
}
