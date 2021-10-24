package services

import (
	"context"
)

type ValidateVATID interface {
	Validate(ctx context.Context, id string) (valid bool, err error)
}
