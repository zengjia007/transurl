package dbutil

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"transurl/util"
)

var Db *sql.DB


// 初始化数据库连接
// connectionInfo为数据库连接的信息
func InitDB(connectionInfo string) *sql.DB {
	DB, err := sql.Open("mysql", connectionInfo)

	if err != nil {
		panic(fmt.Sprintf("数据库连接失败，错误信息为:", err.Error()))
	}
	return DB
}

// 获取数据库连接
func GetConn() *sql.DB {
	conf := util.GetInstanceConf()
	userName := conf.Read("MySql", "USER_NAME")
	password := conf.Read("MySql", "PASSWORD")
	hostName := conf.Read("MySql", "DATABASE_HOST")
	databaseName := conf.Read("MySql", "DATABASE_NAME")
	port := conf.Read("MySql", "PORT")
	queryString := conf.Read("MySql", "QUERY_STRING")
	return InitDB(userName + ":" + password + "@tcp(" + hostName + ":" + port + ")/" + databaseName + "?" + queryString)
}