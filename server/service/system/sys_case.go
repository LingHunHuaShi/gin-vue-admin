package system

/*
type SysCase struct {
	global.GVA_MODEL
	CaseID     uint      `json:"caseID" gorm:"default:0;primaryKey;autoIncrement;index;comment:案例ID"`
	UUID       uuid.UUID `json:"uuid" gorm:"comment:案例发起者UUID"`
	Title      string    `json:"title" gorm:"comment:案例标题"`
	Content    string    `json:"content" gorm:"comment:案例内容"`
	SubmitDate time.Time `json:"submitDate" gorm:"comment:提交时间"`
	Severity   int       `json:"severity" gorm:"default:0;comment:紧急等级"` // 0-3 递增等级
	Status     int       `json:"status" gorm:"default:0;comment:案例状态"`   // 0未处理 1正在处理 2已完成 3异常
	DateClosed time.Time `json:"dateClosed" gorm:"comment:完结日期"`
	Solution   string    `json:"solution" gorm:"comment:解决方案"`
}
*/

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

type CaseService struct{}

// CreateCase 注册案例创建函数
// 参数 caseObj 案例对象
// 返回值 被添加的案例对象 错误信息
func (c *CaseService) CreateCase(caseObj system.SysCase) (caseInter system.SysCase, err error) {
	err = global.GVA_DB.Create(&caseObj).Error
	if err != nil {
		return caseObj, errors.New("创建案例失败")
	}
	return caseObj, err
}

// FindCaseByUUID 根据发起者UUID查找案例
// 参数 uuid 发起者UUID
// 返回值 查找到的案例对象指针 错误信息
func (c *CaseService) FindCaseByUser(user system.SysUser) (caseInter *system.SysCase, err error) {
	var caseObj system.SysCase
	err = global.GVA_DB.Where("UUID = ?", user.UUID).First(&caseObj).Error
	if err != nil {
		return &caseObj, errors.New("案例不存在")
	}
	return &caseObj, err
}

// FindCaseByCaseID 根据案例ID查找案例
// 参数 caseID 案例ID
// 返回值 查找到的案例对象指针 错误信息
func (c *CaseService) FindCaseByCaseID(caseID uint) (caseInter *system.SysCase, err error) {
	var caseObj system.SysCase
	err = global.GVA_DB.Where("ID = ?", caseID).First(&caseObj).Error
	if err != nil {
		return &caseObj, errors.New("案例不存在")
	}
	return &caseObj, err
}

// DeleteCase 根据案例ID删除案例
// 参数 caseID 案例ID
// 返回值 错误信息
func (c *CaseService) DeleteCase(caseID uint) (err error) {
	var caseObj system.SysCase
	err = global.GVA_DB.Where("ID = ?", caseID).Delete(&caseObj).Error
	if err != nil {
		return errors.New("案例不存在")
	}
	return err
}

// UpdateCaseByCaseID 根据案例ID更新案例信息
// 参数 caseID 案例ID, newCase 新的案例信息
// 返回值 更新后的案例对象指针 错误信息
func (c *CaseService) UpdateCaseByCaseID(Case system.SysCase) (updateCase *system.SysCase, err error) {
	var caseObj system.SysCase

	// 查询案例是否存在
	err = global.GVA_DB.Where("ID = ?", Case.ID).First(&caseObj).Error
	if err != nil {
		return &caseObj, errors.New("案例不存在")
	}

	// 更新案例信息
	err = global.GVA_DB.Model(&caseObj).Updates(Case).Error
	if err != nil {
		return &caseObj, errors.New("更新失败")
	}
	return &caseObj, err
}

//	QueryAllCases 查询全部案例
//
// 参数 None
// 返回值 case列表
func (c *CaseService) QueryAllCases() (Container []system.SysCase, err error) {
	err = global.GVA_DB.Find(&Container).Error
	if err != nil {
		return nil, errors.New("查询失败")
	}
	return Container, err
}
