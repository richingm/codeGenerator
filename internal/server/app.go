package server

import (
	_ "embed"
	"gorm.io/gorm"
)

var (
	//go:embed tpl/app.tpl
	appTpl string
)

type AppBuild struct {
	data []gorm.ColumnType
}

func NewAppBuild(data []gorm.ColumnType) *DomainBuild {
	return &DomainBuild{
		data: data,
	}
}
