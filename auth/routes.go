package auth

import (
	"github.com/gin-gonic/gin"
	"happy-server/db"
	"happy-server/lib"
	"happy-server/user"
	"net/http"
	"time"
)

func InitRoutes(r *gin.Engine) {
	g := r.Group("/auth")
	g.POST("/registration", register)
	g.POST("/login", login)
	g.POST("/logout", logout)
	g.POST("/token/refresh", refreshTokens)
}

func register(c *gin.Context) {
	var d user.CreateUserData
	if vErr := c.ShouldBind(&d); vErr != nil {
		c.JSON(http.StatusBadRequest, lib.RequestError{ErrMsg: vErr.Error()})
		return
	}
	newUser, cErr := Register(db.Connection, d)
	if cErr != nil {
		c.JSON(http.StatusBadRequest, lib.RequestError{ErrMsg: cErr.Error()})
		return
	}
	c.JSON(http.StatusOK, newUser)
}

func login(c *gin.Context) {
	var d loginData
	if err := c.ShouldBind(&d); err != nil {
		c.JSON(http.StatusBadRequest, lib.RequestError{ErrMsg: err.Error()})
		return
	}
	tokens, err := Login(db.Connection, d)
	if err != nil {
		c.JSON(http.StatusBadRequest, lib.RequestError{ErrMsg: err.Error()})
		return
	}
	c.SetCookie(
		"accessToken",
		tokens.AccessToken,
		int((time.Hour * 24).Seconds()),
		"/",
		"",
		true,
		true,
	)
	c.SetCookie(
		"refreshToken",
		tokens.RefreshToken,
		int((time.Hour * 24 * 30).Seconds()),
		"/",
		"",
		true,
		true,
	)
	c.JSON(http.StatusOK, tokens)
}

func logout(c *gin.Context) {
}

func refreshTokens(c *gin.Context) {

}
