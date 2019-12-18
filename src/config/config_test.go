package config

import (
	"fmt"
	"testing"
)

func TestGetMysql(t *testing.T) {
	t.Run("InitConfig", func(t *testing.T) {
		InitConfig("../config.ini")
	})
	m := GetMysql()
	fmt.Println(m)
}

func TestInitConfig(t *testing.T) {
	testing.Init()
}
