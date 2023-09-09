package cli

import (
  "context"

  "github.com/juju/errors"

  "github.com/oarkflow/peer-calls/server"
  "github.com/oarkflow/peer-calls/server/logger"
)

type Props struct {
  Log     logger.Logger
  Version string
  Args    []string
  Embed   server.Embed
}

func Exec(ctx context.Context, props Props) error {
  cmd := NewRootCommand(props)
  err := cmd.Exec(ctx, props.Args)

  return errors.Trace(err)
}
