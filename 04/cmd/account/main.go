package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type HttpServer struct {
	name   string
	addr   string
	ctx    context.Context
	cancel func()
}

func main() {
	SetUp()

	ctx, cancel := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(ctx)
	// 启动 http server
	g.Go(func() error {
		hs := &HttpServer{
			name:   "account server",
			addr:   ":8080",
			ctx:    ctx,
			cancel: cancel,
		}
		return hs.Run()
	})

	g.Go(func() error {
		exit := make(chan os.Signal)
		signal.Notify(exit, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
		select {
		case <-ctx.Done():
			return ctx.Err()
		case e := <-exit:
			fmt.Printf("signal %v\r\n", e)
			cancel()
			return nil
		}
	})

	if err := g.Wait(); err != nil {
		return
	}

}

func (hs *HttpServer) Run() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Printf("request: %s\r\n", request.URL.Path)
		_, err := fmt.Fprintf(writer, "%s hello, Gopher %s", hs.name, request.URL.Query().Get("a"))
		if err != nil {
			return
		}
	})

	server := &http.Server{Addr: hs.addr, Handler: mux}

	err := server.ListenAndServe()
	if err != nil {
		hs.cancel()
		return server.Close()
	}

	fmt.Printf("%s run\r\n", hs.name)
	// 监听退出
	select {
	case <-hs.ctx.Done():
		fmt.Printf("close %s http server\r\n", hs.name)
		return server.Close()
	}
}
