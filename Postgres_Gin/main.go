package main

import (
	"Postgres_Gin/Functions"
	"Postgres_Gin/api"
	"Postgres_Gin/configs"
	"Postgres_Gin/postgres"
	"Postgres_Gin/server"
	"Postgres_Gin/service"
)

func main() {
	cfg, err := configs.Load()
	Functions.CheckERR(err)

	repo, err := postgres.NewPostgres(cfg.Config)
	Functions.CheckERR(err)

	svc := service.New(repo)

	svr := server.New(svc)

	r := api.InitRouter(svr)
	r.Run("localhost:8080")
}
