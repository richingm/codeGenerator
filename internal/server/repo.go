package server

import (
	_ "embed"
	"gorm.io/gorm"
)

var (
	//go:embed tpl/domain_po.tpl
	domainPoTpl string

	//go:embed tpl/domain_repo_if.tpl
	domainRepoIfTpl string

	//go:embed tpl/repo.tpl
	repoTpl string
)

type RepoBuild struct {
	data []gorm.ColumnType
}

func NewRepoBuild(data []gorm.ColumnType) *RepoBuild {
	return &RepoBuild{data: data}
}
