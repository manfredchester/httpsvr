package main

import (
	"errors"
	"path/filepath"
	proconfig "protocol/config"
	"pub/zhlog"

	"github.com/hashicorp/consul/api"

	"github.com/micro/go-config/source/consul"

	"github.com/micro/go-config"
)

// 默认consul value值为map类型的字符串
var confmap = make(map[string]interface{})

// 访问key 应为全称且为相对路径
func GetConsulInfo(key string) (b []byte, e error) {
	traceID := zhlog.UUID(8)
	defer func() {
		if e := recover(); e != nil {
			zhlog.Error(traceID, "获取Consul信息失败:%s", e.(error))
		}
	}()

	prefix, _ := filepath.Split(key)
	consulAddr := proconfig.ConsulAddress()
	configSer := config.NewConfig()
	err := configSer.Load(consul.NewSource(
		consul.WithAddress(consulAddr),
		consul.WithPrefix(prefix),
		consul.StripPrefix(true),
	))
	if err != nil {
		zhlog.Log(traceID, "prefix:%s 不存在将立即建立", prefix)
	}

	var defaultConf = api.DefaultConfig()
	defaultConf.Address = consulAddr
	consulCli, err := api.NewClient(defaultConf)
	zhlog.Assert(err)

	value := configSer.Get("path1/path2")
	err = value.Scan(&confmap)
	zhlog.Assert(err)

	// 路径不存在触发404 但是底层包不返回err，且kvPare 为nil
	kvPair, _, err := consulCli.KV().Get(key, nil)
	zhlog.Assert(err)
	if kvPair == nil {
		zhlog.Error(traceID, "该配置信息%s 不存在", key)
		return nil, errors.New("该配置信息不存在")
	}
	// 由应用服务进行格式类型校验
	// r := bytes.NewReader(kvPair.Value)
	// err = json.NewDecoder(r).Decode(&confmap)
	// zhlog.Assert(err)

	return kvPair.Value, nil
}
