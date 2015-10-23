package slacker

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Slacker struct {
	URL      string
	Icon     string
	Username string
}

type payload struct {
	text     string `json:"text"`
	icon     string `json:"icon_url"`
	username string `json:"username"`
}

func (s *Slacker) Send(text string) error {

	payload, err := json.Marshal(payload{
		text:     text,
		icon:     s.Icon,
		username: s.Username,
	})
	if err != nil {
		return err
	}

	_, err = http.Post(s.URL, "application/json", bytes.NewBuffer(payload))
	return err
}
