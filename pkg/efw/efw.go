package efw

import "context"

type EFW struct {
	ctx context.Context
}

// New creates a new EFW instance with the provided context.
func New(ctx context.Context) *EFW {
	return &EFW{
		ctx: ctx,
	}
}
