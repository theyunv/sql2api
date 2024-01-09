#### Give a star before you see it. xie xie xie ~ ~
Generates a api file  for the go-zero framework from your mysql database.  
通过mysql数据库为go-zero框架生成一个api文件
#### Tips
Welcome to modify and optimize the code  
欢迎自行修改和优化代码
#### Use from the command line:
`go install github.com/theyunv/sql2api@latest`


```
$ sql2api -h

Usage of ./sql2api:
  -db string
    	the database type (default "mysql")
  -host string
    	the database host (default "localhost")
  -ignore_columns string
    	a comma spaced list of mysql columns to ignore
  -ignore_tables string
    	a comma spaced list of tables to ignore
  -package string
    	the protocol buffer package. defaults to the database schema.
  -password string
    	the database password (default "root")
  -port int
    	the database port (default 3306)
  -schema string
    	the database schema
  -service_name string
    	the api service name , defaults to the database schema.
  -syntax string
    	the api syntax version. defaults to the database schema.
  -table string
    	the table schema，multiple tables ',' split.  (default "*")
  -user string
    	the database user (default "root")
  -api_style string
    	gen api field style, all |  message | server (default "all")
  -field_style string
    	gen api field style, sql_api | sqlApi (default "sqlApi")
  -group
    	Whether service groups are supported, Note that the group bool must be followed after other type parameters

```

```
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
```

