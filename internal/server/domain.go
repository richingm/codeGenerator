package server

import "gorm.io/gorm"

type DomainBuild struct {
	data []gorm.ColumnType
}

func NewDomainBuild(data []gorm.ColumnType) *DomainBuild {
	return &DomainBuild{
		data: data,
	}
}
