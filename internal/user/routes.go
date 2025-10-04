package user

import "github.com/gin-gonic/gin"

type UserRouter struct {
	handler *UserHandler
}

func GetUserRouter(hanler *UserHandler) *UserRouter {
	return &UserRouter{handler: hanler}
}

func (r *UserRouter) RegisterRouter(router *gin.Engine) {
	userRouter := router.Group("/api/v1/users")
	{
		userRouter.POST("/", r.handler.CreateUser)
	}

}
