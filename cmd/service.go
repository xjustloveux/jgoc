// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"github.com/xjustloveux/jgoc/model"
	"sort"
	"strings"
)

func checkService() error {

	if !root.Service {

		return nil
	}
	/*
		service
		|---dsName
			|---tableName.go
	*/
	path := "service"
	if err := checkExistAndCreateFolder(path); err != nil {

		return err
	}
	path = fmt.Sprint(path, "/")
	if err := createServiceDS(path); err != nil {

		return err
	}
	return nil
}

func createServiceDS(path string) error {

	for _, v := range datasource {

		dsModName := fmt.Sprint(strToL(v.Name), "mod")
		dsSrvName := fmt.Sprint(strToL(v.Name), "srv")
		dsPath := fmt.Sprint(path, dsSrvName)
		if err := checkExistAndCreateFolder(dsPath); err != nil {

			return err
		}
		for _, t := range v.Tables {

			create := len(root.Table) <= 0 || t.Name == root.Table
			if create {

				if err := createServiceTable(fmt.Sprint(dsPath, "/"), dsModName, dsSrvName, v.Name, t); err != nil {

					return err
				}
			}
		}
	}
	return nil
}

func createServiceTable(path, dsModName, dsSrvName, dsName string, table model.Table) error {

	fileName := fmt.Sprint(strToL(table.Name), ".go")
	filePath := fmt.Sprint(path, fileName)
	jPrint(fmt.Sprint("create ", filePath, "..."))
	if err := createFile(filePath, getServiceTableContent(dsModName, dsSrvName, dsName, table)); err != nil {

		return err
	}
	return nil
}

func getServiceTableContent(dsModName, dsSrvName, dsName string, table model.Table) string {

	if root.Gorm {

		return getGormContent(dsModName, dsSrvName, dsName, table)
	}
	return getJgoContent(dsModName, dsSrvName, dsName, table)
}

func getJgoContent(dsModName, dsSrvName, dsName string, table model.Table) string {

	tableNameUCC := strToUCC(table.Name)
	tableNameLCC := strToLCC(table.Name)
	tableComment := table.Schema[0].TableComment
	model := fmt.Sprint(dsModName, ".", tableNameUCC)
	module := []string{`
	"` + root.Name + `/model/` + dsModName + `"`, `
	"github.com/fatih/structs"`, `
	"github.com/xjustloveux/jgo/jsql"`}
	sort.Slice(module, func(i, j int) bool {
		re := []string{" ", "　", "	", "_", `
`}
		a := module[i]
		b := module[j]
		for _, v := range re {
			a = strings.Replace(a, v, "", -1)
			b = strings.Replace(b, v, "", -1)
		}
		return a < b
	})
	mod := ""
	for _, v := range module {
		mod += v
	}
	return `package ` + dsSrvName + `

import (` + mod + `
)

// ` + tableNameUCC + ` ` + tableComment + `
var ` + tableNameUCC + ` = &` + tableNameLCC + `{
	ds:    "` + dsName + `",
	table: "` + table.Name + `",
}

type ` + tableNameLCC + ` struct {
	ds    string
	table string
}

// Create insert data
func (srv *` + tableNameLCC + `) Create(data ` + model + `) (jsql.Result, error) {

	return (&jsql.TableAgent{
		DSKey: srv.ds,
		Table: srv.table,
		Col:   structs.Map(data),
	}).Insert()
}

// CreateTx insert data for tx
func (srv *` + tableNameLCC + `) CreateTx(agent *jsql.Agent, data ` + model + `) (jsql.Result, error) {

	return (&jsql.TableAgent{
		Agent: agent,
		DSKey: srv.ds,
		Table: srv.table,
		Col:   structs.Map(data),
	}).Insert()
}

// FindAll query all data
func (srv *` + tableNameLCC + `) FindAll(param ...*jsql.Param) (jsql.Result, error) {

	return (&jsql.TableAgent{
		DSKey:  srv.ds,
		Table:  srv.table,
		Params: param,
	}).Query()
}

// FindAllTx query all data for tx
func (srv *` + tableNameLCC + `) FindAllTx(agent *jsql.Agent, param ...*jsql.Param) (jsql.Result, error) {

	return (&jsql.TableAgent{
		Agent:  agent,
		DSKey:  srv.ds,
		Table:  srv.table,
		Params: param,
	}).QueryTx()
}

// FindFirst query first data
func (srv *` + tableNameLCC + `) FindFirst(param ...*jsql.Param) (jsql.Result, error) {

	return (&jsql.TableAgent{
		DSKey:  srv.ds,
		Table:  srv.table,
		Params: param,
	}).QueryRow()
}

// FindFirstTx query first data for tx
func (srv *` + tableNameLCC + `) FindFirstTx(agent *jsql.Agent, param ...*jsql.Param) (jsql.Result, error) {

	return (&jsql.TableAgent{
		Agent:  agent,
		DSKey:  srv.ds,
		Table:  srv.table,
		Params: param,
	}).QueryRowTx()
}

// Update update data
func (srv *` + tableNameLCC + `) Update(data map[string]interface{}, param ...*jsql.Param) (jsql.Result, error) {

	return (&jsql.TableAgent{
		DSKey:  srv.ds,
		Table:  srv.table,
		Col:    data,
		Params: param,
	}).Update()
}

// UpdateTx update data for tx
func (srv *` + tableNameLCC + `) UpdateTx(agent *jsql.Agent, data map[string]interface{}, param ...*jsql.Param) (jsql.Result, error) {

	return (&jsql.TableAgent{
		Agent:  agent,
		DSKey:  srv.ds,
		Table:  srv.table,
		Col:    data,
		Params: param,
	}).UpdateTx()
}

// Delete delete data
func (srv *` + tableNameLCC + `) Delete(param ...*jsql.Param) (jsql.Result, error) {

	return (&jsql.TableAgent{
		DSKey:  srv.ds,
		Table:  srv.table,
		Params: param,
	}).Delete()
}

// DeleteTx delete data for tx
func (srv *` + tableNameLCC + `) DeleteTx(agent *jsql.Agent, param ...*jsql.Param) (jsql.Result, error) {

	return (&jsql.TableAgent{
		Agent:  agent,
		DSKey:  srv.ds,
		Table:  srv.table,
		Params: param,
	}).DeleteTx()
}
`
}

