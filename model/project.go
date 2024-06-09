package model

import (
	"github.com/pocketbase/pocketbase/models"
)

var _ models.Model = (*Project)(nil)

const (
	PROJECT = "project"
)

type Project struct {
	Name   string `db:"name"`
	Slug   string `db:"slug"`
	Type   string `db:"type"`
	Parent string `db:"parent"`

	models.BaseModel
}

func (*Project) TableName() string {
	return PROJECT
}
