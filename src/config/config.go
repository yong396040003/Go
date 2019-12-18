package config

import (
	"github.com/Unknwon/goconfig"
	"log"
	model2 "model"
)

var Config *goconfig.ConfigFile

func init() {
}

/**
加载配置文件
*/
func InitConfig(path string) {
	cfg, err := goconfig.LoadConfigFile(path)
	if err != nil {
		log.Fatalf("读取文件失败,失败路径" + path)
		return
	}
	Config = cfg
}

/**
获取port
*/
func GetKV(section string, key string) string {
	value, err := Config.GetValue(section, key)
	if err != nil {
		log.Fatalf("无法读取键值%s:(%s)", value, err)
	}
	return value
}

/**
获取我的数据库信息

*/
func GetMyBam(section string) model2.DB {
	var db model2.DB
	sec, err := Config.GetSection(section)
	if err != nil {
		log.Fatalf("获取sections失败,%s", err)
	}
	db.Section = section
	db.Name = sec["name"]
	db.Password = sec["password"]
	db.IP = sec["ip"]
	db.Database = sec["database"]
	return db
}
