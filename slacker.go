package slacker

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type Slacker struct {
	URL      string
	Icon     string
	Username string
}

type payload struct {
	Text     string `json:"text"`
	Icon     string `json:"icon_url"`
	Username string `json:"username"`
}

func (s *Slacker) Send(text string) error {

	if len(s.URL) == 0 {
		return errors.New("URL is required")
	}

	payload, err := json.Marshal(payload{
		Text:     text,
		Icon:     s.Icon,
		Username: s.Username,
	})
	if err != nil {
		return err
	}

	_, err = http.Post(s.URL, "application/json", bytes.NewBuffer(payload))
	return err
}
