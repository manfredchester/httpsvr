package common

import (
	"github.com/zhxu/go-conf"
)

type Configuration struct {
	HTTP_PORT     string
	READ_TIMEOUT  int
	WRITE_TIMEOUT int
	WebRoot       string
}

var cf Configuration

func loadConfig(fn string) {
	//default values
	cf.HTTP_PORT = "3349"
	cf.READ_TIMEOUT = 60
	cf.WRITE_TIMEOUT = 60
	cf.WebRoot = "/resources"
	if fn != "" {
		assert(conf.ParseFile(fn, &cf))
	}

}
