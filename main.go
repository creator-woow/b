package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"happy-server/auth"
	"happy-server/db"
	"happy-server/user"
)

func init() {
	loadEnv()
	db.Connection = db.NewConnection()
	runAutoMigrations()
}

func main() {
	r := gin.Default()
	extendMainRouter(r)
	r.Run()
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
}

func runAutoMigrations() {
	err := db.Connection.AutoMigrate(&user.User{})
	if err != nil {
		panic("Some of auto migrations failed: " + err.Error())
	}
}

func extendMainRouter(r *gin.Engine) {
	user.InitRoutes(r)
	auth.InitRoutes(r)
}
