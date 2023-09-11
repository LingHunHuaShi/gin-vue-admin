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

// DeleteAlgorithm 根据算法ID删除算法及其相关权限信息
// @param id uint 算法ID
// @return err error 删除操作的错误，如果删除失败则返回错误信息，否则为nil
func (s *AlgorithmApi) DeleteAlgorithm(c *gin.Context) {
	var algorithm system.SysAlgorithm
	err := c.ShouldBindJSON(&algorithm)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = algorithmService.DeleteAlgorithm(algorithm.AlgorithmID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败!"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功!", c)
}

// FindAlgorithmById FindAlgorithmByID 根据算法ID查找算法信息
// @param algorithmID uint 要查找的算法ID
// @return algorithmInter *system.SysAlgorithm 查找到的算法信息，如果找到则为非空指针，否则为nil
// @return err error 查找操作的错误，如果算法不存在或查找失败则返回错误信息，否则为nil
func (s *AlgorithmApi) FindAlgorithmById(c *gin.Context) {
	var algorithm system.SysAlgorithm
	err := c.ShouldBindJSON(&algorithm)
	if err != nil {
		response.FailWithMessage("查找失败!", c)
		return
	}
	algo, err := algorithmService.FindAlgorithmById(algorithm.AlgorithmID)
	if err != nil {
		global.GVA_LOG.Error("查找失败!", zap.Error(err))
		response.FailWithMessage("查找失败!"+err.Error(), c)
		return
	}
	response.OkWithDetailed(algo, "查找成功!", c)
}

// UpdateAlgorithm 更新算法信息
// @param algorithmID uint 要更新的算法ID
// @param newAlgorithm system.SysAlgorithm 新的算法信息
// @return updatedAlgorithmInter *system.SysAlgorithm 更新后的算法信息，如果更新成功则为非空指针，否则为nil
// @return err error 更新操作的错误，如果更新失败则返回错误信息，否则为nil
func (s *AlgorithmApi) UpdateAlgorithm(c *gin.Context) {
	var algorithm system.SysAlgorithm
	err := c.ShouldBindJSON(&algorithm)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	newAlgo, err := algorithmService.UpdateAlgorithm(algorithm.AlgorithmID, algorithm)
	if err != nil {
		global.GVA_LOG.Error("更新失败2 GVA_LOG!", zap.Error(err))
		response.FailWithMessage("更新失败2 RESPONSE!", c)
		return
	}
	response.OkWithDetailed(newAlgo, "更新成功!", c)
}
