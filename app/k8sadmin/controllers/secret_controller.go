/*
*
User: cr-mao
Date: 2024/2/25 09:38
Email: crmao@qq.com
Desc: secret_controller.go
*/
package controllers

import (
	"github.com/cr-mao/k8s-view-server/app/k8sadmin/request/secret_request"
	"github.com/cr-mao/k8s-view-server/app/k8sadmin/services"
	"github.com/cr-mao/k8s-view-server/infra/errcode"
	"github.com/cr-mao/k8s-view-server/infra/response"
	"github.com/gin-gonic/gin"
)

type SecretController struct{}

// 创建或更新secret
func (c *SecretController) CreateOrUpdateSecret(ctx *gin.Context) {
	var secretReq *secret_request.Secret
	if err := ctx.ShouldBind(&secretReq); err != nil {
		response.ErrorAbort(ctx, errcode.ErrCodes.ErrParams, err.Error())
		return
	}
	if err := services.SecretService.CreateOrUpdateSecret(ctx, secretReq); err != nil {
		response.ErrorAbort(ctx, errcode.ErrCodes.ErrInternalServer, err.Error())
		return
	}
	response.Success(ctx, errcode.ErrCodes.ErrNo, nil)
}

// 删除secret
func (c *SecretController) DeleteSecret(ctx *gin.Context) {
	namespace := ctx.Query("namespace")
	name := ctx.Query("name")
	if namespace == "" || name == "" {
		response.ErrorAbort(ctx, errcode.ErrCodes.ErrParams, "namespace ,name 不能为空")
		return
	}
	if err := services.SecretService.DeleteSecret(ctx, namespace, name); err != nil {
		response.ErrorAbort(ctx, errcode.ErrCodes.ErrInternalServer, err.Error())
		return
	}
	response.Success(ctx, errcode.ErrCodes.ErrNo, nil)
}

// secret  详情
func (c *SecretController) GetSecretDetail(ctx *gin.Context) {
	namespace := ctx.Query("namespace")
	name := ctx.Query("name")
	if namespace == "" || name == "" {
		response.ErrorAbort(ctx, errcode.ErrCodes.ErrParams, "namespace ,name 不能为空")
		return
	}
	res, err := services.SecretService.GetSecretDetail(ctx, namespace, name)
	if err != nil {
		response.ErrorAbort(ctx, errcode.ErrCodes.ErrInternalServer, err.Error())
		return
	}
	response.Success(ctx, errcode.ErrCodes.ErrNo, res)
}

// secret 列表
func (c *SecretController) GetSecretList(ctx *gin.Context) {
	namespace := ctx.Query("namespace")
	res, err := services.SecretService.GetSecretList(ctx, namespace)
	if err != nil {
		response.ErrorAbort(ctx, errcode.ErrCodes.ErrInternalServer, err.Error())
		return
	}
	response.Success(ctx, errcode.ErrCodes.ErrNo, res)
}
