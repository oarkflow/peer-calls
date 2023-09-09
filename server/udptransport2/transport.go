package udptransport2

import (
  "sync"

  "github.com/pion/interceptor"

  "github.com/oarkflow/peer-calls/server/identifiers"
  "github.com/oarkflow/peer-calls/server/logger"
  "github.com/oarkflow/peer-calls/server/servertransport"
  "github.com/oarkflow/peer-calls/server/stringmux"
)

type Transport struct {
  *servertransport.Transport
  closeWriteOnce sync.Once
  streamID       identifiers.RoomID
  closeWrite     func()
}

func NewTransport(
  log logger.Logger,
  streamID identifiers.RoomID,
  mediaConn stringmux.Conn,
  dataConn stringmux.Conn,
  metadataConn stringmux.Conn,
  interceptorRegistry *interceptor.Registry,
) *Transport {
  closeWrite := func() {
    mediaConn.CloseWrite()
    dataConn.CloseWrite()
    metadataConn.CloseWrite()
  }

  serverTransportParams := servertransport.Params{
    Log:                 log,
    MediaConn:           mediaConn,
    DataConn:            dataConn,
    MetadataConn:        metadataConn,
    InterceptorRegistry: interceptorRegistry,
    CodecRegistry:       nil,
  }

  return &Transport{
    streamID:       streamID,
    Transport:      servertransport.New(serverTransportParams),
    closeWriteOnce: sync.Once{},
    closeWrite:     closeWrite,
  }
}

func (t *Transport) CloseWrite() {
  t.closeWriteOnce.Do(t.closeWrite)
}

// TODO rename to RoomID.
func (t *Transport) StreamID() identifiers.RoomID {
  return t.streamID
}
