package config

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Name string `mapstructure:"name"`
	Env  string `mapstructure:"env"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
}

type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
}

type LogConfig struct {
	Path       string `mapstructure:"path"`
	Level      string `mapstructure:"level"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Database DatabaseConfig `mapstructure:"database"`
	Redis    RedisConfig    `mapstructure:"redis"`
	Log      LogConfig      `mapstructure:"log"`
}

var GlobalConfig Config

// 初始化配置读取
func InitConfig() {
	// 获取当前文件路径（解决工作目录问题）
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)

	viper.SetConfigName("config")      // 配置文件名 (不带扩展名)
	viper.SetConfigType("yaml")        // 配置文件类型
	viper.AddConfigPath(basePath)      // 当前文件所在目录
	viper.AddConfigPath(".")           // 工作目录
	viper.AddConfigPath("/etc/myapp/") // 全局配置目录

	// 从环境变量读取覆盖配置
	viper.AutomaticEnv()
	viper.SetEnvPrefix("MYAPP") // 环境变量前缀 MYAPP_DATABASE_PASSWORD

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
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		d.Host, d.Port, d.Username, d.Password, d.Name)
}

// 使用示例
func ExampleUsage() {
	// 获取日志路径
	logPath := GlobalConfig.Log.Path
	fmt.Println("Log path:", logPath)

	// 获取数据库连接字符串
	dbConn := GlobalConfig.Database.ConnectionString()
	fmt.Println("DB Connection:", dbConn)

	// 获取Redis密码
	redisPass := GlobalConfig.Redis.Password
	fmt.Println("Redis Password:", redisPass)
}
