package domain

type contextKey string

func (c contextKey) String() string {
	return string(c)
}

const (
	ContextKeyTraceID contextKey = "trace-id"

	CountryCodeGermany = "DE"
)
