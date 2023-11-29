package main

import (
	"github.com/yagoinacio/portfolio-server/config"
	"github.com/yagoinacio/portfolio-server/internal/database/mongodb"
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
}
