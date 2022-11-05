// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"strings"
)

func strToL(str string) string {

	arr := strings.Split(str, "_")
	s := ""
	for _, v := range arr {

		s += strings.ToLower(v)
	}
	return s
}

func strToFU(str string) string {

	if len(str) > 0 {

		return fmt.Sprint(strings.ToUpper(str[0:1]), strings.ToLower(str[1:]))
	}
	return ""
}

func strToLCC(str string) string {

	arr := strings.Split(str, "_")
	s := ""
	for i, v := range arr {

		if i == 0 {

			s += strings.ToLower(v)
		} else {

			s += strToFU(v)
		}
	}
	return s
}

func strToUCC(str string) string {

	arr := strings.Split(str, "_")
	s := ""
	for _, v := range arr {

		s += strToFU(v)
	}
	return s
}
