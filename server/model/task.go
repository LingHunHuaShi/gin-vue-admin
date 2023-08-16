package model

import (
	"time"
)

type Task struct {
	TaskId       uint      `gorm:"primaryKey" json:"taskId"`
	ID           uint      `gorm:"not null" json:"id"`
	CreationDate string    `gorm:"not null" json:"creationDate"`
	VideoSource  string    `gorm:"not null" json:"videoSource"`
	Resolution   string    `gorm:"not null" json:"resolution"`
	AlgorithmId  string    `gorm:"not null" json:"algorithmId"`
	Intensity    int       `gorm:"not null" json:"intensity"`
	Status       int       `gorm:"not null" json:"status"`
	Algorithm    Algorithm `gorm:"foreignKey:AlgorithmId"`
	User         User      `gorm:"foreignKey:ID"`
}
