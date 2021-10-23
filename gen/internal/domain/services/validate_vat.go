package services

import "context"

type ValidateVATID interface {
	Validate(ctx context.Context, id string) (err error)
}
