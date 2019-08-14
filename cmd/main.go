package main

import (
	"flag"
	"os"
	"os/signal"
	"log"
	"syscall"

	"go-admin/conf"
	"go-admin/http"
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}

	//http init
	http.Init(conf.Conf)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT, syscall.SIGHUP)
	for {
		sg := <-c
		log.Printf("main exit by signal(%s)\n", sg.String())
		switch sg {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
			http.Close()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
