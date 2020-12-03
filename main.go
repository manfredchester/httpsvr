package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os/exec"
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
func run(w http.ResponseWriter, r *http.Request) {
	httpExec(w)
}

func httpExec(w http.ResponseWriter) {
	w.Header().Set("Connection", "Keep-Alive")
	w.Header().Set("Transfer-Encoding", "chunked")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	pr, pw := io.Pipe()
	c := exec.Command("ping", "127.0.0.1", "-t")
	c.Stdout = pw
	c.Stderr = pw
	err := c.Start()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	go func() {
		c.Wait()
		pw.Close()
	}()
	defer pr.Close()
	s := bufio.NewScanner(pr)
	for s.Scan() {
		fmt.Fprintln(w, s.Text())
		w.(http.Flusher).Flush()
	}
	err = s.Err()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
