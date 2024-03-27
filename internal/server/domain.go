package server

import (
	_ "embed"
	"gorm.io/gorm"
	"richingm/codeGenerator/internal/conf"
)

var (
	//go:embed tpl/domain.tpl
	domainTpl string

	//go:embed tpl/domain_do.tpl
	domainDoTpl string
)

type DomainBuild struct {
	table string
	data  []gorm.ColumnType
	conf  conf.Config
}

func NewDomainBuild(table string, data []gorm.ColumnType, config conf.Config) *DomainBuild {
	return &DomainBuild{
		table: table,
		data:  data,
		conf:  config,
	}
}

func (d *DomainBuild) Exec() error {
	return nil
}

func (d *DomainBuild) buildDo() {

}

func (d *DomainBuild) buildBiz() {

}
