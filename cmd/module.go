// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"github.com/xjustloveux/jgo/jcast"
	"github.com/xjustloveux/jgo/jfile"
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
	for _, v := range pkg {

		if b, err := regexp.MatchString(v, str); err != nil {

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
