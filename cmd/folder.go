// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"github.com/xjustloveux/jgo/jfile"
	"os"
	"path/filepath"
	"strings"
)

func checkFolder() error {

	if !root.CreateChk {

		return nil
	}
	var err error
	var path, folder string
	if path, err = os.Getwd(); err != nil {

		return err
	} else {

		_, folder = filepath.Split(path)
	}
	arr := strings.Split(root.Name, "/")
	var folderName string
	if l := len(arr); l > 1 {

		folderName = arr[l-1]
	} else {

		folderName = root.Name
	}
	if folder == folderName {

		return nil
	}
	var exist bool
	if exist, err = jfile.Exist(folderName); err != nil {

		return err
	} else if exist {

		return chdir(path, folderName)
	}
	jPrint(fmt.Sprint("create folder(", folderName, ")"))
	if err = os.Mkdir(folderName, 0755); err != nil {

		return err
	}
	if exist, err = jfile.Exist(folderName); err != nil {

		return err
	} else if !exist {

		return jError(fmt.Sprint("create folder fail(", folderName, ")"))
	} else {

		jPrint(fmt.Sprint("create folder success(", folderName, ")"))
	}
	return chdir(path, folderName)
}

func chdir(path, folderName string) error {

	if strings.HasPrefix(path, "/") {

		path = fmt.Sprint("/", strings.Trim(path, "/"))
	} else {

		path = strings.Trim(path, "/")
	}
	path = strings.Trim(path, "\\")
	newPath := fmt.Sprint(path, "/", folderName)
	jPrint(fmt.Sprint("changes the current working directory to ", newPath))
	return os.Chdir(newPath)
}
