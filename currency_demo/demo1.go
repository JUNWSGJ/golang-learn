package main

import (
	"fmt"
	"log"
	"net/http"
)

/**
 使用goroutine监听，主线程用select{} 永远阻塞
 不推荐的做法，
 */
func startHttpServer1(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World!")
	})
	go func() {
		if err:= http.ListenAndServe(":8080", nil); err!=nil {
			log.Fatal(err)
		}
	}()
	select {}
}

/**
 正常的做法，不支持同时启动两个server
 */
func startHttpServer2(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World!")
	})
	if err:= http.ListenAndServe(":8080", nil); err!=nil {
		log.Fatal(err)
	}
}


func serveApp(){
	mux :=http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, QCon!")
	})
	if err:= http.ListenAndServe(":8080", mux);err!=nil {
		log.Fatal(err)
	}
}

func serveDebug(){
	if err:= http.ListenAndServe("127.0.0.1:8001", http.DefaultServeMux);err!=nil {
		log.Fatal(err)
	}
}

/**
 启动多个server， 一个在goroutine执行，一个在主线程执行，
 缺点是 serveDebug退出后，serveApp不能感知
 */
func startMultiHttpServer1(){
	go serveDebug()
	serveApp()
}

/**
 使用fatal退出时不会调用defer
 只建议在 main和init函数里使用fatal
 */
func startMultiHttpServer2(){
	go serveDebug()
	go serveApp()
	select {}
}

/**
 启动httpserver的多种方法
 */
func main() {
	//startHttpServer1()
	//startHttpServer2()
	//startMultiHttpServer1()
	startMultiHttpServer2()
}
