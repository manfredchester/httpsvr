package main

import (
	"github.com/zhxu/go-conf"
)

type Configuration struct {
	WORKPATH string
}

var cf Configuration

func loadConfig(fn string) {
	//default values
	cf.WORKPATH = "D:/gitproject"
	if fn != "" {
		assert(conf.ParseFile(fn, &cf))
	}
}
