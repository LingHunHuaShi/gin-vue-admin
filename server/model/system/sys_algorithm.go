package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

type SysAlgorithm struct {
	global.GVA_MODEL
	AlgorithmName    string    `json:"algorithmName" gorm:"comment:算法名"`
	AlgorithmVersion string    `json:"algorithmVersion" gorm:"comment:算法版本"`
	Description      string    `json:"description" gorm:"comment:算法描述"`
	UpdateDate       time.Time `json:"updateDate" gorm:"comment:更新时间"`
	Size             float32   `json:"size" gorm:"comment:模型文件大小"`
	MD5              string    `json:"MD5" gorm:"comment:MD5校验值"`
}

type SysSyncAlgorithm struct {
	AlgorithmName    string    `json:"algorithmName"`
	AlgorithmVersion string    `json:"algorithmVersion"`
	Description      string    `json:"algorithmDescription"`
	UpdateDate       time.Time `json:"createTime"`
	Size             float32   `json:"algorithmSize"`
	MD5              string    `json:"algorithmMd5"`
	DownLoadLink     string    `json:"algorithmDownloadlink"`
}

func (SysAlgorithm) TableName() string {
	return "sys_algorithms"
}
