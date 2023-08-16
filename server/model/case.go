package model

import (
	"time"
)

type Case struct {
	CaseId         uint   `gorm:"primaryKey" json:"caseId"`
	UserId         uint   `gorm:"not null" json:"userId"`
	Title          string `gorm:"not null" json:"title"`
	Description    string `gorm:"not null" json:"description"`
	Status         string `gorm:"size:255" json:"status"`
	SubmissionDate string `gorm:"not null" json:"submissionDate"`
	CompletionDate string `json:"completionDate"`
	Solution       string `json:"solution"`
	Severity       string `gorm:"size:255" json:"severity"`
	User           User   `gorm:"foreignKey:UserId"`
}
