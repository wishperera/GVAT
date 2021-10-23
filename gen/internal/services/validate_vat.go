package services

import (
	"context"
	"fmt"
	"github.com/wishperera/GVAT/gen/internal/application"
	"github.com/wishperera/GVAT/gen/internal/container"
	"github.com/wishperera/GVAT/gen/internal/domain/adaptors"
	"github.com/wishperera/GVAT/gen/internal/pkg/log"
	"regexp"
)

type ValidateVAT struct {
	log      log.Logger
	adaptors struct {
		euVies adaptors.EUVIESAdaptor
	}
}

func (v *ValidateVAT) Init(c container.Container) error {
	v.log = c.Resolve(application.ModuleLogger).(log.Logger).NewLog("services.vat_validation")
	v.adaptors.euVies = c.Resolve(application.ModuleEUVIESAdaptor).(adaptors.EUVIESAdaptor)

	return nil
}

func (v *ValidateVAT) Validate(ctx context.Context, id string) (err error) {
	err = v.validateFormat(ctx, id)
	if err != nil {
		return err
	}

	return v.checkAgainstVIES(ctx, id)
}

// validateFormat: checks the id is a german id, and validates whether the format is correct
func (v *ValidateVAT) validateFormat(ctx context.Context, id string) error {
	// check the prefix for the country code
	valid, err := regexp.Match("DE[[:digit:]]{9}", []byte(id))
	if err != nil {
		return ValidationError{
			fmt.Errorf("failed to validate id due: %s", err),
		}
	}

	if !valid {
		return ValidationError{
			fmt.Errorf("provided id:[%s] is not a valid German vat number", id),
		}
	}

	return nil
}

// checkAgainstVIES: cross-check the id against the online vies database
func (v *ValidateVAT) checkAgainstVIES(ctx context.Context, id string) error {
	valid, err := v.adaptors.euVies.ValidateVATID(ctx, id)
	if err != nil {
		return DependencyError{
			fmt.Errorf("failed to check against vies database due: %s", err),
		}
	}

	if !valid {
		return ValidationError{
			fmt.Errorf("provided id:[%s] is not a valid German vat number", id),
		}
	}

	return nil
}
