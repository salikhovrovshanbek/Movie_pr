package server

import (
	"Postgres_Gin/Functions"
	"Postgres_Gin/service"
	"Postgres_Gin/structs"
	"github.com/gin-gonic/gin"
)

type Server struct {
	svc service.Service
}

func New(service2 service.Service) Server {
	return Server{
		svc: service2,
	}
}

func (s Server) GetMovieBy(c *gin.Context) {
	m, err := s.svc.GetMovieBy(c)
	if err != nil {
		Functions.CheckSERVERErr(err, c)
		return
	}

	Functions.SERVEROKDATA(m, c)
}

func (s Server) GetAuthMovies(c *gin.Context) {
	id := c.Param("id")

	m, err := s.svc.GetAuthMovies(c, id)
	if err != nil {
		Functions.CheckSERVERErr(err, c)
		return

	}
	Functions.SERVEROKDATA(m, c)
}

func (s Server) CreateMovie(c *gin.Context) {
	var m structs.Movie

	if err := c.ShouldBindJSON(m); err != nil {
		Functions.CheckERR(err)
		return
	}

	if err := s.svc.CreateMovie(c, m); err != nil {
		Functions.CheckSERVERErr(err, c)
		return
	}

	Functions.SERVEROKAY(c)
}

func (s Server) CreateAuthor(c *gin.Context) {
	var a structs.Author

	if err := c.ShouldBindJSON(a); err != nil {
		Functions.CheckERR(err)
		return
	}

	if err := s.svc.CreateAuthor(c, a); err != nil {
		Functions.CheckSERVERErr(err, c)
		return
	}

	Functions.SERVEROKAY(c)
}
