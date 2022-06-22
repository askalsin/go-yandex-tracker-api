package yandextrackerapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Issue struct {
	Summary string `json:"summary"`
	Queue   struct {
		ID  int64  `json:"id"`
		Key string `json:"key"`
	} `json:"queue"`
}

func NewIssue() *Issue {
	return &Issue{}
}

func (c *Connection) CreateNewIssue(issue *Issue) ([]byte, error) {
	url := fmt.Sprintf("%s%s", c.Host, "/issues")

	body, err := json.Marshal(issue)
	if err != nil {
		return nil, err
	}
	responseBody := bytes.NewBuffer(body)

	req, err := http.NewRequest(http.MethodPost, url, responseBody)
	if err != nil {
		return nil, err
	}

	for key, val := range c.Headers {
		req.Header.Add(key, val)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return rbody, nil
}
