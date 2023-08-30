package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AlgorithmApi struct{}

// CreateAlgorithm 算法注册
// 参数 algorithm
// 返回值 被添加的算法 错误信息
func (s *AlgorithmApi) CreateAlgorithm(c *gin.Context) {
	var sysAlgorithm system.SysAlgorithm
	err := c.ShouldBindJSON(&sysAlgorithm)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	_, err = algorithmService.CreateAlgorithm(sysAlgorithm)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}
