package main

import (
	"Postgres_Gin/configs"
	"Postgres_Gin/neccessaryFuncs"
	"Postgres_Gin/postgres"
	"fmt"
)

func main() {
	cfg, err := configs.Load()
	neccessaryFuncs.CheckERR(err)

	repo, err := postgres.NewPostgres(cfg.Config)
	neccessaryFuncs.CheckERR(err)
	fmt.Println(repo)
}
