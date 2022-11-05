[![JGo Web](https://jgo.dev/assets/images/logo_300.svg)](https://jgo.dev/)

[![JGoC release](https://img.shields.io/github/v/release/xjustloveux/jgoc)](https://github.com/xjustloveux/jgoc/releases)
[![codecov](https://codecov.io/gh/xjustloveux/jgoc/branch/master/graph/badge.svg?token=RCO5VO2YU6)](https://codecov.io/gh/xjustloveux/jgoc)
[![Build Status](https://github.com/xjustloveux/jgoc/actions/workflows/go.yml/badge.svg)](https://github.com/xjustloveux/jgoc/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/xjustloveux/jgoc)](https://goreportcard.com/report/github.com/xjustloveux/jgoc)
[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/xjustloveux/jgoc)](https://pkg.go.dev/mod/github.com/xjustloveux/jgoc)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/xjustloveux/jgoc/blob/master/LICENSE)


---

* [Overview](#Overview)
* [Middlewares](#Middlewares)
* [Installation](#Installation)
* [Quick Start](#Quick Start)
* [Command Line Usage](#Command Line Usage)

# Overview

---

JGoC provides an easier way to create Go project, model and schedule.

It is designed on the basis [JGo](https://github.com/xjustloveux/jgo) and [gin](https://github.com/gin-gonic/gin).

# Middlewares

---

Sql middleware use [mysql](https://github.com/go-sql-driver/mysql), [go-mssqldb](https://github.com/denisenkom/go-mssqldb)
and [godror](https://github.com/godror/godror).

***Note:* The middleware used by Oracle is `godror`, which is different from the default `go-oci8` of jgo, so the config file needs to set the `ds` tag.**

# Installation

---

```shell
go install github.com/xjustloveux/jgoc
```

# Quick Start

---

#### create project
```shell
jgoc --name example.com/helloworld --pro
```
#### create project, model
```shell
jgoc --name example.com/helloworld --pro --mod
```
#### create project, model, service
```shell
jgoc --name example.com/helloworld --pro --mod --srv
```
#### create project, model, service, schedule
```shell
jgoc --name example.com/helloworld --pro --mod --srv --sch
```
#### create model
```shell
jgoc --name example.com/helloworld --mod
```
#### create model, service
```shell
jgoc --name example.com/helloworld --pro --srv
```
#### create schedule
```shell
jgoc --name example.com/helloworld --sch
```

# Command Line Usage

---

| Flags         | Type   | Comment                                                                                                                                                              |
|---------------|--------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| --name        | string | module name, required                                                                                                                                                |
| --pro         |        | created project framework                                                                                                                                            |
| --mod         |        | created database model, need config/config.json or config/config.yaml file, configuration refer to [configuration](https://github.com/xjustloveux/jgo#configuration) |
| --pointer     |        | columns of numeric type will be converted to pointer type when creating the model, required flag(s) "mod"                                                            |
| --srv         |        | created model service, required flag(s) "mod"                                                                                                                        |
| --ds          | string | specify the datasource name to be created model and service, required flag(s) "mod"                                                                                  |
| --table       | string | specify the table name to be created model and service, required flag(s) "mod"                                                                                       |
| --gorm        |        | create model and service with gorm, model default only json tag, service default [jgo](https://github.com/xjustloveux/jgo), required flag(s) "mod"                   |
| --sch         |        | created schedule, need config/config.json or config/config.yaml file, configuration refer to [configuration](https://github.com/xjustloveux/jgo#configuration-1)     |
| --job         | string | specify the job name to be created schedule, required flag(s) "sch"                                                                                                  |
| --help, -h    |        | help for jgoc                                                                                                                                                        |
| --version, -v |        | version for jgoc                                                                                                                                                     |
