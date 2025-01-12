package handlers

import (
	"github.com/gin-gonic/gin"
	"happy-server/internal/delivery/http/server"
	"happy-server/internal/models"
	"happy-server/internal/services"
	"net/http"
	"strconv"
)

func UpdateUser(s *services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, parseError := strconv.ParseUint(c.Param(UserIdParam), 10, 64)
		if parseError != nil {
			c.JSON(http.StatusBadRequest, httpserver.RequestError{ErrMsg: parseError.Error()})
			return
		}
		var ud services.UpdateUserData
		if validationErr := c.ShouldBindJSON(&ud); validationErr != nil {
			c.JSON(http.StatusBadRequest, httpserver.RequestError{ErrMsg: validationErr.Error()})
			return
		}
		updatedUser, updateErr := s.UpdateUser(models.ID(userId), ud)
		if updateErr != nil {
			c.JSON(http.StatusBadRequest, httpserver.RequestError{ErrMsg: updateErr.Error()})
			return
		}
		c.JSON(http.StatusOK, updatedUser)
	}

}
