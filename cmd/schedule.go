// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"github.com/xjustloveux/jgo/jcron"
)

func loadSchedule() error {

	if len(root.Job) > 0 {

		jobs = make([]*jcron.SchInfo, 1)
		validName := false
		for _, v := range jcron.GetScheduleInfo() {

			if v.JobName == root.Job {

				jobs[0] = v
				validName = true
				break
			}
		}
		if !validName {

			return jError(fmt.Sprint("job name(", root.Job, ") not exist in schedule"))
		}
	} else {

		jobs = jcron.GetScheduleInfo()
	}
	return nil
}
