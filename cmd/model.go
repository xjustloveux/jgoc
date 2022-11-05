// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"github.com/xjustloveux/jgo/jsql"
	"github.com/xjustloveux/jgoc/model"
	"regexp"
	"strings"
)

func checkModel() error {

	if !root.Model {

		return nil
	}
	jPrint("start load schema...")
	if err := loadSchema(); err != nil {

		return err
	}
	jPrint("load schema success")
	/*
		model
		|---dsName
			|---tableName.go
	*/
	path := "model"
	if err := checkExistAndCreateFolder(path); err != nil {

		return err
	}
	path = fmt.Sprint(path, "/")
	if err := createModelDS(path); err != nil {

		return err
	}
	return nil
}

func createModelDS(path string) error {

	for _, v := range datasource {

		dsName := fmt.Sprint(strToL(v.Name), "mod")
		dsPath := fmt.Sprint(path, dsName)
		if err := checkExistAndCreateFolder(dsPath); err != nil {

			return err
		}
		for _, t := range v.Tables {

			create := len(root.Table) <= 0 || t.Name == root.Table
			if create {

				if err := createModelTable(fmt.Sprint(dsPath, "/"), dsName, t); err != nil {

					return err
				}
			}
		}
	}
	return nil
}

func createModelTable(path, dsName string, table model.Table) error {

	fileName := fmt.Sprint(strToL(table.Name), ".go")
	filePath := fmt.Sprint(path, fileName)
	jPrint(fmt.Sprint("create ", filePath, "..."))
	if err := createFile(filePath, getModelTableContent(dsName, table)); err != nil {

		return err
	}
	return nil
}

func getModelTableContent(dsName string, table model.Table) string {

	module := ""
	if haveTimeCol(table) {

		module = `

import "time"`
	}
	tableName := strToUCC(table.Name)
	tableComment := ""
	if len(table.Schema) > 0 {

		tableComment = table.Schema[0].TableComment
	}
	col := ""
	for _, v := range table.Schema {

		col += getColumnContent(v)
	}
	return `package ` + dsName + module + `

// ` + tableName + ` ` + tableComment + `
type ` + tableName + ` struct {` + col + `
}`
}

func haveTimeCol(table model.Table) bool {

	for _, v := range table.Schema {

		t := strings.ToLower(v.DataType)
		switch t {
		case "date":
			fallthrough
		case "datetime":
			fallthrough
		case "datetime2":
			fallthrough
		case "smalldatetime":
			fallthrough
		case "time":
			fallthrough
		case "timestamp":
			return true
		default:
			if b, err := regexp.MatchString("^timestamp", t); err == nil && b {

				return true
			} else if err != nil {

				jPrint(err)
			}
		}
	}
	return false
}

func getColumnContent(col jsql.TableSchema) string {

	return `
` + getColComment(col) + `
` + getTab() + strToUCC(col.ColumnName) + ` ` + getColType(col) + ` ` + getColJson(col)
}

func getTab() string {

	return `	`
}

func getColComment(col jsql.TableSchema) string {

	return getTab() + `// ` + strToUCC(col.ColumnName) + ` ` + col.ColumnComment
}

func getColType(col jsql.TableSchema) string {

	t := strings.ToLower(col.DataType)
	switch t {
	case "bit":
		return "uint8"
	case "tinyint":
		return "int16"
	case "smallint":
		fallthrough
	case "mediumint":
		return "int32"
	case "int":
		return "int"
	case "bigint":
		return "int64"
	case "smallmoney":
		fallthrough
	case "float":
		return "float32"
	case "money":
		fallthrough
	case "real":
		fallthrough
	case "double":
		fallthrough
	case "numeric":
		fallthrough
	case "decimal":
		return "float64"
	case "date":
		fallthrough
	case "datetime":
		fallthrough
	case "datetime2":
		fallthrough
	case "smalldatetime":
		fallthrough
	case "time":
		fallthrough
	case "timestamp":
		return "time.Time"
	case "blob":
		fallthrough
	case "binary":
		fallthrough
	case "varbinary":
		fallthrough
	case "image":
		return "[]byte"
	default:
		if b, err := regexp.MatchString("^timestamp", t); err == nil && b {

			return "time.Time"
		} else if err != nil {

			jPrint(err)
		}
		if b, err := regexp.MatchString("^number(.*,.*)", t); err == nil && b {

			return "float64"
		} else if err != nil {

			jPrint(err)
		}
		if b, err := regexp.MatchString("^number", t); err == nil && b {

			return "int64"
		} else if err != nil {

			jPrint(err)
		}
		return "string"
	}
}

func getColJson(col jsql.TableSchema) string {

	gorm := ""
	if root.Gorm {

		d := len(col.DataDefault) > 0
		k := col.PrimaryKey != nil
		i := col.IsIdentity == "YES"
		g := d || k
		if g {

			gorm = ` gorm:"`
		}
		t := false
		if d {

			gorm += `default:` + col.DataDefault
			t = true
		}
		if k {

			if t {

				gorm += `;`
			}
			gorm += `primarykey`
			if !i {

				gorm += `;autoIncrement:false`
			}
		}
		if g {

			gorm += `"`
		}
	}
	return "`json:\"" + col.ColumnName + "\"" + gorm + "`"
}
