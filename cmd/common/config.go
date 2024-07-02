package common

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	GosterConfigPath      = "/etc/gost/goster.json"
	GostConfigDefaultPath = "/etc/gost/gost.json"
)

var (
	GlobalConfig = &GostConfig{}
	ConfigPath   string
)

func InitConfig() {
	if GostConfigDefaultPath == ConfigPath {
		gosterC, err := os.ReadFile(GosterConfigPath)
		if err == nil {
			c := &GosterConfig{}
			if err := json.Unmarshal(gosterC, &c); err == nil && c.ConfigName != "" {
				ConfigPath = c.ConfigName
			}
		}
	}
	_, err := os.Stat(ConfigPath)
	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Println("Failed to check file /etc/gost/gost.json")
		}
		return
	}
	b, err := os.ReadFile(ConfigPath)
	if err != nil {
		fmt.Errorf("read config error: %v", err)
		return
	}
	if err := json.Unmarshal(b, GlobalConfig); err != nil {
		fmt.Errorf("unmarshal config error: %v", err)
		return
	}
}
