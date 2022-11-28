// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"github.com/xjustloveux/jgo/jcast"
	"github.com/xjustloveux/jgo/jfile"
	"github.com/xjustloveux/jgo/jsql"
	"github.com/xjustloveux/jgoc/model"
	"sort"
	"strings"
)

func checkProject() error {

	if !root.Project {

		return nil
	}
	/*
		controller
		|---a
			|---a101ctr
				|---a101.go
			|---router.go
		|---router.go
		global
		|---variable.go
		service
		|---a
			|---a101srv
				|---a101.go
		main.go
	*/
	a101ctrgo := model.PathInfo{
		Name:     "a101.go",
		IsFolder: false,
		Child:    nil,
	}
	a101ctr := model.PathInfo{
		Name:     "a101ctr",
		IsFolder: true,
		Child:    []model.PathInfo{a101ctrgo},
	}
	arouter := model.PathInfo{
		Name:     "router.go",
		IsFolder: false,
		Child:    nil,
	}
	actr := model.PathInfo{
		Name:     "a",
		IsFolder: true,
		Child:    []model.PathInfo{a101ctr, arouter},
	}
	router := model.PathInfo{
		Name:     "router.go",
		IsFolder: false,
		Child:    nil,
	}
	controller := model.PathInfo{
		Name:     "controller",
		IsFolder: true,
		Child:    []model.PathInfo{actr, router},
	}
	variable := model.PathInfo{
		Name:     "variable.go",
		IsFolder: false,
		Child:    nil,
	}
	global := model.PathInfo{
		Name:     "global",
		IsFolder: true,
		Child:    []model.PathInfo{variable},
	}
	a101srvgo := model.PathInfo{
		Name:     "a101.go",
		IsFolder: false,
		Child:    nil,
	}
	a101srv := model.PathInfo{
		Name:     "a101srv",
		IsFolder: true,
		Child:    []model.PathInfo{a101srvgo},
	}
	asrv := model.PathInfo{
		Name:     "a",
		IsFolder: true,
		Child:    []model.PathInfo{a101srv},
	}
	service := model.PathInfo{
		Name:     "service",
		IsFolder: true,
		Child:    []model.PathInfo{asrv},
	}
	main := model.PathInfo{
		Name:     "main.go",
		IsFolder: false,
		Child:    nil,
	}
	r := model.PathInfo{
		Name:     "",
		IsFolder: true,
		Child:    []model.PathInfo{controller, service, global, main},
	}
	return createProject(r)
}

func createProject(info model.PathInfo) error {

	for _, v := range info.Child {

		if err := checkInfo("", v); err != nil {

			return err
		}
	}
	return nil
}

func checkInfo(path string, info model.PathInfo) error {

	if info.IsFolder {

		folderPath := fmt.Sprint(path, info.Name)
		if exist, err := jfile.Exist(folderPath); err != nil {

			return err
		} else if !exist {

			jPrint(fmt.Sprint(folderPath, " not exist"))
			jPrint(fmt.Sprint("create folder(", folderPath, ")..."))
			if err = createFolder(folderPath); err != nil {

				return err
			}
			if exist, err = jfile.Exist(folderPath); err != nil {

				return err
			} else if !exist {

				return jError(fmt.Sprint("create folder fail(", folderPath, ")"))
			} else {

				jPrint(fmt.Sprint("create folder success(", folderPath, ")"))
			}
		} else {

			jPrint(fmt.Sprint(folderPath, " exist"))
		}
	} else {

		filePath := fmt.Sprint(path, info.Name)
		content := getProjectFileContent(filePath)
		jPrint(fmt.Sprint("create ", filePath, "..."))
		if err := createFile(filePath, content); err != nil {

			return err
		}
	}
	if info.Child != nil {

		for _, v := range info.Child {

			if err := checkInfo(fmt.Sprint(path, info.Name, "/"), v); err != nil {

				return err
			}
		}
	}
	return nil
}

func getProjectFileContent(path string) string {

	switch path {
	case "controller/a/a101ctr/a101.go":
		module := []string{fmt.Sprintf(ProCtrA101Import1, root.Name), fmt.Sprintf(ProCtrA101Import2, ModuleGin), ProCtrA101Import3}
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
		return fmt.Sprintf(ProCtrA101, mod)
	case "controller/a/router.go":
		return fmt.Sprintf(ProCtrARouter, root.Name, root.Name)
	case "controller/router.go":
		return fmt.Sprintf(ProCtrRouter, root.Name, root.Name)
	case "global/variable.go":
		return fmt.Sprintf(ProVar, ModuleGin)
	case "service/a/a101srv/a101.go":
		return ProSrvA101
	case "main.go":
		module := []string{fmt.Sprintf(ProMainImport1, root.Name), ProMainImport2}
		yamlInit := ""
		if root.Yaml {

			module = append(module, fmt.Sprintf(ProMainImportYaml, root.Name))
			yamlInit = ProMainInitYaml
		}
		jobInit := ""
		if root.Schedule {

			module = append(module, fmt.Sprintf(ProMainImportJob1, root.Name))
			module = append(module, fmt.Sprintf(ProMainImportJob2, ModuleJGo))
			if root.Yaml {

				jobInit = fmt.Sprintf(ProMainInitJob, ProMainInitJobYaml)
			} else {

				jobInit = fmt.Sprintf(ProMainInitJob, "")
			}
		}
		sqlInit := ""
		if root.Service && !root.Gorm {

			module = append(module, fmt.Sprintf(ProMainImportJGoSql, ModuleJGo))
			if root.Yaml {

				sqlInit = fmt.Sprintf(ProMainInitSql, ProMainInitSqlYaml)
			} else {

				sqlInit = fmt.Sprintf(ProMainInitSql, "")
			}
			ds := jsql.GetDataSource()
			for k, v := range ds {

				if len(root.Datasource) <= 0 || root.Datasource == k {

					if m, err := jcast.StringMapString(v); err == nil {

						switch t, _ := jsql.ParseDBType(m["type"]); t {
						case jsql.MySql:
							module = append(module, fmt.Sprintf(ProMainImportSql, ModuleMySql))
						case jsql.MSSql:
							module = append(module, fmt.Sprintf(ProMainImportSql, ModuleMSSql))
						case jsql.Oracle:
							module = append(module, fmt.Sprintf(ProMainImportSql, ModuleOracle))
						case jsql.PostgreSql:
							module = append(module, fmt.Sprintf(ProMainImportSql, ModulePostgreSql))
						}
					}
				}
			}
		}
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
		return fmt.Sprintf(ProMain, mod, yamlInit, sqlInit, jobInit)
	}
	return ""
}
