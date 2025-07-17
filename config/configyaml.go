package config

import (
	"goutils/logger"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

// YAMLConfigProvider 实现了 ConfigProvider 接口，从 YAML 文件中读取配置信息
type YAMLConfigProvider struct {
	SettingFilePath string
	SecretFilePath  string
}

/*
Config 结构体定义了配置信息

	Username string `yaml:"username"`
	Password string `yaml:"password"`

用于指定字段在序列化为 YAML 格式时的键名。
在将结构体序列化为 YAML 格式时，字段的名称将会被替换为标签中指定的键名。
这种方式可以在结构体字段的名称和实际的序列化格式之间建立映射关系，使得序列化和反序列化的过程更加灵活。
*/
type ConfigSetting struct {
	Database struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"database"`
}

type ConfigSecret struct {
	DB_ME_IOT struct {
		Username     string `yaml:"username"`
		Password     string `yaml:"password"`
		Host         string `yaml:"host"`
		Port         string `yaml:"port"`
		Service_name string `yaml:"service_name"`
	} `yaml:"db_me_iot"`
}

var yamlconfig = YAMLConfigProvider{
	SecretFilePath:  "./utils/secret.yaml",
	SettingFilePath: "./utils/setting.yaml",
}

// GetDatabaseConfig 从 YAML 文件中读取数据库配置信息
func GetSecretConfig() (ConfigSecret, error) {
	// 初始化config的日志
	// 记录 log 的文件名
	fileNmae := "log/config" + time.Now().Format("2006-01-02") + ".json"
	filepath := "./"
	log := logger.InitLog(filepath, fileNmae)

	file, err := os.Open(yamlconfig.SecretFilePath)
	if err != nil {
		log.Info("文件解析错误: %s", err)

	}

	// 解析 yaml 文件
	var config ConfigSecret

	decoder := yaml.NewDecoder(file)

	if err := decoder.Decode(&config); err != nil {
		log.Error("Failed to decode config file:", err)
	}

	// fmt.Println("yaml:", config.DB_ME_IOT)

	return config, err
}
