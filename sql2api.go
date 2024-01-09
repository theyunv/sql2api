package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"sql2api/core"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

/*
Usage Example 1: Generate Single Api File  使用示例一：生成单文件Api
sql2api -syntax v1 -host localhost -port 3306 -package user -user root -password 123456 \
-schema testdatabase -service_name usersrv -field_style sqlApi -api_style all -table user,user_test -group true > usersrvdemo1.api

Usage Example 2: Generating Multi Api File  使用示例二：生成多文件Api
sql2api -syntax v1 -host localhost -port 3306 -package user -user root -password 123456 \
-schema testdatabase -service_name usersrv -field_style sqlApi -api_style message -table user -group true > user.api
sql2api -syntax v1 -host localhost -port 3306 -package user -user root -password 123456 \
-schema testdatabase -service_name usersrv -field_style sqlApi -api_style message -table user_test -group true > usertest.api
sql2api -syntax v1 -host localhost -port 3306 -package user -user root -password 123456 \
-schema testdatabase -service_name usersrv -field_style sqlApi -api_style server -table user_test,user_test -group true > usersrvdemo2.api
*/

func main() {
	dbType := flag.String("db", "mysql", "the database type")
	host := flag.String("host", "localhost", "the database host")
	port := flag.Int("port", 3306, "the database port")
	user := flag.String("user", "root", "the database user")
	password := flag.String("password", "root", "the database password")
	schema := flag.String("schema", "", "the database schema")
	table := flag.String("table", "*", "the table schema，multiple tables ',' split. ")
	serviceName := flag.String("service_name", *schema, "the api service name , defaults to the database schema.")
	packageName := flag.String("package", *schema, "the protocol buffer package. defaults to the database schema.")
	goPackageName := flag.String("syntax", "", "the api syntax version. defaults to the database schema.")
	ignoreTableStr := flag.String("ignore_tables", "", "a comma spaced list of tables to ignore")
	ignoreColumnStr := flag.String("ignore_columns", "", "a comma spaced list of mysql columns to ignore")
	fieldStyle := flag.String("field_style", "sqlApi", "gen api field style, sql_api | sqlApi")
	apiStyle := flag.String("api_style", "all", "gen api field style, all |  message | server")
	group := flag.Bool("group", false, "Whether service groups are supported, Note that the group bool must be followed after other type parameters")

	flag.Parse()

	if *schema == "" {
		fmt.Println(" - please input the database schema ")
		return
	}

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", *user, *password, *host, *port, *schema)
	db, err := sql.Open(*dbType, connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	ignoreTables := strings.Split(*ignoreTableStr, ",")
	ignoreColumns := strings.Split(*ignoreColumnStr, ",")

	s, err := core.GenerateSchema(db, *table, ignoreTables, ignoreColumns, *serviceName, *goPackageName, *packageName, *fieldStyle, *apiStyle, *group)

	if nil != err {
		log.Fatal(err)
	}

	if nil != s {
		fmt.Println(s)
	}
}
