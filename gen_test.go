package main

import (
	 "github.com/guoanfamily/gcg/gens/common"
	"testing"
	"github.com/guoanfamily/gcg/gens/orm"
	//"gcg/model/service"
	"os"
	"strings"
	"fmt"
)
//go:generate go-bindata -o ./bindata.go -prefix "./template/" ./template/...
func TestGen(t *testing.T) {
	genutils.Asset = Asset
	genutils.SetValues(map[string]interface{}{
		"RootPath": "./",
	})
	orm.Gen("config.yaml","init","")
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

func TestProjectFolderName(t *testing.T)  {
	s,_:=os.Getwd()
	folders := strings.Split(s,"\\")
	rs:= folders[len(folders)-1]
	fmt.Println(rs)
}
