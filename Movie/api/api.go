package api

import (
	"Postgres_Gin/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(s server.Server) *gin.Engine {
	r := gin.Default()

	r.POST("/movie")
	//r.POST("/author")
	r.POST("/genre")
	r.POST("/director")
	r.POST("/actor")

	r.GET("/movie")
	//r.GET("/authors")
	r.GET("/genres")
	r.GET("/directors")
	r.GET("/actor")

	r.PUT("/movie")
	//r.PUT("/author")
	r.PUT("/genre")
	r.PUT("/director")
	r.PUT("/actor")

	r.DELETE("/movie/:id")
	//r.DELETE("/authors/:id")
	r.DELETE("/genres/:id")
	r.DELETE("/directors/:id")
	r.DELETE("/actor")

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
