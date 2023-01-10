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

// 					CREATE

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
		return
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

//					GET

func (s Server) GetMovies(c *gin.Context) {
	str, err := s.svc.GetMovies(c)
	if err != nil {
		Functions.CheckSERVERErr(err, c)
	}
	Functions.SERVEROKDATA(str, c)
}

func (s Server) GetMovieBy(c *gin.Context) {
	m, err := s.svc.GetMovieBy(c)
	if err != nil {
		Functions.CheckSERVERErr(err, c)
		return
	}

	Functions.SERVEROKDATA(m, c)
}

//func (s Server) GetAuthors(c *gin.Context) {
//	str, err := s.svc.GetAuthors(c)
//	if err != nil {
//		Functions.CheckSERVERErr(err, c)
//	}
//	Functions.SERVEROKDATA(str, c)
//
//}

func (s Server) GetGenres(c *gin.Context) {
	str, err := s.svc.GetGenres(c)
	if err != nil {
		Functions.CheckSERVERErr(err, c)
	}
	Functions.SERVEROKDATA(str, c)

}

func (s Server) GetDirectors(c *gin.Context) {
	str, err := s.svc.GetDirectors(c)
	if err != nil {
		Functions.CheckSERVERErr(err, c)
	}
	Functions.SERVEROKDATA(str, c)

}

func (s Server) GetActors(c *gin.Context) {
	str, err := s.svc.GetActors(c)
	if err != nil {
		Functions.CheckSERVERErr(err, c)
	}
	Functions.SERVEROKDATA(str, c)

}

//					UPDATE

func (s Server) UpdateMovie(c *gin.Context) {
	var str structs.Movies
	if err := c.ShouldBindJSON(str); err != nil {
		Functions.CheckERROR(err, "Updating methods in server pkg")
	}
	err := s.svc.UpdateMovie(c, str)
	if err != nil {
		Functions.CheckSERVERErr(err, c)
	}
	Functions.SERVEROKAY(c)
}

//func (s Server) UpdateAuhtor(c *gin.Context)  {
//	var str structs.Author
//	if err:=c.ShouldBindJSON(str);err!=nil{
//		Functions.CheckERROR(err,"Updating methods in server pkg")
//	}
//	err:=s.svc.UpdateAuthor(c,str)
//	if err!=nil{
//		Functions.CheckSERVERErr(err,c)
//	}
//	Functions.SERVEROKAY(c)
//}

func (s Server) UpdateGenre(c *gin.Context) {
	var str structs.Genres
	if err := c.ShouldBindJSON(str); err != nil {
		Functions.CheckERROR(err, "Updating methods in server pkg")
	}
	err := s.svc.UpdateGenre(c, str)
	if err != nil {
		Functions.CheckSERVERErr(err, c)
	}
	Functions.SERVEROKAY(c)
}

func (s Server) UpdateDirector(c *gin.Context) {
	var str structs.Directors
	if err := c.ShouldBindJSON(str); err != nil {
		Functions.CheckERROR(err, "Updating methods in server pkg")
	}
	err := s.svc.UpdateDirector(c, str)
	if err != nil {
		Functions.CheckSERVERErr(err, c)
	}
	Functions.SERVEROKAY(c)
}

func (s Server) UpdateActor(c *gin.Context) {
	var str structs.Actors
	if err := c.ShouldBindJSON(str); err != nil {
		Functions.CheckERROR(err, "Updating methods in server pkg")
	}
	err := s.svc.UpdateActor(c, str)
	if err != nil {
		Functions.CheckSERVERErr(err, c)
	}
	Functions.SERVEROKAY(c)
}

// 					DELETE

func (s Server) DeleteMovie(c *gin.Context) {
	id := c.Param("id")

	if err := s.svc.DeleteMovie(c, id); err != nil {
		Functions.CheckSERVERErr(err, c)
	}
	Functions.SERVEROKAY(c)
}

//func (s Server) DeleteAuthor(c *gin.Context)  {
//id:=c.Param("id")
//	if err:=s.svc.DeleteAuthor(c,id);err!=nil{
//		Functions.CheckSERVERErr(err,c)
//	}
//	Functions.SERVEROKAY(c)
//}

func (s Server) DeleteGenre(c *gin.Context) {
	id := c.Param("id")

	if err := s.svc.DeleteGenre(c, id); err != nil {
		Functions.CheckSERVERErr(err, c)
	}
	Functions.SERVEROKAY(c)
}

func (s Server) DeleteDirector(c *gin.Context) {
	id := c.Param("id")

	if err := s.svc.DeleteDirector(c, id); err != nil {
		Functions.CheckSERVERErr(err, c)
	}
	Functions.SERVEROKAY(c)
}

func (s Server) DeleteActor(c *gin.Context) {
	id := c.Param("id")

	if err := s.svc.DeleteActor(c, id); err != nil {
		Functions.CheckSERVERErr(err, c)
	}
	Functions.SERVEROKAY(c)
}

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
