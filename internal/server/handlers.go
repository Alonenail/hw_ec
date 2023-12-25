package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func parseQueryDate(until string) (*time.Time, error) {
	formattedTime, err := time.Parse(time.DateOnly, until)
	if err != nil {
		return nil, err
	}
	return &formattedTime, nil
}

func GetDaysLeft(c echo.Context) error {
	toDate := c.QueryParam("until")
	untilDate := time.Date(2030, time.January, 1, 0, 0, 0, 0, time.UTC)

	if toDate != "" {
		parsed, err := parseQueryDate(toDate)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad query. Accept like '?until=2030-01-01'")
		}
		untilDate = *parsed
	}

	days := untilDate.Sub(time.Now()).Hours() / 24

	return c.String(http.StatusOK, fmt.Sprintf("%.f", days))
}
