package user

import (
	"github.com/gin-gonic/gin"
	"happy-server/db"
	"happy-server/lib"
	"net/http"
)

func InitRoutes(r *gin.Engine) {
	g := r.Group("/user")
	g.GET("/", getUsers)
	g.GET("/:userId", getUser)
	g.PATCH("/:userId", patchUser)
	g.DELETE("/:userId", deleteUser)
}

func getUsers(c *gin.Context) {
	users, err := ReadUsers(db.Connection)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, users)
}

func getUser(c *gin.Context) {
	userId, parseError := lib.ParseID(c.Param("userId"))
	if parseError != nil {
		c.JSON(http.StatusBadRequest, lib.RequestError{ErrMsg: parseError.Error()})
		return
	}
	foundUser, foundUserErr := ReadUserById(db.Connection, userId)
	if foundUserErr != nil {
		c.JSON(http.StatusNotFound, lib.RequestError{ErrMsg: foundUserErr.Error()})
		return
	}
	c.JSON(200, foundUser)
}

func patchUser(c *gin.Context) {
	userId, parseError := lib.ParseID(c.Param("userId"))
	if parseError != nil {
		c.JSON(http.StatusBadRequest, parseError.Error())
		return
	}
	var ud updateUserData
	if validationErr := c.ShouldBindJSON(&ud); validationErr != nil {
		c.JSON(http.StatusBadRequest, lib.RequestError{ErrMsg: validationErr.Error()})
		return
	}
	updatedUser, updateErr := UpdateUser(db.Connection, userId, ud)
	if updateErr != nil {
		c.JSON(http.StatusBadRequest, lib.RequestError{ErrMsg: updateErr.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedUser)
}

func deleteUser(c *gin.Context) {
	userId, parseError := lib.ParseID(c.Param("userId"))
	if parseError != nil {
		c.JSON(http.StatusBadRequest, parseError.Error())
		return
	}
	deleteErr := DeleteUser(db.Connection, userId)
	if deleteErr != nil {
		c.JSON(http.StatusNotFound, lib.RequestError{ErrMsg: deleteErr.Error()})
		return
	}
	c.JSON(http.StatusOK, nil)
}
