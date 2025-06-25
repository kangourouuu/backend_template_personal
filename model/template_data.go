package model

import "time"

type ProjectInfo struct {
	ProjectName string    `json:"project_name" gorm:"project_name"`
	Port        int       `json:"port" gorm:"port"`
	CreatedAt   time.Time `json:"created_at" gorm:"created_at"`
}

func (ProjectInfo) TableName() string {
	return "project_info"
}