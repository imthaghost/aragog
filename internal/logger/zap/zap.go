package zap

import (
	"github.com/imthaghost/aragog/internal/logger"

	"go.uber.org/zap"
)

type Service struct {
	Log *zap.Logger
}

// Msg ...
func (s *Service) Msg(message string) {
	s.Log.Info(message)
}

// Error ...
func (s *Service) Error(message string) {
	s.Log.Error(message)
}

// NewService creates a
func NewService() logger.Service {
	log, _ := zap.NewProduction()
	defer log.Sync()

	return &Service{
		Log: log,
	}
}
