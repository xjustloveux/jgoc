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
		Use:           name,
		Short:         CcmdShort,
		Long:          CcmdLong,
		Version:       CcmdVer,
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {

			return execute()
		},
	}
	ccmd.Flags().StringVar(&root.Name, "name", "", FlagsName)
	if err := ccmd.MarkFlagRequired("name"); err != nil {

		fmt.Println(err)
		return
	}
	ccmd.Flags().StringVar(&root.Env, "env", "", FlagsEnv)
	ccmd.Flags().BoolVar(&root.Project, "pro", false, FlagsPro)
	ccmd.Flags().BoolVar(&root.Model, "mod", false, FlagsMod)
	ccmd.Flags().BoolVar(&root.Schedule, "sch", false, FlagsSch)
	ccmd.Flags().BoolVar(&root.Pointer, "pointer", false, FlagsPointer)
	ccmd.Flags().BoolVar(&root.Gorm, "gorm", false, FlagsGorm)
	ccmd.Flags().BoolVar(&root.Service, "srv", false, FlagsSrv)
	ccmd.Flags().StringVar(&root.Datasource, "ds", "", FlagsDs)
	ccmd.Flags().StringVar(&root.Table, "table", "", FlagsTable)
	ccmd.Flags().StringVar(&root.Job, "job", "", FlagsJob)
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
	if err := checkYaml(); err != nil {

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
