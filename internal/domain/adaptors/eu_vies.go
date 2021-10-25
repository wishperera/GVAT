package adaptors

//go:generate mockgen -destination=../../mocks/mock_euvies_adaptor.go -package=mocks -source=./eu_vies.go
import (
	"context"
)

// EUVIESAdaptor : manages the interaction with the EU/VIES online database
type EUVIESAdaptor interface {
	ValidateVATID(ctx context.Context, countryCode string, vatID string) (valid bool, err error)
}
