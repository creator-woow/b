package auth

import "github.com/gin-gonic/gin"

func AddRoutesGroup(router *gin.Engine) {
	authRouter := router.Group("/auth")

	authRouter.POST("/registration", func(c *gin.Context) {})
	authRouter.POST("/login", func(c *gin.Context) {})
	authRouter.POST("/refresh", func(c *gin.Context) {})
}
