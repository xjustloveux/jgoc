// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"github.com/xjustloveux/jgo/jfile"
	"github.com/xjustloveux/jgoc/model"
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
		job
		|---init.go
		|---job001.go
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
		return `package a101ctr

import (
	"` + root.Name + `/service/a/a101srv"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Example(ctx *gin.Context) {

	message := a101srv.DoSomething()
	ctx.JSON(http.StatusOK, gin.H{"success": true, "message": message})
}
`
	case "controller/a/router.go":
		return `package a

import (
	"` + root.Name + `/controller/a/a101ctr"
	"` + root.Name + `/global"
)

func Init() {
	a := global.Router.Group("a")
	{
		a101 := a.Group("a101")
		{
			a101.GET("", a101ctr.Example)
		}
	}
}
`
	case "controller/router.go":
		return `package controller

import (
	"` + root.Name + `/controller/a"
	"` + root.Name + `/global"
)

func Init() error {
	a.Init()
	return global.Router.Run(":8080")
}
`
	case "global/variable.go":
		return `package global

import "github.com/gin-gonic/gin"

var (
	Router = gin.Default()
)
`
	case "service/a/a101srv/a101.go":
		return `package a101srv

func DoSomething() string {

	// Do Something
	return "Example"
}
`
	case "main.go":
		jobModule := ""
		jobInit := ""
		if root.Schedule {

			jobModule = `
	"` + root.Name + `/job"`
			jobInit = `
	if err := job.Init(); err != nil {
		
		fmt.Println(err)
	}`
		}
		return `package main

import (
	"` + root.Name + `/controller"` + jobModule + `
	"fmt"
)

func main() {
	` + jobInit + `
	if err := controller.Init(); err != nil {
		
		fmt.Println(err)
	}
}
`
	}
	return ""
}
