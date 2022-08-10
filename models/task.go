package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Task struct {
	Name        string    `gorm:"size:191;not null" json:"name"`
	Assignee    User      `gorm:"not null;references:AssigneeID;" json:"assignee"`
	AssigneeID  int64     `gorm:"not null;" json:"assignee_id"`
	Deadline    time.Time `gorm:"not null;" json:"deadline"`
	IsDone      bool      `gorm:"not null;" json:"is_done"`
	CreatedByID int64     `gorm:"not null;" json:"created_by"`
	gorm.Model
}

func (t *Task) Create() (*Task, error) {
	var err error
	err = DB.Create(&t).Error

	if err != nil {
		return &Task{}, err
	}

	return t, nil
}

func (t *Task) Update() (*Task, error) {
	var err error
	err = DB.Update(&t).Error

	if err != nil {
		return &Task{}, err
	}

	return t, nil
}

func (t *Task) Delete() (*Task, error) {
	var err error
	err = DB.Delete(&t).Error

	if err != nil {
		return &Task{}, err
	}

	return t, nil
}
