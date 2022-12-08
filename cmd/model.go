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

		module = ModelImport
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
	return fmt.Sprintf(Model, dsName, module, tableName, tableComment, tableName, col)
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
			if b, err := regexp.MatchString("^time", t); err == nil && b {

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

	pointer := ""
	if root.Pointer {

		pointer = "*"
	}
	t := strings.ToLower(col.DataType)
	switch t {
	case "bool":
		fallthrough
	case "boolean":
		return fmt.Sprint(pointer, "bool")
	case "bit":
		return fmt.Sprint(pointer, "uint8")
	case "tinyint":
		return fmt.Sprint(pointer, "int16")
	case "smallint":
		fallthrough
	case "mediumint":
		return fmt.Sprint(pointer, "int32")
	case "int":
		fallthrough
	case "integer":
		return fmt.Sprint(pointer, "int")
	case "bigint":
		return fmt.Sprint(pointer, "int64")
	case "smallmoney":
		fallthrough
	case "float":
		return fmt.Sprint(pointer, "float32")
	case "money":
		fallthrough
	case "real":
		fallthrough
	case "double":
		fallthrough
	case "numeric":
		fallthrough
	case "decimal":
		return fmt.Sprint(pointer, "float64")
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
		fallthrough
	case "bytea":
		return "[]byte"
	case "json":
		return "interface{}"
	default:
		if b, err := regexp.MatchString("^time", t); err == nil && b {

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
	sts := ""
	if root.Gorm {

		v := defValCheck(col)
		d := len(v) > 0
		k := col.PrimaryKey != nil
		i := col.IsIdentity == "YES"
		g := d || k
		if g {

			gorm = ` gorm:"`
		}
		t := false
		if d {

			gorm += `default:` + v
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
	} else {

		sts = fmt.Sprint(` structs:"`, col.ColumnName, `"`)
	}
	return fmt.Sprint("`json:\"", col.ColumnName, "\"", sts, gorm, "`")
}

func defValCheck(col jsql.TableSchema) string {

	val := col.DataDefault
	t := strings.ToLower(col.DataType)
	if strings.HasPrefix(val, "nextval") {

		return ""
	}
	switch t {
	case "tinyint":
		fallthrough
	case "smallint":
		fallthrough
	case "mediumint":
		fallthrough
	case "int":
		fallthrough
	case "integer":
		fallthrough
	case "bigint":
		fallthrough
	case "smallmoney":
		fallthrough
	case "float":
		fallthrough
	case "money":
		fallthrough
	case "real":
		fallthrough
	case "double":
		fallthrough
	case "numeric":
		fallthrough
	case "decimal":
		val = strings.Replace(val, "(", "", -1)
		val = strings.Replace(val, ")", "", -1)
	}
	return val
}
