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
	table string
	data  []gorm.ColumnType
}

func NewRepoBuild(table string, data []gorm.ColumnType) *RepoBuild {
	return &RepoBuild{data: data}
}

func (r *RepoBuild) Exec() error {
	return nil
}

func (r *RepoBuild) buildRepo() {

}

func (r *RepoBuild) buildRepoIf() {

}

func (r *RepoBuild) buildRepoPo() {

}
