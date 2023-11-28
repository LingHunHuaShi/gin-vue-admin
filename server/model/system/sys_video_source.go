package system

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type SysVideoSource struct {
	global.GVA_MODEL
	VideoSource string `json:"videoSource" gorm:"comment:视频源;not null"`
}

func (SysVideoSource) TableName() string {
	return "sys_video_source"
}
