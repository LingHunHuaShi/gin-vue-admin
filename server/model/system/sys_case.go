package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/gofrs/uuid"
	"time"
)

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

func (SysCase) TableName() string {
	return "sys_cases"
}
