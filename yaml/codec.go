// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package yaml

type Codec struct{}

func (Codec) Encode(v map[string]interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

func (Codec) Decode(b []byte, v map[string]interface{}) error {
	return yaml.Unmarshal(b, &v)
}
