package main

import (
	"santrikoding/backend-api/config"
	"santrikoding/backend-api/database"
	"santrikoding/backend-api/routes"
)

func main() {
	//load config .env
	config.LoadEnv()

	//insialisasi database
	database.InitDB()

	//setup router
	r := routes.SetupRouter()

	//mulai server dengan port 3000

	//mulai server
	r.Run(":" + config.GetEnv("APP_PORT", "3000"))
}
