// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package cmd

import (
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/godror/godror"
	"github.com/xjustloveux/jgo/jfile"
	"github.com/xjustloveux/jgoc/yaml"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCreate(t *testing.T) {

	if err := jError("test"); err == nil {

		t.Error("jError must be return error")
	}
	path, _ := os.Getwd()
	testPath := strings.Replace(path, "cmd", "test", -1)
	rootPath, _ := filepath.Split(path)
	if err := os.Chdir(rootPath); err != nil {

		t.Error(err)
		return
	}
	args := []string{"--name", "test2", "--pro", "--test"}
	Execute(args)
	if err := os.Chdir(rootPath); err != nil {

		t.Error(err)
		return
	}
	Execute(args)
	Execute(args)
	if err := os.Chdir(testPath); err != nil {

		t.Error(err)
		return
	}
	jfile.RegisterCodec(jfile.Yaml.String(), yaml.Codec{})
	args = []string{"--name", "test", "--pro", "--test"}
	Execute(args)
	args = []string{"--name", "test", "--pro", "--mod", "--test"}
	Execute(args)
	args = []string{"--name", "test", "--pro", "--mod", "--srv", "--test"}
	Execute(args)
	args = []string{"--name", "test", "--pro", "--mod", "--srv", "--sch", "--test"}
	Execute(args)
	args = []string{"--name", "test", "--mod", "--test"}
	Execute(args)
	args = []string{"--name", "test", "--mod", "--srv", "--test"}
	Execute(args)
	args = []string{"--name", "test", "--sch", "--test"}
	Execute(args)
	args = []string{"--name", "test", "--env", "dev", "--pro", "--mod", "--srv", "--sch", "--test"}
	Execute(args)
	args = []string{"--name", "test", "--pro", "--mod", "--srv", "--ds", "TestMSSql", "--table", "TEST", "--gorm", "--sch", "--job", "job002", "--test"}
	Execute(args)
}
