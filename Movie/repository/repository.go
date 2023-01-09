package repository

import (
	"Postgres_Gin/structs"
	"context"
	"github.com/google/uuid"
)

type Repo interface {
	CreateMovie(ctx context.Context, s structs.Movies) error
	//CreateAuthor(ctx context.Context,s structs.Author)error
	CreateGenre(ctx context.Context, s structs.Genres) error
	CreateDirector(ctx context.Context, s structs.Directors) error
	CreateActor(ctx context.Context, s structs.Actors) error

	GetMovies(ctx context.Context) ([]structs.Movies, error)
	//GetAuthors(ctx context.Context)([]structs.Author,error)
	GetGenres(ctx context.Context) ([]structs.Genres, error)
	GetDirectors(ctx context.Context) ([]structs.Directors, error)
	GetActors(ctx context.Context) ([]structs.Actors, error)

	UpdateMovie(ctx context.Context, movie structs.Movies) error
	//UpdateAuthor(ctx context.Context,author structs.Author)error
	UpdateGenre(ctx context.Context, genres structs.Genres) error
	UpdateDirector(ctx context.Context, directors structs.Directors) error
	UpdateActor(ctx context.Context, actors structs.Actors) error

	DeleteMovie(ctx context.Context, id uuid.UUID) error
	DeleteAuthor(ctx context.Context, id uuid.UUID) error
	DeleteGenre(ctx context.Context, id uuid.UUID) error
	DeleteDirector(ctx context.Context, id uuid.UUID) error
	DeleteActor(ctx context.Context, id uuid.UUID) error

	GetMovieBy(ctx context.Context) ([]structs.Movie, error)
	//GetAuthMovies(ctx context.Context, id string) ([]structs.Movie, error)
	//CreateMovie(ctx context.Context, movie structs.Movie) error
	//CreateActor(ctx context.Context) error
	//CreateDirector(ctx context.Context) error
	//UpdateMovie(ctx context.Context, id string) error
	//UpdateActor(ctx context.Context, id string) error
	//UpdateDirector(ctx context.Context, id string) error
	//DeleteMovie(ctx context.Context, id string) error
	//DeleteAuthor(ctx context.Context, id string) error
	//DeleteDirector(ctx context.Context, id string) error
}
