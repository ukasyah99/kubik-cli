package model

import "time"

type Credential struct {
	ID        string `gorm:"primaryKey;type:varchar(25)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `gorm:"type:varchar(32)"`
	Provider  string `gorm:"type:varchar(15)"`
	Token     string `gorm:"type:varchar(200)"`
}

// type Cluster struct {
// 	ID     string
// 	Name   string
// 	Status string
// }
