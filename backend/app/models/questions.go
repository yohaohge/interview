package models

import (
	"time"
)

type Question struct {
	ID             int       `gorm:"primaryKey" json:"id"`
	Content        string    `json:"content"`
	Tags           string    `json:"tags"`
	Analysis       string    `json:"analysis"`
	RelatedCompany string    `json:"related_company"`
	CreatedAt      time.Time `json:"created_at"`
}

func (Question) TableName() string {
	return "questions"
}
