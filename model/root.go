// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package model

type Root struct {
	Name       string
	Project    bool
	Model      bool
	Schedule   bool
	Gorm       bool
	Service    bool
	Datasource string
	Table      string
	Job        string
	Test       bool
}
