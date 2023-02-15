package db

import (
	"fmt"

	"github.com/ukasyah99/kubik-cli/db/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error

	fmt.Println("init db")

	DB, err = gorm.Open(sqlite.Open("temp/my.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(
		&model.Credential{},
		&model.Cluster{},
	)
}
