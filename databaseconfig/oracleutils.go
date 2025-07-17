package databaseconfig

import (
	"database/sql"
	"fmt"
)

// DBHandler 封装了 Oracle 数据库的连接和操作接口
type DBHandler struct {
	db *sql.DB
}


func OracleUtils(dsn string) *DBHandler {
	// 连接字符串格式：用户名/密码@主机地址:端口号/数据库实例名
	// dsn := "user/password@host:port/servicename"

	// 连接数据库
	db, err := sql.Open("godror", dsn)

	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return nil
	}
	return &DBHandler{db: db}
}

// 关闭数据库连接
func (h *DBHandler) Close() error {
	return h.db.Close()
}

// query 执行查询操作，并且返回结果集
func (h *DBHandler) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return h.db.Query(query, args...)
}

// Exec 执行非查询操作（插入、更新、删除等）
func (h *DBHandler) Exec(query string, args ...interface{}) (sql.Result, error) {
	return h.db.Exec(query, args...)
}
