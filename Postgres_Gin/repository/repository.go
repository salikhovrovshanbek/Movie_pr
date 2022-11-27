package repository

import (
	"Postgres_Gin/structs"
	"context"
)

type Repo interface {
	GetMovieBy(ctx context.Context) ([]structs.Movie, error)
	//GetMovies(ctx context.Context, id string) ([]structs.Movie, error)
	//CreateMovie(ctx context.Context, movie structs.Movie) error
	//CreateAuthor(ctx context.Context, author structs.Author) error
	//UpdateMovie(ctx context.Context, id string) error
	//UpdateAuthor(ctx context.Context, id string) error
	//DeleteMovie(ctx context.Context, id string) error
	//DeleteAuthor(ctx context.Context, id string) error
}
