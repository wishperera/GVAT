package services

import (
	"context"
	"fmt"
	"github.com/wishperera/GVAT/internal/application"
	"github.com/wishperera/GVAT/internal/container"
	"github.com/wishperera/GVAT/internal/domain/adaptors"
	"github.com/wishperera/GVAT/internal/pkg/log"
	"regexp"
)

type ValidateVAT struct {
	log      log.Logger
	adaptors struct {
		euVies adaptors.EUVIESAdaptor
	}
	allowed map[string]int
}

func (v *ValidateVAT) Init(c container.Container) error {
	v.log = c.Resolve(application.ModuleLogger).(log.Logger).NewLog("services.vat_validation")
	v.adaptors.euVies = c.Resolve(application.ModuleEUVIESAdaptor).(adaptors.EUVIESAdaptor)
	v.allowed = map[string]int{
		"DE": 9,
		"FR": 11,
	}

	return nil
}

func (v *ValidateVAT) Validate(ctx context.Context, id string) (valid bool, err error) {
	valid, err = v.validateFormat(ctx, id)
	if err != nil {
		return false, err
	}

	if !valid {
		return valid, nil
	}

	countryCode := id[:2]
	vatNumber := id[2:]

	return v.checkAgainstVIES(ctx, countryCode, vatNumber)
}

// validateFormat: checks the id is a german id, and validates whether the format is correct
func (v *ValidateVAT) validateFormat(_ context.Context, id string) (valid bool, err error) {
	if len(id) < 2 {
		return false, nil
	}

	length, ok := v.allowed[id[:2]]
	if !ok {
		return false, nil
	}

	match := fmt.Sprintf("[[:digit:]]{%d}", length)

	// check the prefix for the country code
	valid, err = regexp.Match(match, []byte(id[2:]))
	if err != nil {
		return false, ValidationError{
			fmt.Errorf("failed to validate id due: %s", err),
		}
	}

	return valid, nil
}

// checkAgainstVIES: cross-check the id against the online vies database
func (v *ValidateVAT) checkAgainstVIES(ctx context.Context, countryCode, id string) (valid bool, err error) {
	valid, err = v.adaptors.euVies.ValidateVATID(ctx, countryCode, id)
	if err != nil {
		return false, DependencyError{
			fmt.Errorf("failed to check against vies database due: %s", err),
		}
	}

	return valid, nil
}
