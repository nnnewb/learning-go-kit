package middleware

import (
	"play/stringsvc"

	"github.com/go-kit/log"
)

type LoggingMiddleware struct {
	stringsvc.StringService
	Logger log.Logger
}

func (m *LoggingMiddleware) Uppercase(s string) (string, error) {
	m.Logger.Log("msg", "call endpoint", "arg", s)
	defer m.Logger.Log("msg", "called endpoint", "arg", s)
	return m.StringService.Uppercase(s)
}
