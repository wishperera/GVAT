package exception

import (
	"context"
	"net/http"
)

type Handler interface {
	HandleException(ctx context.Context, w http.ResponseWriter, err error)
}
