package yandextrackerapi

import (
	"net/http"
)

type Connection struct {
	Client  *http.Client
	Headers map[string]string
	Host    string
	Timeout uint8
	Retries uint8
	Verify  bool
}

func NewConnection(token, orgID string) *Connection {
	return &Connection{
		Client: &http.Client{},
		Headers: map[string]string{
			"Content-Type":  "application/json",
			"Authorization": "OAuth " + token,
			"X-Org-Id":      orgID,
		},
		Host:    "https://api.tracker.yandex.net/v2",
		Timeout: 10,
		Retries: 10,
		Verify:  true,
	}
}
