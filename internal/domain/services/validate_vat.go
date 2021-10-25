package services

//go:generate mockgen -destination=../../mocks/mock_vat_validate_service.go -package=mocks -source=validate_vat.go
import (
	"context"
)

type ValidateVATID interface {
	Validate(ctx context.Context, id string) (valid bool, err error)
}
