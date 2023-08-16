package model

import (
	"time"
)

// User represents the structure of the user data
type User struct {
	ID               uint   `gorm:"primaryKey" json:"id"`
	UserName         string `gorm:"not null" json:"userName"`
	Password         string `gorm:"not null" json:"password"`
	Level            int    `gorm:"not null" json:"level"`
	RegistrationDate string `gorm:"not null" json:"registrationDate"`
	DeletionDate     string `json:"deletionDate"`
	LastUpdateDate   string `json:"lastUpdateDate"`
	AccountStatus    int    `json:"accountStatus"`
}
