package inbound

import (
	"context"
)

type CliInter interface {
	Run(ctx context.Context) error
}
