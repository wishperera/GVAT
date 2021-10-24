package adaptors

import (
	"context"
)

// EUVIESAdaptor : manages the interaction with the EU/VIES online database
type EUVIESAdaptor interface {
	ValidateVATID(ctx context.Context, countryCode string, vatID string) (valid bool, err error)
}
