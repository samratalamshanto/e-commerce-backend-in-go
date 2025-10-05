package utils

import (
	"ecom-backend/internal/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JSONResponse(c *gin.Context, status int, success bool,
	data interface{}, err error, msg string) {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}
	resp := common.Response{
		Code:    status,
		Success: success,
		Message: msg,
		Data:    data,
		Error:   errMsg,
	}
	c.JSON(status, resp)
}

func CreatedResponse(c *gin.Context, data interface{}, msg string) {
	JSONResponse(c, http.StatusCreated, true, data, nil, msg)
}

func SuccessResponse(c *gin.Context, data interface{}, msg string) {
	JSONResponse(c, http.StatusOK, true, data, nil, msg)
}

func BadRequestResponseWithError(c *gin.Context, err error, msg string) {
	JSONResponse(c, http.StatusBadRequest, false, nil, err, msg)
}

func BadRequestResponse(c *gin.Context, msg string) {
	JSONResponse(c, http.StatusBadRequest, false, nil, nil, msg)
}

func UnauthorizedResponse(c *gin.Context, err error, msg string) {
	JSONResponse(c, http.StatusUnauthorized, false, nil, err, msg)
}

func InternalServerErrorResponse(c *gin.Context, err error, msg string) {
	JSONResponse(c, http.StatusInternalServerError, false, nil, err, msg)
}

func PaginatedResponse(c *gin.Context, data interface{}, page, limit, total int,
	msg string) {
	totalPages := (total + limit - 1) / limit // ceiling division
	pagination := &common.Pagination{
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: totalPages,
	}
	resp := common.Response{
		Code:       http.StatusOK,
		Success:    true,
		Message:    msg,
		Data:       data,
		Error:      "",
		Pagination: pagination,
	}
	c.JSON(http.StatusOK, resp)
}
