package structs

import "time"

type Movie struct {
	ID            int
	Title         string
	Author        string
	AuthorID      int
	PublishedData time.Time
}

type Author struct {
	ID        int
	Name      string
	Janr      string
	Movies    []Movie
	BirthData time.Time
	DeathData time.Time
}
