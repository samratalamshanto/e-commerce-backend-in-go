package user

import (
	"ecom-backend/internal/common"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	handler *UserHandler
}

func GetUserRouter(hanler *UserHandler) *UserRouter {
	return &UserRouter{handler: hanler}
}

func (r *UserRouter) RegisterRouter(router *gin.Engine) {
	userRouter := router.Group("/api/v1/users")
	{
		userRouter.GET("/:id", r.handler.GetUserByID)
		userRouter.GET("/", r.handler.GetAllUser)
		userRouter.GET("/pagination/:offset/:limit", r.handler.GetAllUserPagination)

		userRouter.POST("/", r.handler.CreateUser)
		userRouter.PUT("/:id", r.handler.UpdateUser)

		userRouter.DELETE("/:id", r.handler.DeleteUserByID)
	}

}

func LoadUserRouter(r *gin.Engine) {
	repo := GetUserRepository(common.DBInstance)
	service := GetUserService(repo)
	handler := GetUserHandler(service)
	router := GetUserRouter(handler)

	router.RegisterRouter(r)
}
