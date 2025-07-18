/*
 * @Author: your name
 * @Date: 2025-07-16 16:51:54
 * @LastEditTime: 2025-07-18 14:47:59
 * @LastEditors: your name
 * @Description:
 * @FilePath: \package_common_go\config\configViper.go
 * 可以输入预定的版权声明、个性签名、空行等
 */
package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Name string `mapstructure:"name"`
	Env  string `mapstructure:"env"`
}

type MysqlConfig struct {
	Host          string `mapstructure:"host"`
	Port          string `mapstructure:"port"`
	Username      string `mapstructure:"username"`
	Password      string `mapstructure:"password"`
	Database_name string `mapstructure:"database_name"`
}

type LogConfig struct {
	Path       string `mapstructure:"path"`
	Level      string `mapstructure:"level"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type Config struct {
	App             AppConfig   `mapstructure:"app"`
	Mysql_db_config MysqlConfig `mapstructure:"mysql_db_config"`
	Log             LogConfig   `mapstructure:"log"`
}

var GlobalConfig Config

// 初始化配置读取
/** 该部分只定义了部分可用的配置信息，后续可以根据代码做出修改。
 * @description: 初始化配置信息
 * @param {*} setName => 配置文件名
 * @param {*} setType => 配置文件类型
 * @param {string} setPath => 配置文件路径
 * @return {*} 只做初始化，没有返回值。
 */
func InitConfig(setName, setType, setPath string) {
	viper.SetConfigName(setName) // 配置文件名 (不带扩展名)
	viper.SetConfigType(setType) // 配置文件类型
	viper.AddConfigPath(setPath) // 当前文件所在目录WW

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// 解析到结构体
	if err := viper.Unmarshal(&GlobalConfig); err != nil {
		log.Fatalf("Unable to decode config into struct: %v", err)
	}

	log.Println("Config loaded successfully")
}

// 获取数据库连接字符串
func (d *MysqlConfig) ConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		d.Host, d.Port, d.Username, d.Password, d.Database_name)
}

// 使用示例
func ExampleUsage() {
	// 获取日志路径
	logPath := GlobalConfig.Log.Path
	fmt.Println("Log path:", logPath)

	// 获取数据库连接字符串
	dbConn := GlobalConfig.Mysql_db_config.ConnectionString()
	fmt.Println("DB Connection:", dbConn)

	fmt.Println("GlobalConfig: ", GlobalConfig.Log.Level)

}
