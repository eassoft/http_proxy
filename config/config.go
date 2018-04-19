package config

import (
	"encoding/json"

	"os"
	//"encoding/json"

	"hm/eclibs/logger"
)

type ProxyConfig struct {
	Pattern string `json:"pattern"` //子目录 /代表全部
	Scheme  string `json:"scheme"`  //协议：http 或者 https
	Host    string `json:"host"`    //地址
}

var (
	ProxyConfigArr []*ProxyConfig //代理数组
	ListenAddr     = ":88"
	TLSListenAddr  = ":443"

	Dev = "false" //调试true 发布 false
)

func init() {
	//读取需要代理的内容 配置文件
	r, err := os.Open("conf/config.json")
	if err != nil {
		logger.Error(err)
	}
	decoder := json.NewDecoder(r)
	err = decoder.Decode(&ProxyConfigArr)
	if err != nil {
		logger.Error(err)
	} else {
		logger.Debug(ProxyConfigArr)
	}
}
