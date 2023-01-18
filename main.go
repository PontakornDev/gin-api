package main

import (
	"pontakorn322/demo_gin_api/repository"
	router "pontakorn322/demo_gin_api/routers"
)

func main() {
	err := repository.ConnectDB()
	if err != nil {
		panic("failed to connect database")
	}
	repository.MigrationDB()

	r := router.SetUpRouter()

	router.SetUpRouteGroup(r)

	r.Run(":3333")
}
