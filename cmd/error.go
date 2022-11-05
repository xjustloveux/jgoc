// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package cmd

import (
	"errors"
	"fmt"
)

func jError(str string) error {

	return errors.New(fmt.Sprint(name, ": ", str))
}
