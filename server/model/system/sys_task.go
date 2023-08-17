package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/gofrs/uuid"
	"time"
)

type SysTask struct {
	global.GVA_MODEL
	TaskID       uint      `json:"taskID" gorm:"default:0;primaryKey;autoIncrement;index;comment:任务ID"` // 任务ID
	UUID         uuid.UUID `json:"uuid" gorm:"comment:用户UUID"`                                          // 用户UUID
	VideoSource  string    `json:"videoSource" gorm:"comment:视频源"`                                      // 视频源
	CreationDate time.Time `json:"creationDate" gorm:"comment:创建时间"`
	Resolution   int       `json:"resolution" gorm:"comment:分辨率等级"`
	AlgorithmID  uint      `json:"algorithmId" gorm:"comment:任务容器算法ID"`
	Intensity    int       `json:"intensity" gorm:"comment:任务粒度"`
	Status       int       `json:"status" gorm:"comment:任务状态"`
}

func (SysTask) TableName() string {
	return "sys_tasks"
}
