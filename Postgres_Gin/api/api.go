package api

import (
	"Postgres_Gin/server"
	"github.com/gin-gonic/gin"
)

func InitRouter(s server.Server) *gin.Engine {
	r := gin.Default()

	r.GET("/getmovie", s.GetMovieBy)

	return r
}
