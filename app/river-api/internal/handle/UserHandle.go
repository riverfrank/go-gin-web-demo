package handle

import (
	"github.com/gin-gonic/gin"
	"go-gin-web-demo/app/river-api/dto"
	"go-gin-web-demo/app/river-api/internal/service"
	"net/http"
)

func AddUser(c *gin.Context) {
	var user dto.UserCreateDTO
	c.BindJSON(&user)
	service.UserService.Create(&user)
	c.JSON(http.StatusOK, user)

}
