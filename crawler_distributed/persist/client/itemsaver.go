package client

import (
	"log"
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
			item := <-out

			itemCount++
			log.Printf("Item Saver:got item #%d : %v\n", itemCount, item)
			//Call RPC to save item
			result := ""
			err := client.Call("ItemSaverService.Save", item, &result)
			if err != nil {
				log.Printf("Item Saver:error saving item %v:%v", item, err)
			}
		}
	}()
	return out, nil
}
