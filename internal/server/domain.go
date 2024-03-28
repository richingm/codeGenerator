package server

import (
	_ "embed"
	"gorm.io/gorm"
	"os"
	"richingm/codeGenerator/internal/conf"
	"text/template"
)

var (
	//go:embed tpl/domain.tpl
	domainTpl string

	//go:embed tpl/domain_do.tpl
	domainDoTpl string
)

type DomainBuild struct {
	table  string
	data   []gorm.ColumnType
	config conf.Config
}

func NewDomainBuild(table string, data []gorm.ColumnType, config conf.Config) *DomainBuild {
	return &DomainBuild{
		table:  table,
		data:   data,
		config: config,
	}
}

func (d *DomainBuild) Exec() error {
	return nil
}

type domainBuildData struct {
	RepoName        string
	RepoNameToLower string
	RepoIfName      string
	RepoPoName      string
	RepoDoStruct    string
	TableName       string
}

func (r *RepoBuild) buildDomainData() *domainBuildData {
	res := &domainBuildData{}
	//res.RepoName = underscoreToCamelCase(CombineWords(r.table, "Repo"))
	//res.RepoNameToLower = lowerCamelCase(underscoreToCamelCase(CombineWords(r.table, "Repo")))
	//res.RepoIfName = underscoreToCamelCase(CombineWords(r.table, "RepoIf"))
	//res.RepoPoName = underscoreToCamelCase(CombineWords(r.table, "Po"))
	//res.RepoDoStruct = generateStructFromTable(res.RepoPoName, r.data)
	res.TableName = r.table
	return res
}

func (d *DomainBuild) buildDo(data *domainBuildData) error {
	filePath := CombineWords(d.config.WorkDir, d.config.Domain.RelativePath)
	fileName := CombineWords(d.table, "_do.go")
	return d.buildBase(data, filePath, fileName, domainDoTpl)
}

func (d *DomainBuild) buildBiz(data *domainBuildData) error {
	filePath := CombineWords(d.config.WorkDir, d.config.Domain.RelativePath)
	fileName := CombineWords(d.table, ".go")
	return d.buildBase(data, filePath, fileName, domainTpl)
}

func (d *DomainBuild) buildBase(data *domainBuildData, filePath, fileName, tpl string) error {
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
