package crawler

import (
	"github.com/imthaghost/aragog/config"
	"github.com/imthaghost/aragog/internal/tradingview"
)

// Crawler ...
type Crawler struct {
	Config config.Config
}

// New create a new TradingView crawler using Browserless
func New(cfg config.Config) tradingview.Service {

	return &Crawler{
		Config: cfg,
	}
}
