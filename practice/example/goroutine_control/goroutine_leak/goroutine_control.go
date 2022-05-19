package main

import (
	"context"
	"fmt"
	"net/http"
)

func main()  {
	//leak()
	//time.Sleep(time.Second)
	done:=make(chan error,2)
	stop:=make(chan struct{})
	go func() {
		done<-server("127.0.0.1:8888", func() http.Handler{return http.Handler() },stop)
	}()
}
func server(addr string,handler http.Handler,stop<-chan struct{}) error {
	s:=http.Server{
		Addr: addr,
		Handler: handler,
	}
	go func() {
		<-stop
		s.Shutdown(context.Background())
	}()
	return s.ListenAndServe()
}
func leak()  {
	//ch是局部变量，没有任何数据会发生到chan chan也不会关闭
	//leak执行完成后 内部goroutine无法退出 一直阻塞
	ch:=make(chan int)
	go func() {
		val:=<-ch
		fmt.Println("we receive a value",val)
	}()
}
