package transport

import "github.com/oarkflow/peer-calls/server/identifiers"

type TrackJSON struct {
  TrackID identifiers.TrackID `json:"id"`
  PeerID  identifiers.PeerID  `json:"peerId"`
  Codec   Codec               `json:"codec"`
}
