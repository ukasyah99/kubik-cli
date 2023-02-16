package model

import "gorm.io/gorm"

type Credential struct {
	gorm.Model
	Title    string
	Provider string
	Token    string
}

type Cluster struct {
	ID     string
	Name   string
	Status string
}
