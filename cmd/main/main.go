package main

import (
	"strconv"

	"github.com/abhay676/today-menu/pkg/configs"
	"github.com/abhay676/today-menu/pkg/db"
	"github.com/abhay676/today-menu/pkg/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	host, err := configs.GetEnv("host")
	user, err := configs.GetEnv("user")
	password, err := configs.GetEnv("password")
	dbname, err := configs.GetEnv("dbname")
	portString, err := configs.GetEnv("port")
	port, err := strconv.Atoi(portString)

	_, err = db.ConnectToDatabase(host, user, password, dbname, port)

	router := gin.Default()

	router.Group("/api")
	routes.LoadRouter(router)

	router.Run()
}
