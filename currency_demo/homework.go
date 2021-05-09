package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)



/**
  目标：基于 errgroup 实现一个 http server 的启动和关闭，
  以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。
*/

type HttpServer struct {
	name string
	addr string
	ctx context.Context
	server http.Server
	shutdown func()
}


func (s *HttpServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	if path == "/shutdown" {
		fmt.Fprintf(w, "shutdown server[%s] success\n", s.name)
		log.Printf("app[%s]被人为关闭\n", s.name)
		// 模拟APP异常退出，检测其它app是否会同步注销退出
		s.shutdown()
	} else {
		fmt.Fprintf(w,"server[%s], is serving on addr[%s], current URL.Path = %q\n", s.name, s.addr, path)
	}

}

func startHttpServer(ctx context.Context, name string, addr string) error {
	defer func() {
		log.Printf("优雅退出server【%s】\n", name)
	}()

	app := &HttpServer{name: name, addr: addr}

	srv := http.Server{
		Addr:    addr,
		Handler: app,
	}

	app.shutdown = func() {
		srv.Shutdown(ctx)
	}

	stop := make(chan error)
	go func(){
		stop <- srv.ListenAndServe()
	}()

	exit := make(chan os.Signal)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-ctx.Done():
		log.Printf("app[%s]接收到context退出通知，shutdown\n", app.name)
		return srv.Shutdown(ctx)
	case <-exit:
		log.Printf("app[%s]接收到系统退出信号，shutdown\n", app.name)
		return srv.Shutdown(ctx)
	case stopErr := <- stop:
		return errors.Wrapf(stopErr, "app[%s]异常终止", app.name)
	}
}


func main() {

	ctx := context.Background()
	g, ctx:= errgroup.WithContext(ctx)

	g.Go(func() error {
		return startHttpServer(ctx,"app1", ":8001")
	})
	g.Go(func() error {
		return startHttpServer(ctx,"app2", ":8002")
	})
	log.Printf("程序已运行，当前进程pid：%d \n" ,os.Getpid())
	if err := g.Wait(); err != nil {
		log.Printf("find app run error: %v\n", err.Error())
		log.Printf("close all apps and exit.")
	}

}