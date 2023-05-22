// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package yaml

import (
	yamlv3 "gopkg.in/yaml.v3"
)

var yaml = struct {
	Marshal   func(in interface{}) (out []byte, err error)
	Unmarshal func(in []byte, out interface{}) (err error)
}{
	Marshal:   yamlv3.Marshal,
	Unmarshal: yamlv3.Unmarshal,
}
