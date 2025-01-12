package handlers

import (
	"github.com/gin-gonic/gin"
	"happy-server/internal/delivery/http/server"
	"happy-server/internal/models"
	"happy-server/internal/services"
	"net/http"
	"strconv"
)

const UserIdParam = "userId"

func ReadUser(s *services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, parseError := strconv.ParseUint(c.Param(UserIdParam), 10, 64)
		if parseError != nil {
			c.JSON(http.StatusBadRequest, httpserver.RequestError{ErrMsg: parseError.Error()})
			return
		}
		foundUser, foundUserErr := s.ReadUserById(models.ID(userId))
		if foundUserErr != nil {
			c.JSON(http.StatusNotFound, httpserver.RequestError{ErrMsg: foundUserErr.Error()})
			return
		}
		c.JSON(200, foundUser)
	}
}
