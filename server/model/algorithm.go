package model

import (
	"time"
)

type Algorithm struct {
	AlgorithmId      uint    `gorm:"primaryKey" json:"algorithmId"`
	AlgorithmName    string  `gorm:"not null" json:"algorithmName"`
	AlgorithmVersion string  `gorm:"not null" json:"algorithmVersion"`
	Description      string  `json:"description"`
	UpdationDate     string  `gorm:"not null" json:"updationDate"`
	Size             float64 `json:"size"`
	DownloadLink     string  `gorm:"not null" json:"downloadLink"`
	MD5              string  `gorm:"not null" json:"md5"`
}
