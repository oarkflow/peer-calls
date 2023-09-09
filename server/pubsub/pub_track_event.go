package pubsub

import (
  "github.com/oarkflow/peer-calls/server/transport"
)

type PubTrackEvent struct {
  PubTrack PubTrack                 `json:"pubTrack"`
  Type     transport.TrackEventType `json:"type"`
}
