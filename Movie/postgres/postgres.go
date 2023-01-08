package postgres

import (
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

//
//func (p *Postgres) CreateActor(ctx context.Context) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (p *Postgres) CreateDirector(ctx context.Context) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (p *Postgres) UpdateMovie(ctx context.Context, id string) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (p *Postgres) UpdateActor(ctx context.Context, id string) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (p *Postgres) UpdateDirector(ctx context.Context, id string) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (p *Postgres) DeleteMovie(ctx context.Context, id string) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (p *Postgres) DeleteAuthor(ctx context.Context, id string) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (p *Postgres) DeleteDirector(ctx context.Context, id string) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (p *Postgres) GetMovieBy(ctx context.Context) ([]structs.Movie, error) {
//	query := `SELECT * FROM movie;`
//
//	list := make([]structs.Movie, 0)
//	if err := p.db.SelectContext(ctx, &list, query); err != nil {
//		return []structs.Movie{}, err
//	}
//	return list, nil
//}
//
//func (p *Postgres) GetAuthMovies(ctx context.Context, id string) ([]structs.Movie, error) {
//	query := `SELECT * FROM movies where authorid=$1;`
//
//	list := make([]structs.Movie, 0)
//	if err := p.db.SelectContext(ctx, &list, query, id); err != nil {
//		return []structs.Movie{}, err
//	}
//	return list, nil
//}
//func (p *Postgres) CreateMovie(ctx context.Context, m structs.Movie) error {
//	query := `INSERT INTO movie (id,title,author,auhtorid,publisheddata)
//			VALUES ($1,$2,$3,$4,$5)
//	`
//
//	_, err := p.db.ExecContext(
//		ctx, query, m.ID, m.Title, m.Author, m.AuthorID, m.PublishedData,
//	)
//	if err != nil {
//		log.Fatal(err, "executing context in postgres!")
//		return err
//	}
//	return nil
//}
//
//func (p *Postgres) CreateAuthor(ctx context.Context, a structs.Author) error {
//	query := `INSERT INTO author (id,name,genre,movies,birthdata,deathdata)
//			VALUES ($1,$2,$3,$4,$5,$6)
//	`
//
//	_, err := p.db.ExecContext(
//		ctx, query, a.ID, a.Name, a.Janr, a.Movies, a.BirthData, a.DeathData,
//	)
//	if err != nil {
//		log.Fatal(err, "executing context in postgres!")
//		return err
//	}
//	return nil
//}
//
////
////func (p Postgres) UpdateMovie(ctx context.Context, id string) error {
////	//TODO implement me
////	panic("implement me")
////}
////
////func (p Postgres) UpdateAuthor(ctx context.Context, id string) error {
////	//TODO implement me
////	panic("implement me")
////}
////
////func (p Postgres) DeleteMovie(ctx context.Context, id string) error {
////	//TODO implement me
////	panic("implement me")
////}
////
////func (p Postgres) DeleteAuthor(ctx context.Context, id string) error {
////	//TODO implement me
////	panic("implement me")
////}
