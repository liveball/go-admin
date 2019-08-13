package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/yangjian/lib/log"
	"github.com/yangjian/mainsite/conf"
	"github.com/yangjian/mainsite/http"
)

//go run main.go -conf mainsite.toml

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}

	//log init
	log.Init()

	//http init
	http.Init(conf.Conf)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT, syscall.SIGHUP)
	for {
		sg := <-c
		log.Infof("main exit by signal(%s)\n", sg.String())
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
