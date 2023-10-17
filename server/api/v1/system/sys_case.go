package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CaseApi struct{}

// CreateCase 注册案例创建函数
// 参数 caseObj 案例对象
// 返回值 被添加的案例对象 错误信息
func (Api *CaseApi) CreateCase(c *gin.Context) {
	var Case system.SysCase
	err := c.ShouldBindJSON(&Case)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	_, err = caseService.CreateCase(Case)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// FindCaseByUUID 根据发起者UUID查找案例
// 参数 uuid 发起者UUID
// 返回值 查找到的案例对象指针 错误信息
func (Api *CaseApi) FindCaseByUUID(c *gin.Context) {
	var User system.SysUser
	err := c.ShouldBindJSON(&User)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	Case, err := caseService.FindCaseByUser(User)
	if err != nil {
		global.GVA_LOG.Error("查找失败!", zap.Error(err))
		response.FailWithMessage("查找失败!"+err.Error(), c)
		return
	}
	response.OkWithDetailed(Case, "查找成功!", c)
}

// FindCaseByCaseID 根据案例ID查找案例
// 参数 caseID 案例ID
// 返回值 查找到的案例对象指针 错误信息
func (Api *CaseApi) FindCaseByCaseID(c *gin.Context) {
	var Case system.SysCase
	err := c.ShouldBindJSON(&Case)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	MyCase, err := caseService.FindCaseByCaseID(Case.ID)
	if err != nil {
		global.GVA_LOG.Error("查找失败!", zap.Error(err))
		response.FailWithMessage("查找失败!"+err.Error(), c)
		return
	}
	response.OkWithDetailed(MyCase, "查找成功!", c)
}

// DeleteCase 根据案例ID删除案例
// 参数 caseID 案例ID
// 返回值 错误信息
func (Api *CaseApi) DeleteCase(c *gin.Context) {
	var Case system.SysCase
	err := c.ShouldBindJSON(&Case)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = caseService.DeleteCase(Case.ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败!"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功!", c)
}

// UpdateCase 根据案例ID更新案例信息
// 参数 caseID 案例ID, newCase 新的案例信息
// 返回值 更新后的案例对象指针 错误信息
func (Api *CaseApi) UpdateCase(c *gin.Context) {
	var Case system.SysCase
	err := c.ShouldBindJSON(&Case)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	MyCase, err := caseService.UpdateCaseByCaseID(Case)
	if err != nil {
		global.GVA_LOG.Error("更新失败2!", zap.Error(err))
		response.FailWithMessage("更新失败2!", c)
		return
	}
	response.OkWithDetailed(MyCase, "更新成功!", c)
}

//	QueryAllCases 查询全部案例
//
// 参数 None
// 返回值 case列表
func (Api *CaseApi) QueryAllCases(c *gin.Context) {
	var cases []system.SysCase
	cases, err := caseService.QueryAllCases()
	if err != nil {
		global.GVA_LOG.Error("查询失败", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithDetailed(cases, "查询成功", c)
}
