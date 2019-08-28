package main

import (
	proconfig "protocol/config"
	"pub/zhlog"

	"github.com/hashicorp/consul/api"
)

func DelConsulInfo(key string) (e error) {
	traceID := zhlog.UUID(8)
	defer func() {
		if e := recover(); e != nil {
			zhlog.Error(traceID, "删除Consul配置发生错误:", e.(error))
			return
		}
	}()
	consulClient, err := api.NewClient(&api.Config{Scheme: "http", Address: proconfig.CConsulAddr()})
	zhlog.Assert(err)
	_, err = consulClient.KV().Delete(key, nil)
	zhlog.Assert(err)
	return nil
}
