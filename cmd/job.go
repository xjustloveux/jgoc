// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"github.com/xjustloveux/jgo/jcast"
	"github.com/xjustloveux/jgo/jfile"
	"io"
	"os"
	"regexp"
	"strings"
)

func checkJob() error {

	if !root.Schedule {

		return nil
	}
	jPrint("load schedule...")
	if err := loadSchedule(); err != nil {

		return err
	}
	/*
		job
		|---init.go
		|---job001.go
	*/
	path := "job"
	if err := checkExistAndCreateFolder(path); err != nil {

		return err
	}
	path = fmt.Sprint(path, "/init.go")
	initContent := ""
	if exist, err := jfile.Exist(path); err != nil {

		return err
	} else if exist {

		if initContent, err = checkJobInitContent(path); err != nil {

			return err
		}
	} else {

		initContent = getJobInitContent()
	}
	jPrint(fmt.Sprint("create ", path, "..."))
	if err := createFile(path, initContent); err != nil {

		return err
	}
	if err := createJobs(); err != nil {

		return err
	}
	return nil
}

func checkJobInitContent(path string) (string, error) {

	if file, err := os.OpenFile(path, os.O_RDONLY, 0755); err != nil {

		return "", err
	} else {

		defer func() {

			if e := file.Close(); e != nil {

				fmt.Println(e)
			}
		}()

		var b []byte
		if b, err = io.ReadAll(file); err != nil {

			return "", err
		}
		content := jcast.String(b)
		ss := "jobs := []job{"
		si := strings.Index(content, ss)
		ei := strings.Index(content, `
	}
	for _, v := range jobs {`)
		if si < 0 || ei < 0 || ei <= si {

			return "", jError(fmt.Sprint(path, " file not up to specification"))
		}
		chkStr := content[si+len(ss) : ei]
		chkJob := make([]string, 0)
		for _, v := range jobs {

			check := true
			for _, ck := range chkJob {

				if ck == v.JobName {

					check = false
					break
				}
			}
			if check {

				var match bool
				if match, err = regexp.MatchString(fmt.Sprint(".*&", strToL(v.JobName), "{}.*"), chkStr); err != nil {

					return "", err
				}
				if !match {

					chkJob = append(chkJob, v.JobName)
				}
			}
		}
		add := ""
		for _, v := range chkJob {

			add += `
		&` + strToL(v) + `{},`
		}
		return fmt.Sprint(content[:si+len(ss)], chkStr, add, content[ei:]), nil
	}
}

func getJobInitContent() string {

	chkJob := make([]string, 0)
	for _, v := range jobs {

		add := true
		for _, ck := range chkJob {

			if ck == v.JobName {

				add = false
				break
			}
		}
		if add {

			chkJob = append(chkJob, v.JobName)
		}
	}
	add := ""
	for _, v := range chkJob {

		add += `
		&` + strToL(v) + `{},`
	}
	return fmt.Sprintf(JobInit, ModuleJGo, add)
}

func createJobs() error {

	list := make([]string, 0)
	for _, v := range jobs {

		n := strToL(v.JobName)
		add := true
		for _, ck := range list {

			if ck == n {

				add = false
				break
			}
		}
		if add {

			list = append(list, n)
		}
	}
	for _, v := range list {

		path := fmt.Sprint("job/", v, ".go")
		jPrint(fmt.Sprint("create ", path, "..."))
		if err := createFile(path, getJobContent(v)); err != nil {

			return err
		}
	}
	return nil
}

func getJobContent(n string) string {

	return strings.Replace(Job, `%n%`, n, -1)
}
