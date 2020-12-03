package main

import (
	"net/http"
	"time"
)

func main() {
	// val := `"key1":1}`
	// _, e1 := GetConsulInfo("path0/config/inspect/path2")
	// if e1 != nil {
	// 	fmt.Println("写入consul。。。")
	// 	e2 := WriteConsulInfo("path0/config/inspect/path2", val)
	// 	fmt.Println("e2:", e2)
	// } else {
	// 	fmt.Println("consul已存在。。。")
	// 	e3 := DelConsulInfo("/path0/config/inspect/path2")
	// 	fmt.Println("e3:", e3)
	// }
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", home)
	setupRoutes()
	svr := http.Server{
		Addr: ":9966",
		// Handler:      mux,
		ReadTimeout:  time.Duration(10) * time.Second,
		WriteTimeout: time.Duration(10) * time.Second,
	}
	svr.ListenAndServe()
}
