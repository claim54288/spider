package main

import (
	"flag"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"spider/crawler_distributed/config"
	"spider/crawler_distributed/persist"
	"spider/crawler_distributed/rpcsupport"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("必须指定监听端口")
		return
	}
	log.Fatal(serveRpc(fmt.Sprintf(":%d", *port), config.ElasticIndexNovel)) //log.Fatal得到值的话就是表示里面函数出错了，不然死循环出不来,出来了就直接状态1退出
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.159.131:9200"),
		elastic.SetSniff(false),
	)
	if err != nil {
		return err
	}
	err = rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
	return err
}