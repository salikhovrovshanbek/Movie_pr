package api

import (
	"Postgres_Gin/server"
	"github.com/gin-gonic/gin"
)

func InitRouter(s server.Server) *gin.Engine {
	r := gin.Default()

	r.GET("/getmovie", s.GetMovieBy)
	r.GET("/getauth/:id", s.GetAuthMovies)
	r.GET("/newmovie", s.CreateMovie)
	r.GET("/newauth", s.CreateAuthor)

	return r
}
