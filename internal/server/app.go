package server

import (
	_ "embed"
	"gorm.io/gorm"
	"os"
	"richingm/codeGenerator/internal/conf"
	"text/template"
)

var (
	//go:embed tpl/app.tpl
	appTpl string
)

type AppBuild struct {
	table  string
	data   []gorm.ColumnType
	config conf.Config
}

func NewAppBuild(table string, data []gorm.ColumnType, config conf.Config) *AppBuild {
	return &AppBuild{
		table:  table,
		data:   data,
		config: config,
	}
}

func (a *AppBuild) Exec() error {
	data := a.buildAppData()

	// app
	err := a.buildApp(data)
	if err != nil {
		return err
	}

	return nil
}

type appBuildData struct {
	RepoName        string
	RepoNameToLower string
	RepoIfName      string
	RepoPoName      string
	RepoDoStruct    string
	TableName       string
}

func (a *AppBuild) buildAppData() *appBuildData {
	res := &appBuildData{}
	//res.RepoName = underscoreToCamelCase(CombineWords(r.table, "Repo"))
	//res.RepoNameToLower = lowerCamelCase(underscoreToCamelCase(CombineWords(r.table, "Repo")))
	//res.RepoIfName = underscoreToCamelCase(CombineWords(r.table, "RepoIf"))
	//res.RepoPoName = underscoreToCamelCase(CombineWords(r.table, "Po"))
	//res.RepoDoStruct = generateStructFromTable(res.RepoPoName, r.data)
	res.TableName = a.table
	return res
}

func (a *AppBuild) buildApp(data *appBuildData) error {
	filePath := CombineWords(a.config.WorkDir, a.config.App.RelativePath)
	fileName := CombineWords(a.table, ".go")
	return a.buildBase(data, filePath, fileName, appTpl)
}

func (a *AppBuild) buildBase(data *appBuildData, filePath, fileName, tpl string) error {
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
