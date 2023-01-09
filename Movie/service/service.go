package service

import (
	"Postgres_Gin/Functions"
	"Postgres_Gin/repository"
	"Postgres_Gin/structs"
	"context"
	"fmt"
	"github.com/google/uuid"
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
	m.ID = uuid.New()
	if err := s.repo.CreateMovie(ctx, m); err != nil {
		return err
	}
	return nil
}

func (s Service) CreateGenre(ctx context.Context, g structs.Genres) error {
	g.ID = uuid.New()
	if err := s.repo.CreateGenre(ctx, g); err != nil {
		return err
	}
	return nil
}

func (s Service) CreateDirector(ctx context.Context, d structs.Directors) error {
	d.ID = uuid.New()
	if err := s.repo.CreateDirector(ctx, d); err != nil {
		return err
	}
	return nil
}

func (s Service) CreateActor(ctx context.Context, a structs.Actors) error {
	a.ID = uuid.New()
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

//func (s Service) GetAuthors(ctx context.Context) ([]structs.Author, error) {
//	str, err := s.repo.GetAuthors(ctx)
//	er := Functions.CheckERR(err, "getauth service")
//	fmt.Println(er)
//	return str, nil
//}

func (s Service) GetMovieBy(ctx context.Context) ([]structs.Movie, error) {
	m, err := s.repo.GetMovieBy(ctx)
	if err != nil {
		return []structs.Movie{}, err
	}
	fmt.Println(m)
	return m, nil
}

func (s Service) GetGenres(ctx context.Context) ([]structs.Genres, error) {
	str, err := s.repo.GetGenres(ctx)
	if err != nil {
		er := Functions.CheckERR(err, "get methods in service")
		return []structs.Genres{}, er
	}
	return str, nil
}

func (s Service) GetDirectors(ctx context.Context) ([]structs.Directors, error) {
	str, err := s.repo.GetDirectors(ctx)
	if err != nil {
		er := Functions.CheckERR(err, "get methods in service")
		return []structs.Directors{}, er
	}
	return str, nil
}

func (s Service) GetActors(ctx context.Context) ([]structs.Actors, error) {
	str, err := s.repo.GetActors(ctx)
	if err != nil {
		er := Functions.CheckERR(err, "get methods in service")
		return []structs.Actors{}, er
	}
	return str, nil
}

func (s Service) UpdateMovie(ctx context.Context, str structs.Movies) error {
	if err := s.repo.UpdateMovie(ctx, str); err != nil {
		er := Functions.CheckERR(err, "while updating in service pkg")
		return er
	}
	return nil
}

//func (s Service) UpdateAuthor(ctx context.Context,str structs.Author) error {
//	if err:=s.repo.UpdateAuthor(ctx,str);err!=nil{
//		er:=Functions.CheckERR(err,"while updating in service pkg")
//		return er
//	}
//	return nil
//}

func (s Service) UpdateGenre(ctx context.Context, str structs.Genres) error {
	if err := s.repo.UpdateGenre(ctx, str); err != nil {
		er := Functions.CheckERR(err, "while updating in service pkg")
		return er
	}
	return nil
}

func (s Service) UpdateDirector(ctx context.Context, str structs.Directors) error {
	if err := s.repo.UpdateDirector(ctx, str); err != nil {
		er := Functions.CheckERR(err, "while updating in service pkg")
		return er
	}
	return nil
}

func (s Service) UpdateActor(ctx context.Context, str structs.Actors) error {
	if err := s.repo.UpdateActor(ctx, str); err != nil {
		er := Functions.CheckERR(err, "while updating in service pkg")
		return er
	}
	return nil
}

func (s Service) DeleteMovie(ctx context.Context, id string) error {
	if err := s.repo.DeleteMovie(ctx, id); err != nil {
		er := Functions.CheckERR(err, "deleting methods in service pkg")
		return er
	}
	return nil
}

//func (s Service) DeleteAuthor(ctx context.Context, id string) error {
//	if err := s.repo.DeleteAuthor(ctx, id); err != nil {
//		er := Functions.CheckERR(err, "deleting methods in service pkg")
//		return er
//	}
//	return nil
//}

func (s Service) DeleteGenre(ctx context.Context, id string) error {
	if err := s.repo.DeleteGenre(ctx, id); err != nil {
		er := Functions.CheckERR(err, "deleting methods in service pkg")
		return er
	}
	return nil
}

func (s Service) DeleteDirector(ctx context.Context, id string) error {
	if err := s.repo.DeleteDirector(ctx, id); err != nil {
		er := Functions.CheckERR(err, "deleting methods in service pkg")
		return er
	}
	return nil
}

func (s Service) DeleteActor(ctx context.Context, id string) error {
	if err := s.repo.DeleteActor(ctx, id); err != nil {
		er := Functions.CheckERR(err, "deleting methods in service pkg")
		return er
	}
	return nil
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
