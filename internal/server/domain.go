package server

import (
	_ "embed"
	"gorm.io/gorm"
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
}

func NewDomainBuild(table string, data []gorm.ColumnType) *DomainBuild {
	return &DomainBuild{
		table: table,
		data:  data,
	}
}

func (d *DomainBuild) Exec() error {
	return nil
}

func (d *DomainBuild) buildDo() {

}

func (d *DomainBuild) buildBiz() {

}
