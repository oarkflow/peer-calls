package uuid

import (
  "github.com/google/uuid"

  "github.com/oarkflow/peer-calls/server/basen"
)

var defaultBaseNEncoder = basen.NewBaseNEncoder(basen.AlphabetBase62)

func New() string {
  value := uuid.New()

  return defaultBaseNEncoder.Encode(value[:])
}
