package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

type SysAlgorithm struct {
	global.GVA_MODEL
	AlgorithmID      uint      `json:"algorithmID" gorm:"default:0;primary_key;auto_increment;index;comment:算法ID"`
	AlgorithmName    string    `json:"algorithmName" gorm:"comment:算法名"`
	AlgorithmVersion string    `json:"algorithmVersion" gorm:"comment:算法版本"`
	Description      string    `json:"description" gorm:"comment:算法描述"`
	UpdateDate       time.Time `json:"updateDate" gorm:"comment:更新时间"`
	Size             float32   `json:"size" gorm:"comment:模型文件大小"`
	//	DownloadLink     string    `json:"downloadLink" gorm:"comment:模型下载链接"`
	MD5 string `json:"MD5" gorm:"comment:MD5校验值"`
}

func (SysAlgorithm) TableName() string {
	return "sys_algorithms"
}
