package handlers

import (
	"github.com/gin-gonic/gin"
	"happy-server/internal/delivery/http/server"
	"happy-server/internal/services"
	"net/http"
)

func Login(s *services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var d services.LoginData
		if err := c.ShouldBind(&d); err != nil {
			c.JSON(http.StatusBadRequest, httpserver.RequestError{ErrMsg: err.Error()})
			return
		}
		tokens, err := s.Login(d)
		if err != nil {
			c.JSON(http.StatusBadRequest, httpserver.RequestError{ErrMsg: err.Error()})
			return
		}
		c.SetCookie(
			"accessToken",
			tokens.AccessToken,
			int((services.AccessTokenExpireTime).Seconds()),
			"/",
			"",
			true,
			true,
		)
		c.SetCookie(
			"refreshToken",
			tokens.RefreshToken,
			int((services.RefreshTokenExpireTime).Seconds()),
			"/",
			"",
			true,
			true,
		)
		c.JSON(http.StatusOK, tokens)
	}
}
