// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"github.com/xjustloveux/jgo/jfile"
	"github.com/xjustloveux/jgoc/model"
)

func checkYaml() error {

	if !root.Project || !root.Yaml {

		return nil
	}
	/*
		middleware
		|---yaml
			|---codec.go
			|---yaml2.go
	*/
	codec := model.PathInfo{
		Name:     "codec.go",
		IsFolder: false,
		Child:    nil,
	}
	yaml2 := model.PathInfo{
		Name:     "yaml2.go",
		IsFolder: false,
		Child:    nil,
	}
	yaml := model.PathInfo{
		Name:     "yaml",
		IsFolder: true,
		Child:    []model.PathInfo{codec, yaml2},
	}
	middleware := model.PathInfo{
		Name:     "middleware",
		IsFolder: true,
		Child:    []model.PathInfo{yaml},
	}
	r := model.PathInfo{
		Name:     "",
		IsFolder: true,
		Child:    []model.PathInfo{middleware},
	}
	return createYaml(r)
}

func createYaml(info model.PathInfo) error {

	for _, v := range info.Child {

		if err := checkYamlInfo("", v); err != nil {

			return err
		}
	}
	return nil
}

func checkYamlInfo(path string, info model.PathInfo) error {

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
		content := getYamlFileContent(filePath)
		jPrint(fmt.Sprint("create ", filePath, "..."))
		if err := createFile(filePath, content); err != nil {

			return err
		}
	}
	if info.Child != nil {

		for _, v := range info.Child {

			if err := checkYamlInfo(fmt.Sprint(path, info.Name, "/"), v); err != nil {

				return err
			}
		}
	}
	return nil
}

func getYamlFileContent(path string) string {

	switch path {
	case "middleware/yaml/codec.go":
		return YamlCodec
	case "middleware/yaml/yaml2.go":
		return fmt.Sprintf(YamlYaml2, ModuleYaml)
	}
	return ""
}
