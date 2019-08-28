package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

func main() {
	ver := flag.Bool("version", false, "show version info")
	conf := flag.String("conf", "", "configuration file")
	// pkg := flag.String("pack", "", "pack resources under directory")
	// dbg := flag.Bool("debug", false, "debug mode")
	flag.Parse()
	if *ver {
		fmt.Println(verinfo())
		return
	}
	loadConfig(*conf)
	val := `"key1":1}`
	_, e1 := GetConsulInfo("path0/config/inspect/path2")
	if e1 != nil {
		fmt.Println("写入consul。。。")
		e2 := WriteConsulInfo("path0/config/inspect/path2", val)
		fmt.Println("e2:", e2)
	} else {
		fmt.Println("consul已存在。。。")
		e3 := DelConsulInfo("/path0/config/inspect/path2")
		fmt.Println("e3:", e3)

	}
	// if *pkg != "" {
	// 	audit.Assert(res.Pack(*pkg))
	// 	fmt.Printf("resources under '%s' packed.\n", *pkg)
	// 	return
	// }
	// if !*dbg {
	// 	audit.Assert(res.Extract(cf.WebRoot, res.OverwriteIfNewer))
	// }
	// audit.ExpVars(map[string]interface{}{
	// 	"config":  cf,
	// 	"version": _G_REVS + "." + _G_HASH,
	// })
	// audit.SetDebugging(*dbg)

	// mux := http.NewServeMux()
	// mux.HandleFunc("/", home)
	setupRoutes()
	svr := http.Server{
		Addr: ":" + cf.HTTP_PORT,
		// Handler:      mux,
		ReadTimeout:  time.Duration(cf.READ_TIMEOUT) * time.Second,
		WriteTimeout: time.Duration(cf.WRITE_TIMEOUT) * time.Second,
	}
	svr.ListenAndServe()

}
