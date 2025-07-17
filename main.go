package main

import (
	"fmt"
	"goutils/databaseconfig"
	"goutils/logger"
	"time"
)

func main() {
	fmt.Println("this is a utils class with golang program.")

	// 初始化数据库配置
	var config_db = databaseconfig.MysqlConfig{
		Username: "senselink",
		Password: "senselink_2018_local",
		Port:     "3306",
		Address:  "192.168.220.208",
		Database: "bi_slink_base",
	}
	maxOpenConns, maxIdleConns := 20, 10

	// 初始化 Log
	filePath, filename := "./", "package_common"+time.Now().Format("2006-01-02")+".json"
	Log := logger.InitLog(filePath, filename)
	mysql_db, err := databaseconfig.NewMySQLDB(config_db, maxOpenConns, maxIdleConns)

	if err != nil {
		Log.Info("初始化数据库错误： %s", err)
	}

	sql := "select * from t_user"
	res, err := mysql_db.FetchAll(sql)

	if err != nil {
		Log.Info("初始化数据库错误： %s", err)
	}

	fmt.Println(res[0])
	// for index, item := range res {
	// 	fmt.Println(index, item)
	// }

}
