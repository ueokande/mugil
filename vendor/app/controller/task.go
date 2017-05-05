package controller

import (
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

type Task struct {
	Name string
	Time HumanReadableTime
}

func TaskIndex(c echo.Context) error {
	tasks := []Task{
		{
			Name: "wash my head",
			Time: HumanReadableTime(15 * time.Minute),
		},
		{
			Name: "wash my hands",
			Time: HumanReadableTime(60 * time.Minute),
		},
	}

	return c.Render(http.StatusOK, "task_index.html", tasks)
}
