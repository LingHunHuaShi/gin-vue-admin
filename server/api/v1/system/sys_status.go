package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/gin-gonic/gin"
)

type SysStatusApi struct{}

func (Api *SysStatusApi) UploadSystemStatus(c *gin.Context) {
	var status system.SysStatus
	err := c.ShouldBindJSON(&status)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = sysStatusService.UploadSystemStatus(status, "")
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("系统数据成功上传", c)
}
