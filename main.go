package main

import (
	"github.com/lonrover/goutils/config"
	"github.com/lonrover/goutils/databaseconfig"
	"github.com/lonrover/goutils/logger"
)

func main() {
	// 初始化配置信息
	config.InitConfig("setting", "yaml", "D:/code/work/package_common_go/config")
	Golbalconfig := config.GlobalConfig

	// 初始化日志信息
	logfilepath, logfilename := "D:/code/work/package_common_go/logs/", "etl_display.log"
	logger.Init(logfilepath, logfilename)
	log := logger.GetLogger()

	log.Info("mysql 配置信息为： ", Golbalconfig.Mysql_db_config)
	_, err := databaseconfig.NewMySQLDB(Golbalconfig.Mysql_db_config, 10, 5)

	if err != nil {
		log.Error("初始化数据库出错：", err)
	}
}
