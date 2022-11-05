// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"io"
	"os/exec"
)

func runComm(cmd *exec.Cmd) error {

	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {

		return err
	}
	if err = cmd.Start(); err != nil {

		return err
	}
	for {

		tmp := make([]byte, 8192)
		_, err = stdout.Read(tmp)
		fmt.Print(string(tmp))
		if err != nil || err == io.EOF {

			break
		}
	}
	if err = cmd.Wait(); err != nil {

		return err
	}
	return nil
}
