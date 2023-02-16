package credential

import (
	"errors"
	"fmt"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/ukasyah99/kubik-cli/db"
	"github.com/ukasyah99/kubik-cli/db/model"
)

type Credential struct {
	model.Credential
}

func (c *Credential) Save() error {
	if c.ID == "" {
		// Create new credential

		// Generate id
		nanoid, err := gonanoid.New(21)
		if err != nil {
			return err
		}
		c.ID = fmt.Sprintf("cre-%s", nanoid)

		db.DB.Create(c)
	} else {
		// Update existing credential

		// Check if record exists
		exists, err := Exists(c.ID)
		if err != nil {
			return err
		}
		if !exists {
			return errors.New("credential not found")
		}

		db.DB.Save(c)
	}

	return nil
}

func Exists(id string) (bool, error) {
	var count int64
	err := db.DB.Model(&model.Credential{}).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// func Get(id string) (*model.Credential, error) {
// 	data := model.Credential{ID: id}

// 	if err := db.DB.First(&data).Error; err != nil {
// 		return nil, err
// 	}

// 	return &data, nil
// }

// func Describe(id string) (*model.Credential, error) {
// 	return nil, errors.New("not implemented")
// }

// func Update(id string) {
// 	return nil, errors.New("not implemented")
// }

// func Delete() {
// 	fmt.Println("TODO: to be implemented")
// }
