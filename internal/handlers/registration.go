package handlers

import (
	"github.com/gin-gonic/gin"
	"happy-server/internal/delivery/http/server"
	"happy-server/internal/services"
	"net/http"
)

func Registration(as *services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var d services.CreateUserData
		if vErr := c.ShouldBind(&d); vErr != nil {
			c.JSON(http.StatusBadRequest, httpserver.RequestError{ErrMsg: vErr.Error()})
			return
		}
		newUser, cErr := as.Register(d)
		if cErr != nil {
			c.JSON(http.StatusBadRequest, httpserver.RequestError{ErrMsg: cErr.Error()})
			return
		}
		c.JSON(http.StatusOK, newUser)
	}
}
