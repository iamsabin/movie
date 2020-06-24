package imdb

import (
	"context"
)

type Repository interface {
	GetMovies(ctx context.Context, cursor string, num int64)
}
