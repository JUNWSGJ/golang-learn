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
	server *http.Server
	name   string
	addr   string
	ctx    context.Context
}

func (s *HttpServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	if path == "/shutdown" {
		n, err := fmt.Fprintf(w, "shutdown server[%s] success\n", s.name)
		log.Printf("n: %d, err: %v \n", n, err)
		log.Printf("app[%s]被人为关闭\n", s.name)
		// 模拟APP异常退出，检测其它app是否会同步注销退出
		_ = s.Stop()
	} else {
		n, err := fmt.Fprintf(w, "server[%s], is serving on addr[%s], current URL.Path = %q\n", s.name, s.addr, path)
		log.Printf("n: %d, err: %v \n", n, err)
	}
}

func NewHttpServer(ctx context.Context, name string, addr string) *HttpServer {
	srv := &HttpServer{
		name: name,
		addr: addr,
		ctx:  ctx,
	}
	srv.server = &http.Server{Handler: srv, Addr: addr}
	return srv

}

func (s *HttpServer) Run() error {
	defer func() {
		log.Printf("优雅退出server【%s】\n", s.name)
	}()

	stop := make(chan error)
	go func() {
		log.Printf("Server[%s] running\n", s.name)
		stop <- s.server.ListenAndServe()
	}()

	select {
	case <-s.ctx.Done():
		log.Printf("app[%s]接收到context退出通知，shutdown\n", s.name)
		return s.Stop()
	case stopErr := <-stop:
		return errors.Wrapf(stopErr, "app[%s]异常终止", s.name)
	}
}

func (s *HttpServer) Stop() error {

	log.Printf("server[%s] stopping\n", s.name)
	return s.server.Shutdown(s.ctx)
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(ctx)

	server1 := NewHttpServer(ctx, "server1", ":8001")
	server2 := NewHttpServer(ctx, "server2", ":8002")

	g.Go(func() error {
		return server1.Run()
	})
	g.Go(func() error {
		return server2.Run()
	})

	exit := make(chan os.Signal)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	g.Go(func() error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-exit:
			cancel()
			return nil
		}
	})

	log.Printf("程序已运行，当前进程pid：%d \n", os.Getpid())

	if err := g.Wait(); err != nil {
		log.Printf("find app run error: %v\n", err.Error())
		log.Printf("close all apps and exit.")
	}

}
