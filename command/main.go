package main

import (
	"flag"
	"fmt"
	"os/exec"
	"path/filepath"
	"pub/log"
)

func main() {
	defer func() {
		if e := recover(); e != nil {
			log.Error("main()", "%s", e.(error))
		}
	}()
	conf := flag.String("conf", "", "configuration file")
	project := flag.String("p", "", "service name")
	branch := flag.String("b", "", "choose branch default master")
	hash := flag.String("h", "", "commit hash")
	flag.Parse()
	if *project == "" {
		log.Error("command-line-arguments", "project empty")
		return
	}
	if *branch == "" {
		log.Error("command-line-arguments", "branch empty")
		return
	}
	if *hash == "" {
		log.Error("command-line-arguments", "hash empty")
		return
	}
	loadConfig(*conf)

	fmt.Println(*project, *branch, *hash)
	fmt.Println(filepath.Join(cf.WORKPATH, *project))
	data, err := filepath.Glob(filepath.Join(cf.WORKPATH, "/src", "/*"))

	fmt.Println(data, err)
	cmd := "d:"
	arg1 := ""
	c := exec.Command(cmd, arg1)
	fmt.Println("C:", c)
}

// cd /data/goworking/connexctsrc/src/businesscenter &&
// GOPATH=/data/goworking/connexctsrc:/data/goworking/connextlib GOROOT=/data/go /data/go/bin/go build
