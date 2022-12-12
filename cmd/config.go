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

		root.Yaml = false
		root.CreateChk = true
		return nil
	}
	root.CreateChk = false
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

	root.Yaml = false
	fileName := "config.json"
	path := fmt.Sprint("./config/", fileName)
	if exist, err := jfile.Exist(path); err != nil {

		return err
	} else if exist {

		jPrint(fmt.Sprint(path, " exist"))
		if len(root.Env) > 0 {

			fileName = fmt.Sprint("config-", root.Env, ".json")
			path = fmt.Sprint("./config/", fileName)
			if exist, err = jfile.Exist(path); err != nil {

				return err
			} else if exist {

				jPrint(fmt.Sprint(path, " exist"))
				if root.Model {

					jsql.SetEnvVal(root.Env)
				}
				if root.Schedule {

					jcron.SetEnvVal(root.Env)
				}
				return nil
			} else if !exist {

				jPrint(fmt.Sprint(path, " not exist"))
			}
		} else {

			return nil
		}
	} else if !exist {

		jPrint(fmt.Sprint(path, " not exist"))
	}
	root.Yaml = true
	fileName = "config.yaml"
	path = fmt.Sprint("./config/", fileName)
	if exist, err := jfile.Exist(path); err != nil {

		return err
	} else if exist {

		jPrint(fmt.Sprint(path, " exist"))
		if len(root.Env) > 0 {

			fileName = fmt.Sprint("config-", root.Env, ".yaml")
			path = fmt.Sprint("./config/", fileName)
			if exist, err = jfile.Exist(path); err != nil {

				return err
			} else if exist {

				jPrint(fmt.Sprint(path, " exist"))
				if root.Model {

					jsql.SetFormat(jfile.Yaml)
					jsql.SetFileName("config.yaml")
					jsql.SetEnvVal(root.Env)
				}
				if root.Schedule {

					jcron.SetFormat(jfile.Yaml)
					jcron.SetFileName("config.yaml")
					jcron.SetEnvVal(root.Env)
				}
				return nil
			} else if !exist {

				jPrint(fmt.Sprint(path, " not exist"))
			}
		}
		if root.Model {

			jsql.SetFormat(jfile.Yaml)
			jsql.SetFileName("config.yaml")
		}
		if root.Schedule {

			jcron.SetFormat(jfile.Yaml)
			jcron.SetFileName("config.yaml")
		}
		return nil
	}
	return jError("not found any config file")
}
