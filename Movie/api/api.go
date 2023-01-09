package api

import (
	"Postgres_Gin/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(s server.Server) *gin.Engine {
	r := gin.Default()

	r.POST("/movie", s.CreateMovie)
	//r.POST("/author",s.CreateAuthor)
	r.POST("/genre", s.CreateGenre)
	r.POST("/director", s.CreateDirector)
	r.POST("/actor", s.CreateActor)

	r.GET("/movie", s.GetMovies)
	//r.GET("/authors")
	r.GET("/genres", s.GetGenres)
	r.GET("/directors", s.GetDirectors)
	r.GET("/actor", s.GetActors)

	r.PUT("/movie", s.UpdateMovie)
	//r.PUT("/author")
	r.PUT("/genre", s.UpdateGenre)
	r.PUT("/director", s.UpdateDirector)
	r.PUT("/actor", s.UpdateActor)

	r.DELETE("/movie/:id", s.DeleteMovie)
	//r.DELETE("/authors/:id")
	r.DELETE("/genres/:id", s.DeleteGenre)
	r.DELETE("/directors/:id", s.DeleteDirector)
	r.DELETE("/actor/:id", s.DeleteActor)

	r.GET("/getmovie", s.GetMovieBy)
	//r.GET("/getauth/:id", s.GetAuthMovies)
	//r.GET("/newmovie", s.CreateMovie)
	//r.GET("/newauth", s.CreateAuthor)
	//
	//r.POST("/createmovie")
	//r.POST("/createauthor")
	//r.GET("/getmovie/:smth")
	//r.GET("/getmovies/:authid")
	//r.GET("/getauthors")
	//r.DELETE("/deletemovie/:id")
	//r.DELETE("/deleteauthor/:id")
	//r.PATCH("/updatemovie/:id")
	//r.PATCH("/updateauthor/:id")
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "How are you")
	})
	return r
}
