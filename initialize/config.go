package initialize

import (
	"gin-server/global"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

const (
	commonFile = "common.yaml"
	serverFile = "server.yaml"
)

func CommonConfig() (err error) {
	var bytes []byte
	if bytes, err = os.ReadFile(commonFile); err != nil {
		return
	}
	if err = yaml.Unmarshal(bytes, &global.Common); err != nil {
		return
	}
	log.Printf("info %v 解析成功\n%+v", commonFile, global.Common)
	return
}

func ServerConfig() (err error) {
	var bytes []byte
	if bytes, err = os.ReadFile(serverFile); err != nil {
		return
	}
	if err = yaml.Unmarshal(bytes, &global.Server); err != nil {
		return
	}
	log.Printf("info %v 解析成功\n%+v", serverFile, global.Server)
	return
}
