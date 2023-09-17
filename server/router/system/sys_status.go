package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type StatusRouter struct{}

func (r *StatusRouter) InitSysStatusRouter(Router *gin.RouterGroup) {
	statusRouter := Router.Group("status").Use(middleware.OperationRecord())
	statusApi := v1.ApiGroupApp.SystemApiGroup.SysStatusApi
	{
		statusRouter.PUT("uploadSystemStatus", statusApi.UploadSystemStatus)
	}
}
