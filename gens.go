package main

import (
	"github.com/domego/gopt/gens/common"
	"gcg/gens/orm"
)



func genORM(name, desc string) {
	genutils.SetValues(map[string]interface{}{
		"AppName":  appName,
		"AppPort":  appPort,
		"RootPath": rootPath,
	})
	orm.Gen(ormFile,name,tables)
}
