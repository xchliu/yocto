package storage

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"yocto/src/yoctoparser"
)

type table struct {
	name    string
	db      string
	charset string
	define  string
}

func DDL_Table(db string, name string, action string, define string) bool {
	var tb table
	tb.name = name
	tb.db = db
	tb.charset = "utf8"
	tb.define = define
	tb.name = strings.Replace(tb.name, "\n", "", -1)
	tb.name = strings.Replace(tb.name, " ", "", -1)

	if action == "create" {
		return Create_table(tb)
	}
	return false
}

func DDL_ColumnCreateTable(sqlObject yoctoparser.SQLObject) bool {
	fmt.Println("Create new table :" + sqlObject.TableList[0])
	datadir := "/tmp/yocto/data/"
	tabledir := filepath.Join(datadir, sqlObject.DB, sqlObject.TableList[0])
	fmt.Println("Init table in :" + tabledir)
	obj_def, err := os.OpenFile(tabledir+".def", os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return false
	}
	//todo do column action
	//tmp method
	for cIndex, col := range sqlObject.CreateColumns {
		_, _ = obj_def.WriteString(strconv.Itoa(cIndex) + ">" + col.Cname)
		_, _ = obj_def.WriteString(strconv.Itoa(cIndex) + ">" + col.Clength)
		_, _ = obj_def.WriteString(strconv.Itoa(cIndex) + ">" + col.Cprecision)
		_, _ = obj_def.WriteString(strconv.Itoa(cIndex) + ">" + strconv.Itoa(col.Datatype))
		for _, _ = range col.Constraint {
			//todo column constraint
			// column.constraint

		}
	}
	//obj_def.WriteString(sqlObject.CreateColumns)
	obj_data, err := os.OpenFile(tabledir+".ydb", os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return false
	}
	_, _ = obj_data.WriteString("")
	return true
}

func DML_Table(db string, name string, data string) bool {
	var ior IORequest
	ior.data = data
	ior.metadata = strings.Join([]string{db, name}, ".")
	ior.iotype = 1
	//Table level bottle on performance
	ior.key = string(get_next_id(ior.metadata))
	return ior.save()
}

func Create_table(tb table) bool {
	//TODO read from the config file
	fmt.Println("Create new table :" + tb.name)
	datadir := GetConf("datadir")
	tabledir := filepath.Join(datadir, tb.db, tb.name)
	fmt.Println("Init table in :" + tabledir)
	obj_def, err := os.OpenFile(tabledir+".def", os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return false
	}
	obj_def.WriteString(tb.define)
	obj_data, err := os.OpenFile(tabledir+".ydb", os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return false
	}
	obj_data.WriteString("")
	return true
}

func Get_Table_Define(db string, name string) {

}

func get_next_id(metadata string) int {
	v, ok := SEQUENCE[metadata]
	if ok {
		return v
	} else {
		SEQUENCE[metadata] = 1
		return 1
	}
}
