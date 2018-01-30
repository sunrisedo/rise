package main

import (
	"log"
	"net/http"
	"runtime"

	"github.com/sunrisedo/conf"
)

var (
	cfg = conf.NewConfig("init.conf")
)

func init() {
	// 初始化配置文件
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println("init data start...")

	log.Println("init data finish.")
}

func main() {

	//注册路由
	log.Println("init route start...")
	for addr, controller := range RouteMap {
		http.HandleFunc(addr, controller)
	}
	log.Println("init route finish. listen port", cfg.Read("server", "port"))
	// http.Handle("/webroot/", http.FileServer(http.Dir("webroot")))
	log.Println(http.ListenAndServe(cfg.Read("server", "port"), nil))
}
