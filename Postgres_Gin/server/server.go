package server

import (
	"Postgres_Gin/service"
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

func (s Server) GetMovieBy(c *gin.Context) {
	m, err := s.svc.GetMovieBy(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
			"ok":    false,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"ok":   true,
		"data": m,
	})
}
