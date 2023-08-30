package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	system2 "github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/gin-gonic/gin"
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
	newAlgorithm, err := system2.AlgorithmService.CreateAlgorithm(sysAlgorithm)
}
