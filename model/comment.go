package model

import "time"

type Comment struct {
	ID          int       `gorm:"primary_key;column:id"`
	Description string    `gorm:"column:description"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (c Comment) TableName() string {
	return "comments"
}
