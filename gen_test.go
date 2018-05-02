package main

import (
	 "gcg/gens/common"
	"testing"
	"gcg/gens/orm"
	//"gcg/model/service"
)
//go:generate go-bindata -o ./bindata.go -prefix "./template/" ./template/...
func TestGen(t *testing.T) {
	genutils.Asset = Asset
	genutils.SetValues(map[string]interface{}{
		"RootPath": "./",
	})
	orm.Gen("config.yaml")
}

//func TestSave(t *testing.T)  {
//	var ci  service.CompanyInfo
//	var i int32 = 20
//	s := "test2"
//	ci.Id = 49
//	ci.CompanyName=&s
//	ci.PeopleNum = &i
//	ci.Save()
//}
