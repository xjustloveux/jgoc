// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"github.com/xjustloveux/jgo/jcron"
	"github.com/xjustloveux/jgo/jfile"
	"github.com/xjustloveux/jgo/jsql"
)

func checkConfig() error {

	if !root.Model && !root.Schedule {

		return nil
	}
	if err := checkConfigFile(); err != nil {

		return err
	}
	if root.Model {

		if err := jsql.Init(); err != nil {

			return err
		}
	}
	if root.Schedule {

		if err := jcron.Init(); err != nil {

			return err
		}
	}
	return nil
}

func checkConfigFile() error {

	fileName := "config.json"
	path := fmt.Sprint("./config/", fileName)
	if exist, err := jfile.Exist(path); err != nil {

		return err
	} else if exist {

		jPrint(fmt.Sprint(path, " exist"))
		return nil
	} else if !exist {

		jPrint(fmt.Sprint(path, " not exist"))
	}
	fileName = "config.yaml"
	path = fmt.Sprint("./config/", fileName)
	if exist, err := jfile.Exist(path); err != nil {

		return err
	} else if exist {

		if root.Model {

			jsql.SetFormat(jfile.Yaml)
			jsql.SetFileName(fileName)
		}
		if root.Schedule {

			jcron.SetFormat(jfile.Yaml)
			jcron.SetFileName(fileName)
		}
		jPrint(fmt.Sprint(path, " exist"))
		return nil
	}
	return jError("not found any config file")
}
