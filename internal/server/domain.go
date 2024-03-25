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
	data []gorm.ColumnType
}

func NewDomainBuild(data []gorm.ColumnType) *DomainBuild {
	return &DomainBuild{
		data: data,
	}
}
