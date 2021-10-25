package exception

//go:generate mockgen -destination=../../../../mocks/mock_exception_handler.go -package=mocks -source=./exception.go
import (
	"context"
	"net/http"
)

type Handler interface {
	HandleException(ctx context.Context, w http.ResponseWriter, err error)
}
