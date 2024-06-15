package clickup

import (
	"encoding/json"
	"io"
	"net/http"
)

type ClickUpAPI struct {
	Token string
	url   string
}

func New(token string) ClickUpAPI {
	return ClickUpAPI{
		Token: token,
		url:   "https://api.clickup.com/api/v2",
	}
}

func (c ClickUpAPI) sendRequest(path string) *http.Response {
	req, err := http.NewRequest("GET", c.url+path, nil)

	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", c.Token)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}
	return res
}

func (c *ClickUpAPI) getRequestBody(res *http.Response, placeholder any) {
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(body, placeholder); err != nil {
		panic("ERROR: can not decode json response")
	}
}
