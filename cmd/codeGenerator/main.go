package main

import (
	"errors"
	"fmt"
	"github.com/samber/lo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"richingm/codeGenerator/internal/conf"
	"richingm/codeGenerator/internal/server"
)

func main() {
	// 获取数据表信息
	tablesInfo, err := getTablesInfo(nil)
	if err != nil {
		panic(err)
	}

	config := conf.GetConfig()

	for table, columns := range tablesInfo {
		// repo
		repoBuild := server.NewRepoBuild(table, columns, config)
		err := repoBuild.Exec()
		if err != nil {
			panic(err)
		}

		// domain
		domainBuild := server.NewDomainBuild(table, columns, config)
		err = domainBuild.Exec()
		if err != nil {
			panic(err)
		}

		// app
		appBuild := server.NewAppBuild(table, columns, config)
		err = appBuild.Exec()
		if err != nil {
			panic(err)
		}
	}
}

// ColumnInfo 使用Raw SQL获取表的列信息
type ColumnInfo struct {
	CID        string  `gorm:"column:cid"`
	Name       string  `gorm:"column:name"`
	Type       string  `gorm:"column:type"`
	Notnull    int     `gorm:"column:notnull"`
	Dflt_value *string `gorm:"column:dflt_value"`
	Pk         int     `gorm:"column:pk"`
}

func getTablesInfo(tables []string) (map[string][]gorm.ColumnType, error) {
	// 连接数据库
	dsn := "host=localhost user=blog password=blog dbname=blogs port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 获取数据表
	all, err := db.Migrator().GetTables()
	if err != nil {
		return nil, err
	}
	if len(tables) == 0 {
		tables = all
	} else {
		for _, table := range tables {
			if !lo.Contains(all, table) {
				return nil, errors.New(fmt.Sprintf("table:%s,不存在", table))
			}
		}
	}

	// 获取数据表信息
	res := make(map[string][]gorm.ColumnType)
	migrator := db.Migrator()
	for _, table := range tables {
		columns, err := migrator.ColumnTypes(table)
		if err != nil {
			return nil, errors.New("获取字段信息失败: " + err.Error())
		}
		res[table] = columns
	}
	return res, nil
}
