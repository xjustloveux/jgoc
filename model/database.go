// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package model

import "github.com/xjustloveux/jgo/jsql"

type Datasource struct {
	Name   string
	Tables []Table
}

type Table struct {
	Name   string
	Schema []jsql.TableSchema
}
