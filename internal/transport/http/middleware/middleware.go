package middleware

//go:generate mockgen -destination=../../../mocks/mock_middleware.go -package=mocks -source=./middleware.go
import "net/http"

type Middleware interface {
	Handle(next http.Handler) http.Handler
}
