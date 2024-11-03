package services

import (
	"context"
)

type LocationInferrer[T any] interface {
	InferLocation(ctx context.Context) (T, error)
}
