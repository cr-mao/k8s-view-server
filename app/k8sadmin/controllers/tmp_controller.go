package controllers

import (
	"fmt"
	"github.com/cr-mao/k8s-view-server/app/k8sadmin/global"
	"github.com/gin-gonic/gin"
	"github.com/google/martian/log"

	"github.com/cr-mao/k8s-view-server/infra/errcode"
	"github.com/cr-mao/k8s-view-server/infra/response"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type TmpController struct {
}

// pod 测试
func (cc *TmpController) PodList(c *gin.Context) {

	list, err := global.KubeConfigSet.CoreV1().Pods("").List(c, metav1.ListOptions{})
	if err != nil {
		log.Errorf("err :%v", err)
		response.ErrorAbort(c, errcode.ErrCodes.ErrInternalServer)
	}
	for _, value := range list.Items {
		fmt.Println(value.Namespace, value.Name)
	}
	response.Success(c, errcode.ErrCodes.ErrNo, map[string]interface{}{
		"token": "a111",
	})
	return
}
