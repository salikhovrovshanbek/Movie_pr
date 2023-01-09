package structs

import (
	"github.com/google/uuid"
	"time"
)

type Movie struct {
	ID            int       `json:"id"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	AuthorID      int       `json:"authorid"`
	PublishedData time.Time `json:"publishedData"`
}

type Author struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Janr      string    `json:"janr"`
	Movies    []int     `json:"movies"`
	BirthData time.Time `json:"birth_data"`
	DeathData time.Time `json:"death_data"`
}

type Genres struct {
	ID        uuid.UUID `json:"id"`
	GenreName string    `json:"genrename"`
}

type MoiveGenres struct {
	MoiveID uuid.UUID `json:"moiveid"`
	GenreID uuid.UUID `json:"genreID"`
}

type Movies struct {
	ID          uuid.UUID `json:"id"`
	DirectorID  uuid.UUID `json:"directorid"`
	GenreID     uuid.UUID `json:"genreid"`
	Title       string    `json:"title"`
	ReleaseYear int       `json:"releaseyear"`
	Rating      int       `json:"rating"`
	Plot        int       `json:"plot"`
	MovieLength time.Time `json:"movielength"`
}

type MovieActor struct {
	MovieID uuid.UUID `json:"movieid"`
	ActorID uuid.UUID `json:"actorid"`
}

type Actors struct {
	ID          uuid.UUID `json:"id"`
	FirstName   string    `json:"firstname"`
	LastName    string    `json:"lastname"`
	Nationality string    `json:"nationality"`
	Birth       time.Time `json:"birth"`
}

type Directors struct {
	ID          uuid.UUID `json:"id"`
	FirstName   string    `json:"firstname"`
	LastName    string    `json:"lastname"`
	Nationality string    `json:"nationality"`
	Birth       time.Time `json:"birth"`
}
