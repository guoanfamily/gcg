package orm

import (
	"database/sql"
	"fmt"
	sh "github.com/codeskyblue/go-sh"
	_ "github.com/go-sql-driver/mysql"
	"github.com/guoanfamily/gcg/gens/common"
	"github.com/guoanfamily/gcg/gens/funcs"
	"github.com/guoanfamily/gcg/gokits"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var typeMap = [][]string{
	{"int", "int32", "*int32"},
	{"tinyint", "byte", "*byte"},
	{"bigint", "int64", "*int64"},
	{"varchar", "string", "*string"},
	{"char", "string", "*string"},
	{"text", "string", "*string"},
	{"langtext", "string", "*string"},
	{"json", "string", "*string"},
	{"tinytext", "string", "*string"},
	{"datetime", "time.Time", "*time.Time"},
	{"date", "time.Time", "*time.Time"},
	{"timestamp", "time.Time", "*time.Time"},
	{"decimal", "float64", "*float64"},
	{"float", "float64", "*float64"},
	{"double", "float64", "*float64"},
}

// ColumnInfo table column info
type ColumnInfo struct {
	Field      string
	Type       string
	Null       string
	Key        string
	Default    *string
	Extra      string
	GoType     string
	ThriftType string
	Required   string
}

var isInit = false

// Gen gen
func Gen(dbFile string, name string, tbs string) {
	var databaseDir string
	if name == "init" {
		isInit = true
	}
	dirs := []string{
		"controller",
		"router",
	}
	model := map[string]interface{}{
		"RootPath": genutils.Values["RootPath"],
	}
	funcs.FuncMap["setDefault"] = SetDefault
	funcs.FuncMap["getTableFieldNames"] = GetTableFieldNames
	funcs.FuncMap["getTableFieldCounts"] = GetTableFieldCounts
	funcs.FuncMap["getObjColumn"] = GetObjColumn
	funcs.FuncMap["getUpdateColumn"] = GetUpdateColumn
	funcs.FuncMap["getInsertColumn"] = GetInsertColumn

	// mkdirs
	genutils.InitDirs(dirs)

	dbs := parseDBFile(dbFile)
	for _, v := range dbs {
		//生成service层代码
		db := v.(map[interface{}]interface{})
		genSource(db)
		databaseDir = "model/" + db["Name"].(string)
		genutils.MkDirIfNotExists(databaseDir)
		//生成表service代码
		loadDBMetaInfo(databaseDir, tbs, db)
		// log.Debugf("%v", db)
		//生成db连接代码
		if isInit {
			genutils.GenFileWithTargetPath("model/database/db.go.tmpl", databaseDir+"/gen_db.go", db)
		}
		sh.Command("gofmt", "-w", ".", sh.Dir(databaseDir)).Run()
		//生成controller层表相关代码
		for _, tableMeta := range db["TableMetas"].([]interface{}) {
			model := tableMeta.(map[interface{}]interface{})
			genutils.GenFileWithTargetPath("controller/gen_controller.go.tmpl", "controller/gen_"+model["TableName"].(string)+".go", tableMeta)
		}

		//生成controller层公共代码
		if isInit {
			genutils.GenFileWithTargetPath("controller/controller.go.tmpl", "controller/controller.go", nil)
		}
		sh.Command("gofmt", "-w", ".", sh.Dir("controller")).Run()
		//生成router层代码
		var tables = make(map[string]interface{})
		var sts []string
		for _, tableMeta := range db["TableMetas"].([]interface{}) {
			model := tableMeta.(map[interface{}]interface{})
			sts = append(sts, model["TableName"].(string))
			genutils.GenFileWithTargetPath("router/gen_router.go.tmpl", "router/gen_"+model["TableName"].(string)+".go", tableMeta)
		}
		sh.Command("gofmt", "-w", ".", sh.Dir("router")).Run()
		if isInit {
			tables["TableNames"] = sts
			tables["ServicePort"] = db["ServicePort"]
			genutils.GenFileWithTargetPath("router/router.go.tmpl", "router/router.go", tables)
			genutils.GenFileWithTargetPath("router/duplicateQuest.go.tmpl", "router/duplicateQuest.go", tables)
			sh.Command("gofmt", "-w", ".", sh.Dir("router")).Run()
			//生成main
			db["ProjectName"] = getProjectFolderName()
			genutils.GenFileWithTargetPath("main.go.tmpl", "main.go", db)
		}

	}
	model["DB"] = dbs
	// gen model/all/all.go
	//	log.Debugf("%v", model)
	//genutils.GenFile("model/all/all.go.tmpl", model)

	//sh.Command("gofmt", "-w", "model/all/all.go").Run()

}

// parseDBFile 解析出db配置文件信息
func parseDBFile(dbFile string) []interface{} {
	var bs []byte
	var err error
	var dbs interface{}
	if bs, err = ioutil.ReadFile(dbFile); err != nil {
		log.Fatalf("%s", err)
	}
	if err = yaml.Unmarshal(bs, &dbs); err != nil {
		log.Fatalf("%s", err)
	}
	return dbs.([]interface{})
}

// genSource 生成db连接字符串
func genSource(db map[interface{}]interface{}) {
	pwd := db["Password"].(string)
	if pwd != "" {
		pwd = ":" + pwd
	}
	port := 3306
	portObj, ok := db["Port"]
	if ok {
		port = portObj.(int)
	}
	source := fmt.Sprintf("%s%s@tcp(%s:%d)/%s?charset=utf8", db["Username"], pwd, db["Host"], port, db["Database"])
	db["Source"] = source
}

// loadDBMetaInfo 查询db元信息
func loadDBMetaInfo(databaseDir string, tables string, dbInfo map[interface{}]interface{}) {
	var (
		db         *sql.DB
		rows       *sql.Rows
		err        error
		tableMeta  interface{}
		tableMetas []interface{}
	)
	source := dbInfo["Source"].(string)
	if db, err = sql.Open("mysql", source); err != nil {
		log.Fatalf("%s", err)
	}
	defer db.Close()
	if rows, err = db.Query("SHOW TABLES"); err != nil {
		log.Fatalf("%s", err)
	}
	dbName := dbInfo["Name"].(string)
	//获取项目路径
	projectName := getProjectFolderName() //dbInfo["ProjectName"].(string)
	if tables == "" && isInit {
		tables = dbInfo["Table"].(string)
	}
	genTablesArray := strings.Split(tables, ",")
	for rows.Next() {
		var rowName string
		if err = rows.Scan(&rowName); err != nil {
			log.Fatalf("%s", err)
		}
		// 只加载配置的表
		if tables == "*" || utils.IsInArray(genTablesArray, rowName) {
			tableMeta = loadTableMetaInfo(db, rowName, dbName, projectName)
			model := tableMeta.(map[interface{}]interface{})
			//service层代码生成
			genutils.GenFileWithTargetPath("model/database/table.go.tmpl", databaseDir+"/gen_"+model["TableName"].(string)+".go", tableMeta)
			tableMetas = append(tableMetas, tableMeta)
		}
	}
	dbInfo["TableMetas"] = tableMetas
}

// laodTableMetaInfo 查询表结构元信息
func loadTableMetaInfo(db *sql.DB, tableName, dbName string, projectName string) interface{} {
	var (
		rows              *sql.Rows
		err               error
		primaryKey        = "id"
		primaryKeyType    = "int"
		primaryKeyDefault = "0"
		primaryKeyExtra   = ""
		autoIncrement     = false
		columnInfoList    []*ColumnInfo
	)
	if rows, err = db.Query("SHOW COLUMNS FROM `" + tableName + "`"); err != nil {
		log.Fatalf("%s", err)
	}

	for rows.Next() {
		c := new(ColumnInfo)
		if err = rows.Scan(&c.Field, &c.Type, &c.Null, &c.Key, &c.Default, &c.Extra); err != nil {
			log.Fatalf("%s", err)
		}
		c.GoType = toGoType(c.Type, c.Null)
		c.Required = "required"
		if c.Null == "YES" {
			c.Required = ""
		}
		// log.Debugf("%v", c)
		columnInfoList = append(columnInfoList, c)
		if c.Key == "PRI" {
			primaryKey = c.Field
			primaryKeyType = c.GoType
			primaryKeyExtra = c.Extra
			if primaryKeyType != "int32" {
				primaryKeyDefault = "\"\""
			}
			if c.Extra == "auto_increment" {
				autoIncrement = true
			}
		}
	}
	return map[interface{}]interface{}{
		"ProjectName":       projectName,
		"DBName":            dbName,
		"TableName":         tableName,
		"PrimaryKey":        primaryKey,
		"PrimaryKeyType":    primaryKeyType,
		"PrimaryKeyDefault": primaryKeyDefault,
		"PrimaryKeyExtra":   primaryKeyExtra,
		"AutoIncrement":     autoIncrement,
		"Columns":           columnInfoList,
	}
}

/*
转换数据库类型为go类型
*/
func toGoType(s, null string) string {
	for _, v := range typeMap {
		if strings.HasPrefix(s, v[0]) {
			if null == "YES" {
				return v[2]
			}
			return v[1]
		}
	}
	log.Fatalf("unsupport type %s", s)
	return ""
}

func SetDefault(ci *ColumnInfo) string {
	if ci.Default == nil {
		return ""
	}
	if *ci.Default == "CURRENT_TIMESTAMP" {
		return ""
	}
	if strings.Contains(ci.GoType, "string") && *ci.Default == "" {
		return ""
	}
	if (strings.Contains(ci.GoType, "int") || strings.Contains(ci.GoType, "byte") || strings.Contains(ci.GoType, "float")) && !strings.Contains(ci.GoType, "*") && *ci.Default == "0" {
		return ""
	}
	r := funcs.UpperCamel(ci.Field) + ": "
	if ci.Null == "YES" {
		if strings.Contains(ci.GoType, "Time") {
			r += "utils.ToTime(\"2006-01-02 15:04:05\", \""
		} else {
			r += "utils." + funcs.CapitalizeFirst(ci.GoType[1:]) + "("
		}

	}
	if strings.Contains(ci.GoType, "string") {
		r += "\""
	}
	r += *ci.Default
	if strings.Contains(ci.GoType, "string") {
		r += "\""
	}
	if ci.Null == "YES" {
		if strings.Contains(ci.GoType, "Time") {
			r += "\""
		}
		r += ")"
	}
	r += ","
	return r
}

func GetTableFieldNames(args []*ColumnInfo) string {
	names := []string{}
	for _, a := range args {
		//if a.Field == "id"{
		//	continue
		//}
		names = append(names, fmt.Sprintf("`%s`", a.Field))
	}
	return strings.Join(names, ", ")
}

func GetTableFieldCounts(args []*ColumnInfo) string {
	names := []string{}
	for _, a := range args {
		if a.Field == "id" || a.Field == "updatetime" || a.Field == "deleted" {
			continue
		}
		if a.Field == "createtime" {
			names = append(names, "now()")
		} else {
			names = append(names, "?")
		}
	}
	return strings.Join(names, ", ")
}

func GetObjColumn(args []*ColumnInfo) string {
	names := []string{}
	for _, a := range args {
		if a.Field == "id" || a.Field == "createtime" || a.Field == "updatetime" || a.Field == "deleted" {
			continue
		}
		names = append(names, fmt.Sprintf("obj.%s", strFirstToUpper(a.Field)))
	}
	return strings.Join(names, ", ")
}

func GetUpdateColumn(args []*ColumnInfo) string {
	names := []string{}
	for _, a := range args {
		if a.Field == "id" || a.Field == "createtime" || a.Field == "updatetime" || a.Field == "deleted" {
			continue
		}
		names = append(names, fmt.Sprintf("%s=?", a.Field))
	}
	return strings.Join(names, ", ")
}

func GetInsertColumn(args []*ColumnInfo) string {
	names := []string{}
	for _, a := range args {
		if a.Field == "id" || a.Field == "updatetime" || a.Field == "deleted" {
			continue
		}
		names = append(names, fmt.Sprintf("`%s`", a.Field))
	}
	return strings.Join(names, ", ")
}

/**
 * 字符串首字母转化为大写 ios_bbbbbbbb -> IosBbbbbbbbb
 */
func strFirstToUpper(str string) string {
	temp := strings.Split(str, "_")
	var upperStr string
	for y := 0; y < len(temp); y++ {
		vv := []rune(temp[y])
		for i := 0; i < len(vv); i++ {
			if i == 0 {
				vv[i] -= 32
				upperStr += string(vv[i]) // + string(vv[i+1])
			} else {
				upperStr += string(vv[i])
			}
		}
	}
	return upperStr
}

/**
 * 获取项目文件夹名称
 */
func getProjectFolderName() string {
	s, _ := os.Getwd()
	folders := strings.Split(s, "\\")
	return folders[len(folders)-1]
}
