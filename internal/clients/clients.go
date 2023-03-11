/*
Package clients serves as the data access layer.
The Clients struct hosts a set of service objects which can be used to query
various Bearish services. In order for your endpoint to make outbound calls,
it must utilize the clients exposed here.
*/
package clients

import (
	"github.com/imthaghost/aragog/config"
	"github.com/imthaghost/aragog/internal/clients/monopoly"
)

// Clients holds the set of objects used to query backend services from endpoints
type Clients struct {
	Monopoly monopoly.ClientWrapper
}

// SetupClients ...
func SetupClients(cfg *config.Config) (*Clients, error) {
	// monopoly client
	monopolyClient := monopoly.NewClient(cfg)

	return &Clients{
		Monopoly: monopolyClient,
	}, nil
}
