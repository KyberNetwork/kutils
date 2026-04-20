package tracer

import (
	"context"

	"github.com/KyberNetwork/kyber-trace-go/pkg/tracer"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type Span struct {
	span trace.Span
}

func (s Span) SetTag(name string, value string) {
	s.span.SetAttributes(attribute.String(name, value))
}

func (s Span) End() {
	s.span.End()
}

func StartSpanFromContext(ctx context.Context, operationName string) (Span, context.Context) {
	ctx, span := tracer.Tracer().Start(ctx, operationName)
	return Span{
		span: span,
	}, ctx
}
