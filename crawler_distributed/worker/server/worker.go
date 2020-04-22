package main

import (
	"flag"
	"fmt"
	"log"
	"spider/crawler_distributed/rpcsupport"
	"spider/crawler_distributed/worker"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("必须指定监听端口")
		return
	}
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", *port), worker.CrawlService{}))
}
