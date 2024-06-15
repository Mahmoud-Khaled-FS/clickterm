package clickup

import (
	"errors"
	"net/http"
)

type User struct {
	ID                int         `json:"id"`
	Username          string      `json:"username"`
	Email             string      `json:"email"`
	Color             string      `json:"color"`
	ProfilePicture    interface{} `json:"profilePicture"`
	Initials          string      `json:"initials"`
	WeekStartDay      int         `json:"week_start_day"`
	GlobalFontSupport bool        `json:"global_font_support"`
	Timezone          string      `json:"timezone"`
}

func (c ClickUpAPI) GetUser() (User, error) {
	res := c.sendRequest("/user")
	if res.StatusCode != http.StatusOK {
		return User{}, errors.New("Unauthorized")
	}
	body := struct {
		User User `json:"user"`
	}{}
	c.getRequestBody(res, &body)
	return body.User, nil
}
