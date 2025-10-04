package user

import (
	"ecom-backend/internal/common"
	"ecom-backend/pkg/utils"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *UserService
}

func GetUserHandler(service *UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (r *UserHandler) CreateUser(c *gin.Context) {
	var dto CreateUserDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		utils.BadRequestResponse(c, err, utils.MsgInvalidInput)
		return
	}
	saved, err := r.service.CreateUser(dto)
	if err != nil {
		utils.InternalServerErrorResponse(c, err, utils.MsgDatabaseError)
		return
	}
	userResponse, mapErr := common.Mapper[User, UserResponse](*saved)
	if mapErr != nil {
		utils.InternalServerErrorResponse(c, mapErr, utils.MsgMappingError)
		return
	}
	utils.CreatedResponse(c, userResponse, utils.MsgCreateSuccess)
}
