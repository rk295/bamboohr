package bamboohr

import (
	"context"
	"net/http"
	"net/url"
	"path"
	"time"
)

// Users is a list of Users
type Users []User

// User holds the details of a BambooHR user
type User struct {
	ID         int       `json:"id"`
	EmployeeID int       `json:"employeeId"`
	FirstName  string    `json:"firstName"`
	LastName   string    `json:"lastName"`
	Email      string    `json:"email"`
	Status     string    `json:"status"`
	LastLogin  time.Time `json:"lastLogin"`
}

// GetUserList returns a list of users
func (c *Client) GetUserList(ctx context.Context) (Users, error) {
	var userList Users

	u, err := url.Parse(c.BaseURL)
	if err != nil {
		return userList, err
	}
	u.Path = path.Join(u.Path, "meta/users/")

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return userList, err
	}

	mp := make(map[string]User)

	req = req.WithContext(ctx)
	if err := c.makeRequest(req, &mp); err != nil {
		return userList, err
	}

	for _, u := range mp {
		userList = append(userList, u)
	}

	return userList, nil
}
