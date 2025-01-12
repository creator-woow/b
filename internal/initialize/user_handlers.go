package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"happy-server/internal/handlers"
	"happy-server/internal/services"
)

func UserHandlers(r *gin.Engine, s *services.UserService) {
	userUrl := fmt.Sprintf("/:%v", handlers.UserIdParam)
	g := r.Group("/user")
	g.GET("/", handlers.ReadUsersList(s))
	g.GET(userUrl, handlers.ReadUser(s))
	g.PATCH(userUrl, handlers.UpdateUser(s))
	g.DELETE(userUrl, handlers.DeleteUser(s))
}
