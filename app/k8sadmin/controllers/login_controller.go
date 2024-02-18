package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/cr-mao/k8s-view-server/infra/errcode"
	"github.com/cr-mao/k8s-view-server/infra/response"
)

type AdminLoginController struct {
}

//后台登录
func (cc *AdminLoginController) Login(c *gin.Context) {
	response.Success(c, errcode.ErrCodes.ErrNo, map[string]interface{}{
		"token": "a111",
	})
	return
}
