// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xjustloveux/jgo/jcron"
	"github.com/xjustloveux/jgoc/model"
	"os"
)

var (
	name       = "jgoc"
	root       = &model.Root{}
	datasource []*model.Datasource
	jobs       []*jcron.SchInfo
)

func Execute(args []string) {

	ccmd := &cobra.Command{
		Use:   name,
		Short: `create project, model or schedule as fast and easily as possible`,
		Long: `JGoC provides an easier way to create Go project, model and schedule.
JGoC goal is to simplify the created project framework, model and schedule steps while providing variant customization options for all steps.
Check out github for more information: https://github.com/xjustloveux/jgof`,
		Version:       "v1.0.0",
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {

			return execute()
		},
	}
	ccmd.Flags().StringVar(&root.Name, "name", "", "module name")
	if err := ccmd.MarkFlagRequired("name"); err != nil {

		fmt.Println(err)
		return
	}
	ccmd.Flags().BoolVar(&root.Project, "pro", false, "created project framework")
	ccmd.Flags().BoolVar(&root.Model, "mod", false, "created database model, need config/config.json or config/config.yaml file, configuration refer to https://github.com/xjustloveux/jgo#configuration")
	ccmd.Flags().BoolVar(&root.Schedule, "sch", false, "created schedule, need config/config.json or config/config.yaml file, configuration refer to https://github.com/xjustloveux/jgo#configuration-1")
	ccmd.Flags().BoolVar(&root.Pointer, "pointer", false, "columns of numeric type will be converted to pointer type when creating the model")
	ccmd.Flags().BoolVar(&root.Gorm, "gorm", false, `create model and service with gorm, required flag(s) "mod"`)
	ccmd.Flags().BoolVar(&root.Service, "srv", false, `created model service, required flag(s) "mod"`)
	ccmd.Flags().StringVar(&root.Datasource, "ds", "", `specify the datasource name to be created model and service, required flag(s) "mod"`)
	ccmd.Flags().StringVar(&root.Table, "table", "", `specify the table name to be created model and service, required flag(s) "mod"`)
	ccmd.Flags().StringVar(&root.Job, "job", "", `specify the job name to be created schedule, required flag(s) "sch"`)
	test := ccmd.Flags()
	test.BoolVar(&root.Test, "test", false, "")
	if err := test.MarkHidden("test"); err != nil {

		fmt.Println(err)
		return
	}
	ccmd.SetArgs(args)
	if err := ccmd.Execute(); err != nil {

		fmt.Println(err)
		if os.Getenv("github-action") == "Y" {

			os.Exit(0)
		} else {

			os.Exit(1)
		}
	}
}

func execute() error {

	if !root.Project && !root.Model && !root.Schedule {

		return errors.New(`required flag(s) "pro", "mod" or "sch" not set`)
	}
	if !root.Model && (root.Pointer || root.Gorm || root.Service || len(root.Datasource) > 0 || len(root.Table) > 0) {

		return errors.New(`required flag(s) "mod" not set`)
	}
	if !root.Schedule && len(root.Job) > 0 {

		return errors.New(`required flag(s) "sch" not set`)
	}
	if err := checkConfig(); err != nil {

		return err
	}
	if err := checkModule(); err != nil {

		return err
	}
	if err := checkProject(); err != nil {

		return err
	}
	if err := checkModel(); err != nil {

		return err
	}
	if err := checkService(); err != nil {

		return err
	}
	if err := checkJob(); err != nil {

		return err
	}
	return nil
}
