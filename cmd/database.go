// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"github.com/xjustloveux/jgo/jsql"
	"github.com/xjustloveux/jgoc/model"
)

func loadSchema() error {

	if len(root.Datasource) > 0 {

		datasource = make([]*model.Datasource, 1)
		datasource[0] = &model.Datasource{
			Name: root.Datasource,
		}
	} else {

		datasource = make([]*model.Datasource, 0)
		ds := jsql.GetDataSource()
		for k := range ds {

			newDs := &model.Datasource{
				Name:   k,
				Tables: nil,
			}
			datasource = append(datasource, newDs)
		}
	}
	for i, v := range datasource {

		if root.Test {

			var k int8
			k = 1
			datasource[i].Tables = []model.Table{
				{
					Name: "TEST",
					Schema: []jsql.TableSchema{
						{
							ColumnName:    "SEQ",
							DataType:      "int",
							IsNullable:    "NO",
							DataDefault:   "",
							PrimaryKey:    &k,
							IsIdentity:    "YES",
							ColumnComment: "sequence",
							TableComment:  "JGoC Test Table",
						},
					},
				},
			}
			continue
		}
		if agent, err := jsql.GetAgent(v.Name); err != nil {

			return err
		} else {

			if datasource[i], err = getTableSchema(agent, v); err != nil {

				return err
			}
		}
	}
	return nil
}

func getTableSchema(agent *jsql.Agent, ds *model.Datasource) (*model.Datasource, error) {

	if err := agent.UseTx(func() error {

		jPrint("load tables...")
		if list, err := agent.TablesTx(); err != nil {

			return err
		} else {

			ds.Tables = make([]model.Table, len(list))
			for i, v := range list {

				ds.Tables[i] = model.Table{
					Name:   v,
					Schema: nil,
				}
				jPrint(fmt.Sprint("load ", v, " table schema..."))
				if ds.Tables[i].Schema, err = agent.TableSchemaTx(v); err != nil {

					return err
				}
			}
		}

		return nil
	}); err != nil {

		return nil, err
	}
	return ds, nil
}
