package main

import (
	"context"
	"fmt"
	"net/http"
)

func serve(addr string, handler http.Handler, stop <-chan struct{}) error  {
	s := http.Server{Addr: addr, Handler: handler}
	s.Handler = http.DefaultServeMux
	go func(){
		<-stop
		s.Shutdown(context.Background())
	}()
	return s.ListenAndServe()
}

type debugHandler struct {}

func (h *debugHandler)ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "debug")
}

type appHandler struct {
}

func (h *appHandler)ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func main() {
	done := make(chan error, 2)
	stop := make(chan struct{})
	go func() {
		done <- serve("0.0.0.0:8081",&debugHandler{}, stop)
	}()

	go func() {
		done <- serve("0.0.0.0:8080",&appHandler{}, stop)
	}()

	var stopped bool
	for i:=0; i<cap(done); i++ {
		if err := <-done; err!=nil {
			fmt.Println("error: %v", err)
		}
		if !stopped {
			stopped = true
			close(stop)
		}
	}
	
}