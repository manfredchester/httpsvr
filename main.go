package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"time"
)

type helloHandler struct{}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func main() {
	ser := "http"
	switch ser {
	case "http":
		httpServer()
	case "mux":
		muxServer()
	case "server":
		server()
	}
}

func server() {
	http.HandleFunc("/run", run)
	svr := http.Server{
		Addr: ":9966",
		// Handler:      mux,
		ReadTimeout:  time.Duration(10) * time.Second,
		WriteTimeout: time.Duration(10) * time.Second,
	}
	svr.ListenAndServe()
}

func muxServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/run", run)
	http.ListenAndServe(":9966", mux)
}

func httpServer() {
	// Handle 对象结构需要实现ServeHTTP的方法
	http.Handle("/hello", &helloHandler{})
	// HandleFunc 中的对象HandleFunc实则是个适配器，最终还是会指向ServeHTTP
	http.HandleFunc("/run", run)
	// 文件服务器 或 负责重定向的RedirectHandler
	// http.ListenAndServe(":9966", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":9966", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TODO: handle requests\n")
}

func run(w http.ResponseWriter, r *http.Request) {
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
