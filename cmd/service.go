// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"github.com/xjustloveux/jgoc/model"
	"sort"
	"strings"
)

func checkService() error {

	if !root.Service {

		return nil
	}
	/*
		service
		|---dsName
			|---tableName.go
	*/
	path := "service"
	if err := checkExistAndCreateFolder(path); err != nil {

		return err
	}
	path = fmt.Sprint(path, "/")
	if err := createServiceDS(path); err != nil {

		return err
	}
	return nil
}

func createServiceDS(path string) error {

	for _, v := range datasource {

		dsModName := fmt.Sprint(strToL(v.Name), "mod")
		dsSrvName := fmt.Sprint(strToL(v.Name), "srv")
		dsPath := fmt.Sprint(path, dsSrvName)
		if err := checkExistAndCreateFolder(dsPath); err != nil {

			return err
		}
		for _, t := range v.Tables {

			create := len(root.Table) <= 0 || t.Name == root.Table
			if create {

				if err := createServiceTable(fmt.Sprint(dsPath, "/"), dsModName, dsSrvName, v.Name, t); err != nil {

					return err
				}
			}
		}
	}
	return nil
}

func createServiceTable(path, dsModName, dsSrvName, dsName string, table model.Table) error {

	fileName := fmt.Sprint(strToL(table.Name), ".go")
	filePath := fmt.Sprint(path, fileName)
	jPrint(fmt.Sprint("create ", filePath, "..."))
	if err := createFile(filePath, getServiceTableContent(dsModName, dsSrvName, dsName, table)); err != nil {

		return err
	}
	return nil
}

func getServiceTableContent(dsModName, dsSrvName, dsName string, table model.Table) string {

	if root.Gorm {

		return getGormContent(dsModName, dsSrvName, dsName, table)
	}
	return getJGoContent(dsModName, dsSrvName, dsName, table)
}

func getJGoContent(dsModName, dsSrvName, dsName string, table model.Table) string {

	tableNameUCC := strToUCC(table.Name)
	tableNameLCC := strToLCC(table.Name)
	tableComment := table.Schema[0].TableComment
	tableModel := fmt.Sprint(dsModName, ".", tableNameUCC)
	module := []string{fmt.Sprintf(SrvImport1, root.Name, dsModName), fmt.Sprintf(SrvImport2, ModuleStructs), fmt.Sprintf(SrvJGoImport, ModuleJGo)}
	sort.Slice(module, func(i, j int) bool {
		re := []string{" ", "　", "	", "_", `
`}
		a := module[i]
		b := module[j]
		for _, v := range re {
			a = strings.Replace(a, v, "", -1)
			b = strings.Replace(b, v, "", -1)
		}
		return a < b
	})
	mod := ""
	for _, v := range module {
		mod += v
	}
	srv := SrvJGo
	srv = strings.Replace(srv, `%UCC%`, tableNameUCC, -1)
	srv = strings.Replace(srv, `%LCC%`, tableNameLCC, -1)
	srv = strings.Replace(srv, `%model%`, tableModel, -1)
	return fmt.Sprintf(srv, dsSrvName, mod, tableComment, dsName, table.Name)
}

func getGormContent(dsModName, dsSrvName, dsName string, table model.Table) string {

	tableNameUCC := strToUCC(table.Name)
	tableNameLCC := strToLCC(table.Name)
	tableComment := table.Schema[0].TableComment
	tableModel := fmt.Sprint(dsModName, ".", tableNameUCC)
	module := []string{fmt.Sprintf(SrvImport1, root.Name, dsModName), fmt.Sprintf(SrvImport2, "fmt"), fmt.Sprintf(SrvImport2, ModuleGorm)}
	sort.Slice(module, func(i, j int) bool {
		re := []string{" ", "　", "	", "_", `
`}
		a := module[i]
		b := module[j]
		for _, v := range re {
			a = strings.Replace(a, v, "", -1)
			b = strings.Replace(b, v, "", -1)
		}
		return a < b
	})
	mod := ""
	for _, v := range module {
		mod += v
	}
	srv := SrvGorm
	srv = strings.Replace(srv, `%UCC%`, tableNameUCC, -1)
	srv = strings.Replace(srv, `%LCC%`, tableNameLCC, -1)
	srv = strings.Replace(srv, `%model%`, tableModel, -1)
	return fmt.Sprintf(srv, dsSrvName, mod, tableComment, table.Name)
}
