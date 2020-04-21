package client

import (
	"log"
	"spider/crawler_distributed/config"
	"spider/crawler_distributed/rpcsupport"
	"spider/engine"
)

//存储
func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host) //启一个rpc客户端
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out //这边是只有一个chan在进行存储内容获取，并实时获得存储结果返回，开多个同时进行也行
			itemCount++
			log.Printf("Item Saver:got item #%d : %v\n", itemCount, item)
			//Call RPC to save item 分布式，调用一个服务器上的RPC服务来进行内容存储
			result := ""
			err := client.Call(config.ItemSaverRpc, item, &result)
			if err != nil {
				log.Printf("Error saving item %v:%v", item, err)
			}
		}
	}()
	return out, nil
}
