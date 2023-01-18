package main

import (
	"pontakorn322/demo_gin_api/repository"
	router "pontakorn322/demo_gin_api/routers"
	"pontakorn322/demo_gin_api/utils"
)

func main() {
	err := repository.ConnectDB()
	if err != nil {
		panic("failed to connect database")
	}
	repository.MigrationDB()

	r := router.SetUpRouter()

	router.SetUpRouteGroup(r)
	config := utils.ReadConfigs()

	r.Run(config.Port)
}
