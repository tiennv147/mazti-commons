package endpoints

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)

type RequestWithID struct {
	ID int64
}

type ListRequest struct {
	Limit  int64
	Offset int64
}

// LoggingMiddleware returns an endpoint middleware that logs the
// duration of each invocation, and the resulting error, if any.
//noinspection GoUnhandledErrorResult
func LoggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {

			defer func(begin time.Time) {
				logger.Log("transport_error", err, "took", time.Since(begin))
			}(time.Now())
			return next(ctx, request)

		}
	}
}
