package bamboohr

import (
	"context"
	"net/http"
	"net/url"
	"path"
	"time"
)

const (
	timeOffDateFormatStr = "2006-01-02"
)

// TimeOffDetailList contains a list of TimeOffDetails
type TimeOffDetailList []TimeOffDetails

// TimeOffDetails contains the details of a time off request
type TimeOffDetails struct {
	ID         int    `json:"id"`
	Type       string `json:"type"`
	EmployeeID int    `json:"employeeId"`
	Name       string `json:"name"`
	Start      string `json:"start"`
	End        string `json:"end"`
}

// GetWhosOut gets a list of whos out, optionally between start and end dates
func (c *Client) GetWhosOut(ctx context.Context, start, end *time.Time) (TimeOffDetailList, error) {
	var list TimeOffDetailList

	u, err := url.Parse(c.BaseURL)
	if err != nil {
		return list, err
	}
	u.Path = path.Join(u.Path, "time_off/whos_out")

	query := u.Query()
	if start != nil {
		query.Add("start", start.Format(timeOffDateFormatStr))
	}
	if end != nil {
		query.Add("end", end.Format(timeOffDateFormatStr))
	}
	u.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return list, err
	}

	req = req.WithContext(ctx)
	if err := c.makeRequest(req, &list); err != nil {
		return list, err
	}

	return list, nil
}
