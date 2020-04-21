package main

import (
	"gopkg.in/olivere/elastic.v5"
	"log"
	"spider/crawler_distributed/persist"
	"spider/crawler_distributed/rpcsupport"
)

func main() {
	log.Fatal(serveRpc(":1234", "dating_novel")) //log.Fatal得到值的话就是表示里面函数出错了，不然死循环出不来,出来了就直接状态1退出
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
