package main

import (
	"github.com/PaguhEsatrio/task-5-pbi-btpns-PaguhEsatrio/router"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	r := router.SetupRouter()
	r.Run()
}
