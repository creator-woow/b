package initialize

import (
	"github.com/gin-gonic/gin"
	"happy-server/internal/handlers"
	"happy-server/internal/services"
)

func AuthHandlers(r *gin.Engine, s *services.AuthService) {
	g := r.Group("/auth")
	g.POST("/registration", handlers.Registration(s))
	g.POST("/login", handlers.Login(s))
}
