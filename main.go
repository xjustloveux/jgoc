// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.
package main

import (
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/godror/godror"
	"github.com/xjustloveux/jgo/jfile"
	"github.com/xjustloveux/jgoc/cmd"
	"github.com/xjustloveux/jgoc/yaml"
	"os"
)

func main() {

	jfile.RegisterCodec(jfile.Yaml.String(), yaml.Codec{})
	cmd.Execute(os.Args[1:])
}
