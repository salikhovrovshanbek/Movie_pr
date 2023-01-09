package server

import (
	"Postgres_Gin/Functions"
	"Postgres_Gin/service"
	"Postgres_Gin/structs"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	svc service.Service
}

func New(service2 service.Service) Server {
	return Server{
		svc: service2,
	}
}

func (s Server) CreateMovie(c *gin.Context) {
	var m structs.Movies
	if err := c.ShouldBindJSON(m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ok":    false,
			"error": err,
		})
	}
	if err := s.svc.CreateMovie(c, m); err != nil {
		Functions.CheckSERVERErr(err, c)
		return
	}
	Functions.SERVEROKAY(c)
	return
}

func (s Server) CreateGenre(c *gin.Context) {
	var m structs.Genres
	if err := c.ShouldBindJSON(m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ok":    false,
			"error": err,
		})
	}
	if err := s.svc.CreateGenre(c, m); err != nil {
		Functions.CheckSERVERErr(err, c)
		return
	}
	Functions.SERVEROKAY(c)
	return
}

func (s Server) CreateDirector(c *gin.Context) {
	var m structs.Directors
	if err := c.ShouldBindJSON(m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ok":    false,
			"error": err,
		})
	}
	if err := s.svc.CreateDirector(c, m); err != nil {
		Functions.CheckSERVERErr(err, c)
		return
	}
	Functions.SERVEROKAY(c)
	return
}

func (s Server) CreateActor(c *gin.Context) {
	var m structs.Actors
	if err := c.ShouldBindJSON(m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ok":    false,
			"error": err,
		})
	}
	if err := s.svc.CreateActor(c, m); err != nil {
		Functions.CheckSERVERErr(err, c)
		return
	}
	Functions.SERVEROKAY(c)
	return
}

func (s Server) GetMovies(c *gin.Context) ([]structs.Movies, error) {
	str, err := s.svc.GetMovies(c)
	if err != nil {
		Functions.CheckSERVERErr(err, c)
		return []structs.Movies{}, err
	}
	Functions.SERVEROKDATA(str, c)
	return str, nil
}

func (s Server) GetMovieBy(c *gin.Context) {
	m, err := s.svc.GetMovieBy(c)
	if err != nil {
		Functions.CheckSERVERErr(err, c)
		return
	}

	Functions.SERVEROKDATA(m, c)
}

//
//func (s Server) GetAuthMovies(c *gin.Context) {
//	id := c.Param("id")
//
//	m, err := s.svc.GetAuthMovies(c, id)
//	if err != nil {
//		Functions.CheckSERVERErr(err, c)
//		return
//
//	}
//	Functions.SERVEROKDATA(m, c)
//}
//
//func (s Server) CreateMovie(c *gin.Context) {
//	var m structs.Movie
//
//	if err := c.ShouldBindJSON(m); err != nil {
//		Functions.CheckERR(err)
//		return
//	}
//
//	if err := s.svc.CreateMovie(c, m); err != nil {
//		Functions.CheckSERVERErr(err, c)
//		return
//	}
//
//	Functions.SERVEROKAY(c)
//}
//
//func (s Server) CreateAuthor(c *gin.Context) {
//	var a structs.Author
//
//	if err := c.ShouldBindJSON(a); err != nil {
//		Functions.CheckERR(err)
//		return
//	}
//
//	if err := s.svc.CreateAuthor(c); err != nil {
//		Functions.CheckSERVERErr(err, c)
//		return
//	}
//
//	Functions.SERVEROKAY(c)
//}
