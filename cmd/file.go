// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"github.com/xjustloveux/jgo/jfile"
	"os"
)

func createFolder(folderPath string) error {

	return os.MkdirAll(folderPath, 0755)
}

func createFile(filePath, content string) error {

	if file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755); err != nil {

		return err
	} else {

		defer func() {

			if e := file.Close(); e != nil {

				fmt.Println(e)
			}
		}()
		if err = file.Truncate(0); err != nil {

			return err
		}
		if _, err = file.Seek(0, 0); err != nil {

			return err
		}
		if _, err = file.Write([]byte(content)); err != nil {

			return err
		}
	}
	return nil
}

func checkExistAndCreateFolder(path string) error {

	if exist, err := jfile.Exist(path); err != nil {

		return err
	} else if !exist {

		jPrint(fmt.Sprint(path, " not exist"))
		jPrint(fmt.Sprint("create folder(", path, ")..."))
		if err = createFolder(path); err != nil {

			return err
		}
		if exist, err = jfile.Exist(path); err != nil {

			return err
		} else if !exist {

			return jError(fmt.Sprint("create folder fail(", path, ")"))
		} else {

			jPrint(fmt.Sprint("create folder success(", path, ")"))
		}
	} else {

		jPrint(fmt.Sprint(path, " exist"))
	}
	return nil
}
