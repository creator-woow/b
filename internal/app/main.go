package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"happy-server/internal/db"
	"happy-server/internal/initialize"
	"happy-server/internal/migrations"
	"happy-server/internal/services"
	"log"
)

func Run() {
	envConfig := initialize.Env()
	newDb := db.New(
		envConfig.DBUser,
		envConfig.DBPass,
		envConfig.DBHost,
		envConfig.DBPort,
		envConfig.DBName,
	)
	migrations.AutoMigrations(newDb)
	r := gin.Default()
	userService := &services.UserService{DB: newDb}
	tokenService := &services.TokenService{
		DB:        newDb,
		JWTSecret: envConfig.JwtSecret,
	}
	authService := &services.AuthService{
		DB: newDb, UserService: userService,
		TokenService: tokenService,
	}
	initialize.UserHandlers(r, userService)
	initialize.AuthHandlers(r, authService)
	runErr := r.Run(fmt.Sprintf(":%v", envConfig.AppPort))
	if runErr != nil {
		log.Panicln("error while starting server: ", runErr.Error())
	}
}
