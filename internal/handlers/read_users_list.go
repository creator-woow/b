package handlers

import (
	"github.com/gin-gonic/gin"
	"happy-server/internal/services"
	"net/http"
)

func ReadUsersList(s *services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := s.ReadUsers()
		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		c.JSON(http.StatusOK, users)
	}
}
