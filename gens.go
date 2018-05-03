package main

import (
	"github.com/guoanfamily/gcg/gens/common"
	"github.com/guoanfamily/gcg/gens/orm"
)



func genORM(name, desc string) {
	genutils.SetValues(map[string]interface{}{
		"AppName":  appName,
		"AppPort":  appPort,
		"RootPath": rootPath,
	})
	orm.Gen(ormFile,name,tables)
}
