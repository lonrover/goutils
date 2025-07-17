package databaseconfig

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
)

type MySQLDB struct {
	db *sql.DB
}

type MysqlConfig struct {
	Username string
	Password string
	Port     string
	Address  string
	Database string
}

/**
 * @description: 创建一个新的 MySQLDB 实例，并初始化连接池
 * @param {MysqlConfig} config 【数据库的配置信息】
 * @param {*} maxOpenConns 【最大支持的连接数】
 * @param {int} maxIdleConns 【最大空闲的连接数】
 * @return {*}
 */
func NewMySQLDB(config MysqlConfig, maxOpenConns, maxIdleConns int) (*MySQLDB, error) {
	// 使用 mysql.Config 结构体构建 DSN，更安全且支持更多配置选项
	cfg := mysql.Config{
		User:                 config.Username,
		Passwd:               config.Password,
		Net:                  "tcp",
		Addr:                 config.Address + ":" + config.Port,
		DBName:               config.Database,
		Timeout:              5 * time.Second, // 连接超时
		ReadTimeout:          3 * time.Second, // 读取超时
		WriteTimeout:         3 * time.Second, // 写入超时
		AllowNativePasswords: true,            // 允许本地密码认证
		ParseTime:            true,            // 自动解析时间类型
		Loc:                  time.UTC,        // 设置时区为 UTC
	}

	// 使用配置生成 DSN 字符串
	dsn := cfg.FormatDSN()

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &MySQLDB{db: db}, nil
}

// Close 关闭数据库连接
func (db *MySQLDB) Close() error {
	return db.db.Close()
}

// Insert 插入数据
func (db *MySQLDB) Insert(query string, args ...interface{}) (sql.Result, error) {
	stmt, err := db.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.Exec(args...)
}

// Read 查询数据
func (db *MySQLDB) Read(query string, args ...interface{}) (*sql.Rows, error) {
	return db.db.Query(query, args...)
}

// Update 更新数据
func (db *MySQLDB) Update(query string, args ...interface{}) (sql.Result, error) {
	stmt, err := db.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.Exec(args...)
}

// Delete 删除数据
func (db *MySQLDB) Delete(query string, args ...interface{}) (sql.Result, error) {
	stmt, err := db.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.Exec(args...)
}

/** 添加一个安全的查询方法，自动处理 rows.Close()
 * @description: 查询所有数据并返回result结果集
 * @param {string} query => 查询sql
 * @param {...interface{}} args
 * @return {*}
 */
func (db *MySQLDB) FetchAll(query string, args ...interface{}) ([]map[string]interface{}, error) {
	rows, err := db.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // 确保 rows 被关闭

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))

	for rows.Next() {
		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, err
		}

		row := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			row[col] = v
		}
		result = append(result, row)
	}

	return result, nil
}
