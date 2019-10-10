package main

import (
	"net/http"
)

// Golang开启http服务的三种方式
// https://www.jianshu.com/p/fe502c586034

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/bye", sayBye)

	server := &http.Server{
		Addr:    ":9898",
		Handler: mux,
	}
	server.ListenAndServe()
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bye bye ,this is v3 httpServer"))
}

type myHandler struct{}

func (e *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is version 3"))
}
