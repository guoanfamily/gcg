# gcg
go 代码生成工具 基于sqlx echo mysql

go code generater based on sqlx echo by mysql
# install

go get github.com/guoanfamily/gcg

# use

1. create a folder in go src folder named with your projectname
```sh
mkdir projectname
cd projectname
```
2. create a config.yaml file and edit it like this:
```yaml
- Name: service
  Database: mysql
  Username: root
  Password: ""
  Host: localhost
  Port: 3306
  Table: "company_info,ebook"
  ServicePort: 9091
```
3.init the project use command in your project folder
```sh
gcg init [-config=config.yaml]
```

4.the file struct like this
```folder
projectname
--controller
----controller.go
----gen_table.go
--model
----service
------gen_table.go
------gen_db.go
--router
----gen_table.go
----router.go
--main.go
```
5.run the project
```sh
go run main.go
```

6.additional command

if you want add table in your exist project,just run ,split tablename with ','
```sh
gcg add [-tables=tablename1,tablename2]
```
and then add this code in your router.go file
```golang
{{tablename}}Add(e)
//start server
```
