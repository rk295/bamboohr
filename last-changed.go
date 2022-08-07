package bamboohr

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"time"
)

type ChangeType int

const (
	ChangeUpdated  = "updated"
	ChangeInserted = "inserted"
	ChangeDeleted  = "deleted"

	Pending ChangeType = iota
	Approved
	Rejected
	end
)

type ChangeList struct {
	Latest    time.Time                `json:"latest"`
	Employees map[string]ChangeDetails `json:"employees"`
}

type ChangeDetails struct {
	ID          string    `json:"id"`
	Action      string    `json:"action"`
	LastChanged time.Time `json:"lastChanged"`
}

// GetChanges gets all changes since the given time, supports filtering by inserted, updated or deleted changes
func (c *Client) GetChanges(ctx context.Context, since time.Time, t string) (ChangeList, error) {

	var ch ChangeList

	u, err := url.Parse(c.BaseURL)
	if err != nil {
		return ch, err
	}
	u.Path = path.Join(u.Path, "employees/changed")

	query := u.Query()
	query.Add("since", since.Format(time.RFC3339))

	switch t {
	case ChangeUpdated:
		query.Add("type", ChangeUpdated)
	case ChangeInserted:
		query.Add("type", ChangeInserted)
	case ChangeDeleted:
		query.Add("type", ChangeDeleted)
	default:
		return ch, fmt.Errorf("invalid type: %s. Possible values are %s, %s, %s", t, ChangeUpdated, ChangeInserted, ChangeDeleted)
	}
	u.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return ch, err
	}

	req = req.WithContext(ctx)
	if err := c.makeRequest(req, &ch); err != nil {
		return ch, err
	}

	return ch, nil
}
