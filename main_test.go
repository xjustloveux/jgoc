// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.
package main

import (
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-oci8"
	"github.com/xjustloveux/jgo/jfile"
	"github.com/xjustloveux/jgoc/cmd"
	"github.com/xjustloveux/jgoc/yaml"
	"os"
	"testing"
)

func TestCreate(t *testing.T) {

	path, _ := os.Getwd()
	if err := os.Chdir(path + "/test"); err != nil {

		t.Error(err)
		return
	}
	args := []string{"--name", "test", "--pro", "--mod", "--srv", "--sch", "--test"}
	jfile.RegisterCodec(jfile.Yaml.String(), yaml.Codec{})
	cmd.Execute(args)
}
