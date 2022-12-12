// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package cmd

const (
	CcmdShort = `create project, model or schedule as fast and easily as possible`
	CcmdLong  = `JGoC provides an easier way to create Go project, model and schedule.
JGoC goal is to simplify the created project framework, model and schedule steps while providing variant customization options for all steps.
Check out github for more information: https://github.com/xjustloveux/jgof`
	CcmdVer = `v1.0.21`

	FlagsName    = `module name`
	FlagsEnv     = `jgo config environment value`
	FlagsPro     = `created project framework`
	FlagsMod     = `created database model, need config/config.json or config/config.yaml file, configuration refer to https://github.com/xjustloveux/jgo#configuration`
	FlagsSch     = `created schedule, need config/config.json or config/config.yaml file, configuration refer to https://github.com/xjustloveux/jgo#configuration-1`
	FlagsPointer = `columns of numeric type will be converted to pointer type when creating the model`
	FlagsGorm    = `create model and service with gorm, required flag(s) "mod"`
	FlagsSrv     = `created model service, required flag(s) "mod"`
	FlagsDs      = `specify the datasource name to be created model and service, required flag(s) "mod"`
	FlagsTable   = `specify the table name to be created model and service, required flag(s) "mod"`
	FlagsJob     = `specify the job name to be created schedule, required flag(s) "sch"`

	GoMod = `go.mod`

	ModuleGin        = `github.com/gin-gonic/gin`
	ModuleYaml       = `gopkg.in/yaml.v2`
	ModuleMySql      = `github.com/go-sql-driver/mysql`
	ModuleMSSql      = `github.com/denisenkom/go-mssqldb`
	ModuleOracle     = `github.com/godror/godror`
	ModulePostgreSql = `github.com/lib/pq`
	ModuleGorm       = `gorm.io/gorm`
	ModuleJGo        = `github.com/xjustloveux/jgo`
	ModuleStructs    = `github.com/fatih/structs`
	ModuleGovaluate  = `github.com/Knetic/govaluate`

	ProCtrA101Import1 = `
	"%v/service/a/a101srv"`
	ProCtrA101Import2 = `
	"%v"`
	ProCtrA101Import3 = `
	"net/http"`
	ProCtrA101 = `package a101ctr

import (%v
)

func Example(ctx *gin.Context) {

	message := a101srv.DoSomething()
	ctx.JSON(http.StatusOK, gin.H{"success": true, "message": message})
}
`
	ProCtrARouter = `package a

import (
	"%v/controller/a/a101ctr"
	"%v/global"
)

func Init() {

	a := global.Router.Group("a")
	{
		a101 := a.Group("a101")
		{
			a101.GET("", a101ctr.Example)
		}
	}
}
`
	ProCtrRouter = `package controller

import (
	"%v/controller/a"
	"%v/global"
)

func Init() error {

	a.Init()
	return global.Router.Run(":8080")
}
`
	ProVar = `package global

import "%v"

var (
	Router = gin.Default()
)
`
	ProSrvA101 = `package a101srv

func DoSomething() string {

	// Do Something
	return "Example"
}
`
	ProMainImport1 = `
	"%v/controller"`
	ProMainImport2 = `
	"fmt"`
	ProMainImportYaml1 = `
	"%v/middleware/yaml"`
	ProMainImportYaml2 = `
	"%v/jfile"`
	ProMainInitYaml = `
	jfile.RegisterCodec(jfile.Yaml.String(), yaml.Codec{})`
	ProMainImportJob1 = `
	"%v/job"`
	ProMainImportJob2 = `
	"%v/jcron"`
	ProMainInitJobYaml = `
	jcron.SetFormat(jfile.Yaml)
	jcron.SetFileName("config.yaml")`
	ProMainInitJob = `%v
	if err := jcron.Init(); err != nil {

		fmt.Println(err)
	}
	if err := job.Init(); err != nil {
		
		fmt.Println(err)
	}`
	ProMainImportJGoSql = `
	"%v/jsql"`
	ProMainInitSqlYaml = `
	jsql.SetFormat(jfile.Yaml)
	jsql.SetFileName("config.yaml")`
	ProMainInitSql = `%v
	if err := jsql.Init(); err != nil {

		fmt.Println(err)
	}`
	ProMainImportSql = `
	_ "%v"`
	ProMain = `package main

import (%v
)

func main() {
	%v%v%v
	if err := controller.Init(); err != nil {
		
		fmt.Println(err)
	}
}
`

	YamlCodec = `package yaml

type Codec struct{}

func (Codec) Encode(v map[string]interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

func (Codec) Decode(b []byte, v map[string]interface{}) error {
	return yaml.Unmarshal(b, &v)
}
`
	YamlYaml2 = `package yaml

import yamlv2 "%v"

var yaml = struct {
	Marshal   func(in interface{}) (out []byte, err error)
	Unmarshal func(in []byte, out interface{}) (err error)
}{
	Marshal:   yamlv2.Marshal,
	Unmarshal: yamlv2.Unmarshal,
}
`

	ModelImport = `

import "time"`
	Model = `package %v%v

// %v %v
type %v struct {%v
}`

	SrvImport1 = `
	"%v/model/%v"`
	SrvImport2 = `
	"%v"`
	SrvJGoImport = `
	"%v/jsql"`
	SrvJGo = `package %v

import (%v
)

// %UCC% %v
var %UCC% = &%LCC%{
	ds:    "%v",
	table: "%v",
}

type %LCC% struct {
	ds    string
	table string
}

// Ds return ds
func (srv *%LCC%) Ds() string {

	return srv.ds
}

// Table return table name
func (srv *%LCC%) Table() string {

	return srv.table
}

// Create insert data
func (srv *%LCC%) Create(data %model%) (jsql.Result, error) {

	return (&jsql.TableAgent{
		DSKey: srv.ds,
		Table: srv.table,
		Col:   structs.Map(data),
	}).Insert()
}

// CreateTx insert data for tx
func (srv *%LCC%) CreateTx(agent *jsql.Agent, data %model%) (jsql.Result, error) {

	return (&jsql.TableAgent{
		Agent: agent,
		DSKey: srv.ds,
		Table: srv.table,
		Col:   structs.Map(data),
	}).Insert()
}

// FindAll query all data
func (srv *%LCC%) FindAll(param ...*jsql.Param) (jsql.Result, error) {

	return (&jsql.TableAgent{
		DSKey:  srv.ds,
		Table:  srv.table,
		Params: param,
	}).Query()
}

// FindAllTx query all data for tx
func (srv *%LCC%) FindAllTx(agent *jsql.Agent, param ...*jsql.Param) (jsql.Result, error) {

	return (&jsql.TableAgent{
		Agent:  agent,
		DSKey:  srv.ds,
		Table:  srv.table,
		Params: param,
	}).QueryTx()
}

// FindFirst query first data
func (srv *%LCC%) FindFirst(param ...*jsql.Param) (jsql.Result, error) {

	return (&jsql.TableAgent{
		DSKey:  srv.ds,
		Table:  srv.table,
		Params: param,
	}).QueryRow()
}

// FindFirstTx query first data for tx
func (srv *%LCC%) FindFirstTx(agent *jsql.Agent, param ...*jsql.Param) (jsql.Result, error) {

	return (&jsql.TableAgent{
		Agent:  agent,
		DSKey:  srv.ds,
		Table:  srv.table,
		Params: param,
	}).QueryRowTx()
}

// Update update data
func (srv *%LCC%) Update(data map[string]interface{}, param ...*jsql.Param) (jsql.Result, error) {

	return (&jsql.TableAgent{
		DSKey:  srv.ds,
		Table:  srv.table,
		Col:    data,
		Params: param,
	}).Update()
}

// UpdateTx update data for tx
func (srv *%LCC%) UpdateTx(agent *jsql.Agent, data map[string]interface{}, param ...*jsql.Param) (jsql.Result, error) {

	return (&jsql.TableAgent{
		Agent:  agent,
		DSKey:  srv.ds,
		Table:  srv.table,
		Col:    data,
		Params: param,
	}).UpdateTx()
}

// Delete delete data
func (srv *%LCC%) Delete(param ...*jsql.Param) (jsql.Result, error) {

	return (&jsql.TableAgent{
		DSKey:  srv.ds,
		Table:  srv.table,
		Params: param,
	}).Delete()
}

// DeleteTx delete data for tx
func (srv *%LCC%) DeleteTx(agent *jsql.Agent, param ...*jsql.Param) (jsql.Result, error) {

	return (&jsql.TableAgent{
		Agent:  agent,
		DSKey:  srv.ds,
		Table:  srv.table,
		Params: param,
	}).DeleteTx()
}
`
	SrvGorm = `package %v

import (%v
)

// %UCC% %v
var %UCC% = &%LCC%{
	table: "%v",
}

type %LCC% struct {
	table string
}

// Table return table name
func (srv *%LCC%) Table() string {

	return srv.table
}

// Create insert data
func (srv *%LCC%) Create(db *gorm.DB, data *%model%) error {

	return db.Table(srv.table).Create(data).Error
}

// FindAll query all data
func (srv *%LCC%) FindAll(db *gorm.DB, data *[]%model%, param ...map[string]interface{}) error {

	db = db.Table(srv.table)
	for _, v1 := range param {

		for k2, v2 := range v1 {

			db = db.Where(fmt.Sprint(k2, " = ?"), v2)
		}
	}
	return db.Find(data).Error
}

// FindFirst query first data
func (srv *%LCC%) FindFirst(db *gorm.DB, data *%model%, param ...map[string]interface{}) error {

	db = db.Table(srv.table)
	for _, v1 := range param {

		for k2, v2 := range v1 {

			db = db.Where(fmt.Sprint(k2, " = ?"), v2)
		}
	}
	return db.First(data).Error
}

// Update update data
func (srv *%LCC%) Update(db *gorm.DB, data map[string]interface{}, param ...map[string]interface{}) error {

	db = db.Table(srv.table).Where("1 = 1")
	for _, v1 := range param {

		for k2, v2 := range v1 {

			db = db.Where(fmt.Sprint(k2, " = ?"), v2)
		}
	}
	return db.Updates(&data).Error
}

// Delete delete data
func (srv *%LCC%) Delete(db *gorm.DB, param ...map[string]interface{}) error {

	db = db.Unscoped().Table(srv.table).Where("1 = 1")
	for _, v1 := range param {

		for k2, v2 := range v1 {

			db = db.Where(fmt.Sprint(k2, " = ?"), v2)
		}
	}
	return db.Delete(&%model%{}).Error
}
`

	JobInit = `package job

import (
	"%v/jcron"
)

type job interface {
	Name() string
	Run(map[string]interface{})
}

func Init() error {

	jobs := []job{%v
	}
	for _, v := range jobs {

		if err := jcron.AddJob(v.Name(), v); err != nil {

			return err
		}
	}
	return nil
}
`
	Job = `package job

import "fmt"

type %n% struct{}

func (j *%n%) Name() string {

	return "%n%"
}

func (j *%n%) Run(map[string]interface{}) {

	fmt.Println("%n%: start")
	// Do Something
	fmt.Println("%n%: end")
}
`
)
