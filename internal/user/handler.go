package user

import (
	"ecom-backend/internal/common"
	"ecom-backend/pkg/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	service *UserService
}

func GetUserHandler(service *UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	idString := c.Param("id")
	if idString == "" {
		utils.BadRequestResponse(c, utils.MsgMissingMandatoryParams)
		return
	}
	id, err := utils.ParseUintFromString(idString)
	if err != nil {
		utils.InternalServerErrorResponse(c, err, utils.MsgParsingError)
		return
	}
	user, dbErr := h.service.GetUserByID(id)
	if dbErr != nil {
		if dbErr == gorm.ErrRecordNotFound {
			utils.BadRequestResponse(c, utils.MsgRecordNotFound)
		} else {
			utils.InternalServerErrorResponse(c, dbErr, utils.MsgDatabaseError)
		}
		return
	}
	userResponse, mapErr := common.Mapper[User, UserResponse](*user)
	if mapErr != nil {
		utils.InternalServerErrorResponse(c, mapErr, utils.MsgMappingError)
		return
	}
	utils.SuccessResponse(c, userResponse, utils.MsgSuccess)
}

func (h *UserHandler) GetAllUser(c *gin.Context) {
	userList, err := h.service.GetAllUsers()
	if err != nil {
		utils.InternalServerErrorResponse(c, err, utils.MsgDatabaseError)
		return
	}
	userResponseList, mapperErr := common.MapperForList[User, UserResponse](userList)
	if mapperErr != nil {
		utils.InternalServerErrorResponse(c, mapperErr, utils.MsgMappingError)
		return
	}
	utils.CreatedResponse(c, userResponseList, utils.MsgCreateSuccess)
}

func (h *UserHandler) GetAllUserPagination(c *gin.Context) {
	offsetStr := c.Param("offset")
	if offsetStr == "" {
		utils.BadRequestResponse(c, utils.MsgMissingMandatoryParams)
		return
	}
	offset, err := utils.ParseIntFromString(offsetStr)
	if err != nil {
		utils.InternalServerErrorResponse(c, err, utils.MsgParsingError)
		return
	}

	limitStr := c.Param("limit")
	if offsetStr == "" {
		utils.BadRequestResponse(c, utils.MsgMissingMandatoryParams)
		return
	}
	limit, err := utils.ParseIntFromString(limitStr)
	if err != nil {
		utils.InternalServerErrorResponse(c, err, utils.MsgParsingError)
		return
	}

	userList, err := h.service.GetAllUsersPagination(offset, limit)
	if err != nil {
		utils.InternalServerErrorResponse(c, err, utils.MsgDatabaseError)
		return
	}

	userResponseList, mapperErr := common.MapperForList[User, UserResponse](userList)
	if mapperErr != nil {
		utils.InternalServerErrorResponse(c, mapperErr, utils.MsgMappingError)
		return
	}
	utils.CreatedResponse(c, userResponseList, utils.MsgCreateSuccess)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var dto CreateUserDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		utils.BadRequestResponseWithError(c, err, utils.MsgInvalidInput)
		return
	}
	saved, err := h.service.CreateUser(dto)
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

func (h *UserHandler) UpdateUser(c *gin.Context) {}

func (h *UserHandler) DeleteUserByID(c *gin.Context) {
	idString := c.Param("id")
	if idString == "" {
		utils.BadRequestResponse(c, utils.MsgMissingMandatoryParams)
		return
	}
	id, err := utils.ParseUintFromString(idString)
	if err != nil {
		utils.InternalServerErrorResponse(c, err, utils.MsgParsingError)
		return
	}
	success, dbErr := h.service.DeleteUserByID(id)
	if dbErr != nil {
		if dbErr == gorm.ErrRecordNotFound {
			utils.BadRequestResponse(c, utils.MsgRecordNotFound)
		} else {
			utils.InternalServerErrorResponse(c, dbErr, utils.MsgDatabaseError)
		}
		return
	}
	utils.SuccessResponse(c, success, utils.MsgDeleteSuccess)
}