func getGormContent(dsModName, dsSrvName, dsName string, table model.Table) string {

	tableNameUCC := strToUCC(table.Name)
	tableNameLCC := strToLCC(table.Name)
	tableComment := table.Schema[0].TableComment
	model := fmt.Sprint(dsModName, ".", tableNameUCC)
	module := []string{`
	"` + root.Name + `/model/` + dsModName + `"`, `
	"fmt"`, `
	"gorm.io/gorm"`}
	sort.Slice(module, func(i, j int) bool {
		re := []string{" ", "　", "	", "_", `
`}
		a := module[i]
		b := module[j]
		for _, v := range re {
			a = strings.Replace(a, v, "", -1)
			b = strings.Replace(b, v, "", -1)
		}
		return a < b
	})
	mod := ""
	for _, v := range module {
		mod += v
	}
	return `package ` + dsSrvName + `

import (` + mod + `
)

// ` + tableNameUCC + ` ` + tableComment + `
var ` + tableNameUCC + ` = &` + tableNameLCC + `{
	table: "` + table.Name + `",
}

type ` + tableNameLCC + ` struct {
	table string
}

// Create insert data
func (srv *` + tableNameLCC + `) Create(db *gorm.DB, data ` + model + `) error {

	return db.Table(srv.table).Create(&data).Error
}

// FindAll query all data
func (srv *` + tableNameLCC + `) FindAll(db *gorm.DB, data []` + model + `, param ...map[string]interface{}) error {

	db = db.Table(srv.table)
	for _, v1 := range param {

		for k2, v2 := range v1 {

			db = db.Where(fmt.Sprint(k2, " = ?"), v2)
		}
	}
	return db.Find(&data).Error
}

// FindFirst query first data
func (srv *` + tableNameLCC + `) FindFirst(db *gorm.DB, data ` + model + `, param ...map[string]interface{}) error {

	db = db.Table(srv.table)
	for _, v1 := range param {

		for k2, v2 := range v1 {

			db = db.Where(fmt.Sprint(k2, " = ?"), v2)
		}
	}
	return db.First(&data).Error
}

// Update update data
func (srv *` + tableNameLCC + `) Update(db *gorm.DB, data map[string]interface{}, param ...map[string]interface{}) error {

	db = db.Table(srv.table).Where("1 = 1")
	for _, v1 := range param {

		for k2, v2 := range v1 {

			db = db.Where(fmt.Sprint(k2, " = ?"), v2)
		}
	}
	return db.Updates(&data).Error
}

// Delete delete data
func (srv *` + tableNameLCC + `) Delete(db *gorm.DB, param ...map[string]interface{}) error {

	db = db.Unscoped().Table(srv.table).Where("1 = 1")
	for _, v1 := range param {

		for k2, v2 := range v1 {

			db = db.Where(fmt.Sprint(k2, " = ?"), v2)
		}
	}
	return db.Delete(&` + model + `{}).Error
}
`
}
