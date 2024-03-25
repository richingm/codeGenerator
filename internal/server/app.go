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
	table string
	data  []gorm.ColumnType
}

func NewAppBuild(table string, data []gorm.ColumnType) *AppBuild {
	return &AppBuild{
		table: table,
		data:  data,
	}
}

func (a *AppBuild) Exec() error {
	return nil
}

func (a *AppBuild) buildApp() {

}
