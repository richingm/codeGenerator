package biz

import "time"

type KnowledgePo struct {
	Id          int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Pid         int64
	Name        string
	ImportLevel string
	Notes       string
}

func (KnowledgePo) TableName() string {
	return "knowledges"
}

func (k *KnowledgePo) ConvertToDo(do *KnowledgeDo) {
	do.Id = k.Id
	do.CreatedAt = k.CreatedAt
	do.UpdatedAt = k.UpdatedAt
	do.Pid = k.Pid
	do.Name = k.Name
	do.ImportLevel = k.ImportLevel
	do.Notes = k.Notes
}
