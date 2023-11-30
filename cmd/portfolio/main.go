package main

import (
	"github.com/yagoinacio/portfolio-server/config"
	"github.com/yagoinacio/portfolio-server/pkg/api/http"
	"github.com/yagoinacio/portfolio-server/pkg/database/mongodb"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	err = mongodb.OpenConnection()
	if err != nil {
		panic(err)
	}

	server := http.NewServer()

	port := ":" + config.GetServerPort()

	server.Start(port)
}
