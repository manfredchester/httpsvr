package main

import (
	"errors"
	"path"
	proconfig "protocol/config"
	"pub/zhlog"

	"github.com/hashicorp/consul/api"
)

func WriteConsulInfo(key, val string) (e error) {
	traceID := zhlog.UUID(8)
	defer func() {
		if e := recover(); e != nil {
			zhlog.Error(traceID, "写入Consul配置发生错误:", e.(error))
		}
	}()
	if path.IsAbs(key) {
		zhlog.Error(traceID, "key应为相对路径")
		return errors.New("key应为相对路径")
	}
	value := []byte(val)
	consulClient, err := api.NewClient(&api.Config{Scheme: "http", Address: proconfig.CConsulAddr()})
	zhlog.Assert(err)
	_, err = consulClient.KV().Put(&api.KVPair{Key: key, Value: value}, nil)
	zhlog.Assert(err)
	return nil
}
