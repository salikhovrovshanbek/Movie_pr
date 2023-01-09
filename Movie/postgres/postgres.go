package postgres

import (
	"Postgres_Gin/Functions"
	"Postgres_Gin/structs"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
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

func (p *Postgres) CreateMovie(ctx context.Context, s structs.Movies) error {
	query := `INSERT INTO movies(id,directorid,genreid,title,releaseyear,
                   rating,plot,movielength)
                   VALUES ($1,$2,$3,$4,$5,$6,$7)`

	_, err := p.db.ExecContext(ctx, query, s.ID, s.DirectorID,
		s.GenreID, s.Title, s.ReleaseYear, s.Rating, s.Plot, s.MovieLength)
	er := Functions.CheckERR(err, "postgres create movie !")
	return er
}

//func (p Postgres) CreateAuthor(ctx context.Context, s structs.Author) error {
//	//TODO implement me
//	panic("implement me")
//}

func (p *Postgres) CreateGenre(ctx context.Context, s structs.Genres) error {
	query := `INSERT INTO genres(id,genrename)
					VALUES ($1,$2)`

	_, err := p.db.ExecContext(ctx, query, s.ID, s.GenreName)
	er := Functions.CheckERR(err, "postgres create genre")
	return er
}

func (p *Postgres) CreateDirector(ctx context.Context, s structs.Directors) error {
	query := `INSERT INTO directors(id,firtsname,lastname,nationality,birth)
					VALUES ($1,$2,$3,$4,$5)`

	_, err := p.db.ExecContext(ctx, query, s.ID, s.FirstName, s.LastName, s.Nationality, s.Birth)
	er := Functions.CheckERR(err, "postgres create director")
	return er
}

func (p *Postgres) CreateActor(ctx context.Context, s structs.Actors) error {
	query := `INSERT INTO actor(id,firtsname,lastname,nationality,birth)
					VALUES ($1,$2,$3,$4,$5)`

	_, err := p.db.ExecContext(ctx, query, s.ID, s.FirstName, s.LastName, s.Nationality, s.Birth)
	er := Functions.CheckERR(err, "postgres create director")
	return er
}

func (p *Postgres) GetMovies(ctx context.Context) ([]structs.Movies, error) {
	query := `SELECT * FROM movies`

	var list []structs.Movies
	if err := p.db.GetContext(ctx, &list, query); err != nil {
		log.Fatalln("getmoives postgres ", err.Error())
	}
	return list, nil
}

//func (p *Postgres) GetAuthors(ctx context.Context) ([]structs.Author, error) {
//	//TODO implement me
//	panic("implement me")
//}

func (p *Postgres) GetGenres(ctx context.Context) ([]structs.Genres, error) {
	query := `SELECT * FROM genres`

	var list []structs.Genres
	if err := p.db.GetContext(ctx, &list, query); err != nil {
		log.Fatalln("getgenres postgres ", err.Error())
	}
	return list, nil
}

func (p *Postgres) GetDirectors(ctx context.Context) ([]structs.Directors, error) {
	query := `SELECT * FROM directors`

	var list []structs.Directors
	if err := p.db.GetContext(ctx, &list, query); err != nil {
		log.Fatalln("getdirectors postgres ", err.Error())
	}
	return list, nil
}

func (p *Postgres) GetActors(ctx context.Context) ([]structs.Actors, error) {
	query := `SELECT * FROM actors`

	var list []structs.Actors
	if err := p.db.GetContext(ctx, &list, query); err != nil {
		log.Fatalln("getactors postgres ", err.Error())
	}
	return list, nil
}

func (p *Postgres) UpdateMovie(ctx context.Context, m structs.Movies) error {
	query := `
	UPDATE movies 
		SET directorid=$1,genreid=$2,title=$3,releaseyear=$4,rating=$5,plot=$6,movielength=$7
	WHERE id=$8			
	`
	_, err := p.db.ExecContext(
		ctx, query,
		m.DirectorID, m.GenreID, m.Title, m.ReleaseYear, m.Rating, m.Plot, m.MovieLength, m.ID,
	)
	er := Functions.CheckERR(err, "while executing updatemovie postgres pkg")
	return er
}

//func (p *Postgres) UpdateAuthor(ctx context.Context, author structs.Author) error {
//	//TODO implement me
//	panic("implement me")
//}

func (p *Postgres) UpdateGenre(ctx context.Context, s structs.Genres) error {
	query := `
	UPDATE  genres
		SET genrename=$1
	WHERE id=$2
	`
	_, err := p.db.ExecContext(
		ctx, query,
		s.GenreName, s.GenreName,
	)
	er := Functions.CheckERR(err, "while executing in postgrespkg")
	return er
}

func (p *Postgres) UpdateDirector(ctx context.Context, s structs.Directors) error {
	query := `
	UPDATE  directors
		SET	firstname=$1, lastname=$2, nationality=$3, birth=$4
	WHERE id=$5
	`
	_, err := p.db.ExecContext(
		ctx, query,
		s.FirstName, s.LastName, s.Nationality, s.Birth, s.ID,
	)
	er := Functions.CheckERR(err, "while executing in postgrespkg")
	return er
}

func (p *Postgres) UpdateActor(ctx context.Context, s structs.Actors) error {
	query := `
	UPDATE  actors
		SET	firstname=$1, lastname=$2, nationality=$3, birth=$4
	WHERE id=$5
	`
	_, err := p.db.ExecContext(
		ctx, query,
		s.FirstName, s.LastName, s.Nationality, s.Birth, s.ID,
	)
	er := Functions.CheckERR(err, "while executing in postgrespkg")
	return er
}

func (p *Postgres) DeleteMovie(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM movies
           WHERE id=$1`
	_, err := p.db.ExecContext(ctx, query, id)
	er := Functions.CheckERR(err, "deleting  in postgres pkg")
	return er
}

//func (p *Postgres) DeleteAuthor(ctx context.Context, id uuid.UUID) error {
//	query := `DELETE FROM auths
//		WHERE id=$1`
//
//	_, err := p.db.ExecContext(ctx, query, id)
//	er := Functions.CheckERR(err, "deleting  in postgres pkg")
//	return er
//}

func (p *Postgres) DeleteGenre(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM genres
           WHERE id=$1`
	_, err := p.db.ExecContext(ctx, query, id)
	er := Functions.CheckERR(err, "deleting  in postgres pkg")
	return er
}

func (p *Postgres) DeleteDirector(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM directors
           WHERE id=$1`
	_, err := p.db.ExecContext(ctx, query, id)
	er := Functions.CheckERR(err, "deleting  in postgres pkg")
	return er
}

func (p *Postgres) DeleteActor(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM actors
           WHERE id=$1`
	_, err := p.db.ExecContext(ctx, query, id)
	er := Functions.CheckERR(err, "deleting  in postgres pkg")
	return er

}

func (p *Postgres) GetMovieBy(ctx context.Context) ([]structs.Movie, error) {
	query := `SELECT * FROM movie;`

	list := make([]structs.Movie, 0)
	if err := p.db.SelectContext(ctx, &list, query); err != nil {
		fmt.Println(list, err)
		return []structs.Movie{}, err
	}
	fmt.Println(list)
	return list, nil
}

//	func (p *Postgres) CreateActor(ctx context.Context) error {
//		//TODO implement me
//		panic("implement me")
//	}
//
//	func (p *Postgres) CreateDirector(ctx context.Context) error {
//		//TODO implement me
//		panic("implement me")
//	}
//
//	func (p *Postgres) UpdateMovie(ctx context.Context, id string) error {
//		//TODO implement me
//		panic("implement me")
//	}
//
//	func (p *Postgres) UpdateActor(ctx context.Context, id string) error {
//		//TODO implement me
//		panic("implement me")
//	}
//
//	func (p *Postgres) UpdateDirector(ctx context.Context, id string) error {
//		//TODO implement me
//		panic("implement me")
//	}
//
//	func (p *Postgres) DeleteMovie(ctx context.Context, id string) error {
//		//TODO implement me
//		panic("implement me")
//	}
//
//	func (p *Postgres) DeleteAuthor(ctx context.Context, id string) error {
//		//TODO implement me
//		panic("implement me")
//	}
//
//	func (p *Postgres) DeleteDirector(ctx context.Context, id string) error {
//		//TODO implement me
//		panic("implement me")
//	}

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
