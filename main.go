package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"happy-server/auth"
	"happy-server/db"
	"happy-server/user"
)

func main() {
	initEnv()

	r := gin.Default()
	db.Connection = db.NewConnection()

	runAutoMigrations()
	extendMainRouter(r)

	r.Run()
}

func initEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
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
	user.AddRoutesGroup(r)
	auth.AddRoutesGroup(r)
}
