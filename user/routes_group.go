package user

import (
	"github.com/gin-gonic/gin"
	"happy-server/db"
	"happy-server/shared"
	"net/http"
)

func AddRoutesGroup(r *gin.Engine) {
	userRouter := r.Group("/user")

	userRouter.GET("/", func(c *gin.Context) {
		users, err := GetUsers(db.Connection)
		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		c.JSON(http.StatusOK, users)
	})

	userRouter.POST("/", func(c *gin.Context) {
		var ud createUserData
		if validationErr := c.ShouldBindJSON(&ud); validationErr != nil {
			c.JSON(http.StatusBadRequest, shared.RequestError{ErrMsg: validationErr.Error()})
			return
		}
		newUser, newUserErr := CreateUser(db.Connection, ud)
		if newUserErr != nil {
			c.JSON(http.StatusBadRequest, shared.RequestError{ErrMsg: newUserErr.Error()})
			return
		}
		c.JSON(http.StatusCreated, newUser)
	})

	userRouter.GET("/:userId", func(c *gin.Context) {
		userId, parseError := shared.ParseID(c.Param("userId"))
		if parseError != nil {
			c.JSON(http.StatusBadRequest, shared.RequestError{ErrMsg: parseError.Error()})
			return
		}
		foundUser, foundUserErr := GetUser(db.Connection, userId)
		if foundUserErr != nil {
			c.JSON(http.StatusNotFound, shared.RequestError{ErrMsg: foundUserErr.Error()})
			return
		}
		c.JSON(200, foundUser)
	})

	userRouter.PATCH("/:userId", func(c *gin.Context) {
		userId, parseError := shared.ParseID(c.Param("userId"))
		if parseError != nil {
			c.JSON(http.StatusBadRequest, parseError.Error())
			return
		}
		var ud updateUserData
		if validationErr := c.ShouldBindJSON(&ud); validationErr != nil {
			c.JSON(http.StatusBadRequest, shared.RequestError{ErrMsg: validationErr.Error()})
			return
		}
		updatedUser, updateErr := UpdateUser(db.Connection, userId, ud)
		if updateErr != nil {
			c.JSON(http.StatusBadRequest, shared.RequestError{ErrMsg: updateErr.Error()})
			return
		}
		c.JSON(http.StatusOK, updatedUser)
	})

	userRouter.DELETE("/:userId", func(c *gin.Context) {
		userId, parseError := shared.ParseID(c.Param("userId"))
		if parseError != nil {
			c.JSON(http.StatusBadRequest, parseError.Error())
			return
		}
		deleteErr := DeleteUser(db.Connection, userId)
		if deleteErr != nil {
			c.JSON(http.StatusNotFound, shared.RequestError{ErrMsg: deleteErr.Error()})
			return
		}
		c.JSON(http.StatusOK, nil)
	})
}
