package context

import (
	"context"
	"github.com/wishperera/GVAT/gen/internal/domain"
	"github.com/wishperera/GVAT/gen/internal/pkg/uuid"
)

func ExtractTrace(ctx context.Context) string {
	trace, ok := ctx.Value(domain.ContextKeyTraceID).(uuid.UUID)
	if !ok {
		return uuid.New().String()
	}

	return trace.String()
}
