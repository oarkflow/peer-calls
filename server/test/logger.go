package test

import (
  "github.com/oarkflow/peer-calls/server/logformatter"
  "github.com/oarkflow/peer-calls/server/logger"
)

func NewLogger() logger.Logger {
  return logger.NewFromEnv("PEERCALLS_LOG").WithFormatter(logformatter.New())
}
