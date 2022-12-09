package repository

import (
	"Postgres_Gin/structs"
	"context"
)

type Repo interface {
	GetMovieBy(ctx context.Context) ([]structs.Movie, error)
	GetAuthMovies(ctx context.Context, id string) ([]structs.Movie, error)
	CreateMovie(ctx context.Context, movie structs.Movie) error
	CreateActor(ctx context.Context) error
	CreateDirector(ctx context.Context) error
	UpdateMovie(ctx context.Context, id string) error
	UpdateActor(ctx context.Context, id string) error
	DeleteMovie(ctx context.Context, id string) error
	DeleteAuthor(ctx context.Context, id string) error
}
