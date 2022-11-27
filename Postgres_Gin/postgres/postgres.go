package postgres

import (
	"Postgres_Gin/structs"
	"context"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgres(cfg Config) (*Postgres, error) {
	db, err := Connect(cfg)
	if err != nil {
		return nil, err
	}

	return &Postgres{
		db: db,
	}, nil
}

type Postgres struct {
	db *sqlx.DB
}

func (p Postgres) GetMovieBy(ctx context.Context) ([]structs.Movie, error) {
	query := `SELECT * FROM movie;`

	var m structs.Movie
	list := make([]structs.Movie, 0)
	if err := p.db.SelectContext(ctx, &m, query); err != nil {
		return []structs.Movie{}, err
	}
	return list, nil
}

//func (p Postgres) GetMovies(ctx context.Context, id string) ([]structs.Movie, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (p Postgres) CreateMovie(ctx context.Context, movie structs.Movie) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (p Postgres) CreateAuthor(ctx context.Context, author structs.Author) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (p Postgres) UpdateMovie(ctx context.Context, id string) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (p Postgres) UpdateAuthor(ctx context.Context, id string) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (p Postgres) DeleteMovie(ctx context.Context, id string) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (p Postgres) DeleteAuthor(ctx context.Context, id string) error {
//	//TODO implement me
//	panic("implement me")
//}
