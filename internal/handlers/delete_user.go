package handlers

import (
	"github.com/gin-gonic/gin"
	"happy-server/internal/delivery/http/server"
	"happy-server/internal/models"
	"happy-server/internal/services"
	"net/http"
	"strconv"
)

func DeleteUser(s *services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, parseError := strconv.ParseUint(c.Param(UserIdParam), 10, 64)
		if parseError != nil {
			c.JSON(http.StatusBadRequest, httpserver.RequestError{ErrMsg: parseError.Error()})
			return
		}
		deleteErr := s.DeleteUser(models.ID(userId))
		if deleteErr != nil {
			c.JSON(http.StatusNotFound, httpserver.RequestError{ErrMsg: deleteErr.Error()})
			return
		}
		c.JSON(http.StatusOK, nil)
	}
}
