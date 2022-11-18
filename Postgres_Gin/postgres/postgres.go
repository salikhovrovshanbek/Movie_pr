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

func (p *Postgres) GetMovie(ctx context.Context) (structs.Movie, error) {

}
