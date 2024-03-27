package server

import (
	_ "embed"
	"gorm.io/gorm"
	"os"
	"richingm/codeGenerator/internal/conf"
	"text/template"
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
	table  string
	data   []gorm.ColumnType
	config conf.Config
}

func NewRepoBuild(table string, data []gorm.ColumnType, config conf.Config) *RepoBuild {
	return &RepoBuild{table: table, data: data, config: config}
}

func (r *RepoBuild) Exec() error {
	data := r.buildRepoData()
	// repoIf
	err := r.buildRepoIf(data)
	if err != nil {
		return err
	}

	// repoPo
	err = r.buildRepo(data)
	if err != nil {
		return err
	}

	// repo
	err = r.buildRepoPo(data)
	if err != nil {
		return err
	}
	return nil
}

type repoBuildData struct {
	RepoName        string
	RepoNameToLower string
	RepoIfName      string
	RepoPoName      string
	RepoDoStruct    string
	TableName       string
}

func (r *RepoBuild) buildRepoData() *repoBuildData {
	res := &repoBuildData{}
	res.RepoName = underscoreToCamelCase(CombineWords(r.table, "Repo"))
	res.RepoNameToLower = lowerCamelCase(underscoreToCamelCase(CombineWords(r.table, "Repo")))
	res.RepoIfName = underscoreToCamelCase(CombineWords(r.table, "RepoIf"))
	res.RepoPoName = underscoreToCamelCase(CombineWords(r.table, "Po"))
	res.RepoDoStruct = generateStructFromTable(res.RepoPoName, r.data)
	res.TableName = r.table
	return res
}

func (r *RepoBuild) buildRepo(data *repoBuildData) error {
	filePath := CombineWords(r.config.WorkDir, r.config.Repo.RelativePath)
	fileName := CombineWords(r.table, "_repo.go")
	return r.buildBase(data, filePath, fileName, repoTpl)
}

func (r *RepoBuild) buildRepoIf(data *repoBuildData) error {
	filePath := CombineWords(r.config.WorkDir, r.config.Domain.RelativePath)
	fileName := CombineWords(r.table, "_if.go")
	return r.buildBase(data, filePath, fileName, domainRepoIfTpl)
}

func (r *RepoBuild) buildRepoPo(data *repoBuildData) error {
	filePath := CombineWords(r.config.WorkDir, r.config.Domain.RelativePath)
	fileName := CombineWords(r.table, "_po.go")
	return r.buildBase(data, filePath, fileName, domainPoTpl)
}

func (r *RepoBuild) buildBase(data *repoBuildData, filePath, fileName, tpl string) error {
	t := template.Must(template.New("t").Parse(tpl))

	err := CreateFile(filePath, fileName, "")
	if err != nil {
		return err
	}

	fileFullPath := CombineWords(filePath, CombineWords("/", fileName))
	file, err := os.Create(fileFullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	err = t.Execute(file, data)
	if err != nil {
		return err
	}

	return nil
}
