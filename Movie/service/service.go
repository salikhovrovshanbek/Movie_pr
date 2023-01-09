package service

import (
	"Postgres_Gin/Functions"
	"Postgres_Gin/repository"
	"Postgres_Gin/structs"
	"context"
	"fmt"
)

type Service struct {
	repo repository.Repo
}

func New(repo repository.Repo) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) CreateMovie(ctx context.Context, m structs.Movies) error {
	if err := s.repo.CreateMovie(ctx, m); err != nil {
		return err
	}
	return nil
}

func (s Service) CreateGenre(ctx context.Context, g structs.Genres) error {
	if err := s.repo.CreateGenre(ctx, g); err != nil {
		return err
	}
	return nil
}

func (s Service) CreateDirector(ctx context.Context, d structs.Directors) error {
	if err := s.repo.CreateDirector(ctx, d); err != nil {
		return err
	}
	return nil
}

func (s Service) CreateActor(ctx context.Context, a structs.Actors) error {
	if err := s.repo.CreateActor(ctx, a); err != nil {
		return err
	}
	return nil
}

func (s Service) GetMovies(ctx context.Context) ([]structs.Movies, error) {
	str, err := s.repo.GetMovies(ctx)
	er := Functions.CheckERR(err, "getmovie service")
	fmt.Println(er)
	return str, nil
}

func (s Service) GetMovieBy(ctx context.Context) ([]structs.Movie, error) {
	m, err := s.repo.GetMovieBy(ctx)
	if err != nil {
		return []structs.Movie{}, err
	}
	fmt.Println(m)
	return m, nil
}

//
//func (s Service) GetAuthMovies(ctx context.Context, id string) ([]structs.Movie, error) {
//	m, err := s.repo.GetAuthMovies(ctx, id)
//	if err != nil {
//		return []structs.Movie{}, err
//	}
//	return m, nil
//}
//
//func (s Service) CreateMovie(ctx context.Context, movie structs.Movie) error {
//	if err := s.repo.CreateMovie(ctx, movie); err != nil {
//		return err
//	}
//	return nil
//}
//
//func (s Service) CreateAuthor(ctx context.Context) error {
//	if err := s.repo.CreateActor(ctx); err != nil {
//		return err
//	}
//	return nil
//}
