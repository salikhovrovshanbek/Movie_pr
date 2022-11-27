package service

import (
	"Postgres_Gin/repository"
	"Postgres_Gin/structs"
	"context"
)

type Service struct {
	repo repository.Repo
}

func New(repo repository.Repo) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) GetMovieBy(ctx context.Context) ([]structs.Movie, error) {
	m, err := s.repo.GetMovieBy(ctx)
	if err != nil {
		return []structs.Movie{}, err
	}
	return m, nil
}
