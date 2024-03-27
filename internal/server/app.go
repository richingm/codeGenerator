package server

import (
	_ "embed"
	"gorm.io/gorm"
	"richingm/codeGenerator/internal/conf"
)

var (
	//go:embed tpl/app.tpl
	appTpl string
)

type AppBuild struct {
	table string
	data  []gorm.ColumnType
	conf  conf.Config
}

func NewAppBuild(table string, data []gorm.ColumnType, config conf.Config) *AppBuild {
	return &AppBuild{
		table: table,
		data:  data,
		conf:  config,
	}
}

func (a *AppBuild) Exec() error {
	return nil
}

func (a *AppBuild) buildApp() {

}
