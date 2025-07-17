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

type DatabaseConfig struct {
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
	App             AppConfig      `mapstructure:"app"`
	Mysql_db_config DatabaseConfig `mapstructure:"mysql_db_config"`
	Log             LogConfig      `mapstructure:"log"`
}

var GlobalConfig Config

// 初始化配置读取
func InitConfig() {
	viper.SetConfigName("setting")                                // 配置文件名 (不带扩展名)
	viper.SetConfigType("yaml")                                   // 配置文件类型
	viper.AddConfigPath("D:/code/work/package_common_go/config/") // 当前文件所在目录WW

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
func (d *DatabaseConfig) ConnectionString() string {
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
