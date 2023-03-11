package monopoly

import (
	"bytes"
	"encoding/json"
	"github.com/imthaghost/aragog/config"
	"net/http"
)

// ClientWrapper ...
type ClientWrapper interface {
	HealthCheck() (string, error)
	SendMessage(msg string) (string, error)
}

type clientWrapper struct {
	Client *http.Client
	Config *config.Config
}

// NewClient ...
func NewClient(cfg *config.Config) ClientWrapper {

	return &clientWrapper{
		Config: cfg,
		Client: http.DefaultClient,
	}
}

// HealthCheck ...
func (c *clientWrapper) HealthCheck() (string, error) {
	return "", nil
}

// SendMessage ...
func (c *clientWrapper) SendMessage(msg string) (string, error) {
	values := map[string]string{"msg": msg}
	jsonData, err := json.Marshal(values)

	if err != nil {

		return "", err
	}

	resp, err := c.Client.Post(c.Config.Monopoly.Host+"api/v1/aragog/webhooks", "application/json",
		bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	return resp.Status, nil
}
