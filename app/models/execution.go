package models

import (
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
)

type Execution struct {
	gorm.Model
	Application string `gorm:"size:255;not null" json:"application"`
	Version     string `gorm:"size:255;not null" json:"version"`
	TypeExec    string `gorm:"size:255;not null" json:"typexec"`
	Date        string `gorm:"size:255;not null" json:"date"`
	Status      bool   `gorm:"not null" json:"status"`
	User        string `gorm:"size:255;not null" json:"user"`
}

func AddExecution(e Execution) error {

	db, _ := OpenSQL()

	if err := db.Create(&e).Error; err != nil {
		revel.AppLog.Error(err.Error())
		CloseSQL(db)
		return err

	}

	CloseSQL(db)
	return nil
}

func ListExecution() []Execution {

	var result []Execution

	db, _ := OpenSQL()
	db.Find(&result)
	db.Close()

	return result
}
