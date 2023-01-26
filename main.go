package main

import (
	"fmt"
	"os"
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

	fmt.Printf("Server Listen PORT %s\n", os.Getenv("PORT"))

	if err := r.Run(":" + os.Getenv("PORT")); err != nil {
		fmt.Println(err)
	}
}
