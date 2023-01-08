package service

import (
	"Postgres_Gin/repository"
)

type Service struct {
	repo repository.Repo
}

func New(repo repository.Repo) Service {
	return Service{
		repo: repo,
	}
}

//
//func (s Service) GetMovieBy(ctx context.Context) ([]structs.Movie, error) {
//	m, err := s.repo.GetMovieBy(ctx)
//	if err != nil {
//		return []structs.Movie{}, err
//	}
//	return m, nil
//}
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
