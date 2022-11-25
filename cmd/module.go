// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"github.com/xjustloveux/jgo/jcast"
	"github.com/xjustloveux/jgo/jfile"
	"github.com/xjustloveux/jgo/jsql"
	"os/exec"
	"regexp"
)

func checkModule() error {

	file := "go.mod"
	if exist, err := jfile.Exist(file); err != nil {

		return err
	} else if !exist {

		jPrint(fmt.Sprint(file, " not exist"))
		if err = runComm(exec.Command("go", "mod", "init", root.Name)); err != nil {

			return err
		}
		if exist, err = jfile.Exist(file); err != nil {

			return err
		} else if !exist {

			return jError("go mod init fail")
		} else {

			jPrint("go mod init success")
		}
	} else {

		jPrint(fmt.Sprint(file, " exist"))
	}
	if modFile, err := loadModFile(); err != nil {

		return err
	} else {

		return checkModInit(modFile)
	}
}

func loadModFile() (string, error) {

	if b, err := jfile.Load("go.mod"); err != nil {

		return "", err
	} else {

		return jcast.String(b), nil
	}
}

func checkModInit(str string) error {

	pkg := make([]string, 0)
	if root.Project {

		pkg = append(pkg, "github.com/gin-gonic/gin")

		if root.Service && !root.Gorm {

			ds := jsql.GetDataSource()
			for k, v := range ds {

				if len(root.Datasource) <= 0 || root.Datasource == k {

					if m, err := jcast.StringMapString(v); err != nil {

						return err
					} else {

						switch t, _ := jsql.ParseDBType(m["type"]); t {
						case jsql.MySql:
							pkg = append(pkg, "github.com/go-sql-driver/mysql")
						case jsql.MSSql:
							pkg = append(pkg, "github.com/denisenkom/go-mssqldb")
						case jsql.Oracle:
							pkg = append(pkg, "github.com/godror/godror")
						case jsql.PostgreSql:
							pkg = append(pkg, "github.com/lib/pq")
						}
					}
				}
			}
		}
	}
	if root.Service {

		if root.Gorm {

			pkg = append(pkg, "gorm.io/gorm")
		} else {

			pkg = append(pkg, "github.com/xjustloveux/jgo")
			pkg = append(pkg, "github.com/fatih/structs")
		}
	}
	if root.Schedule && (!root.Service || root.Gorm) {

		pkg = append(pkg, "github.com/xjustloveux/jgo")
	}
	l := len(fmt.Sprint("module ", root.Name))
	ckStr := str[l:]
	for _, v := range pkg {

		if b, err := regexp.MatchString(v, ckStr); err != nil {

			return err
		} else if b {

			jPrint(fmt.Sprint("package(", v, ") exist"))
		} else {

			jPrint(fmt.Sprint("package(", v, ") not exist"))
			if err = runComm(exec.Command("go", "get", "-u", v)); err != nil {

				return err
			}
		}
	}
	return nil
}
