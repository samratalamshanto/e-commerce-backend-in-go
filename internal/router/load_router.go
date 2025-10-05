package router

import (
	"ecom-backend/internal/user"

	"github.com/gin-gonic/gin"
)

func LoadRouter(r *gin.Engine) {
	user.LoadUserRouter(r)

}
