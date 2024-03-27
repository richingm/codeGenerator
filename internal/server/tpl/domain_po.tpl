package biz

import "time"

{{.RepoDoStruct}}

func ({{.RepoPoName}}) TableName() string {
	return "{{.TableName}}"
}

func (k *{{.RepoPoName}}) ConvertToDo(do *{{.RepoPoName}}) {
	do.Id = k.Id
	do.CreatedAt = k.CreatedAt
	do.UpdatedAt = k.UpdatedAt
	do.Pid = k.Pid
	do.Name = k.Name
	do.ImportLevel = k.ImportLevel
	do.Notes = k.Notes
}
