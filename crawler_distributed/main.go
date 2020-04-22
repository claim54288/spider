package main

import (
	"flag"
	"log"
	"net/rpc"
	"spider/crawler_distributed/config"
	itemsaver "spider/crawler_distributed/persist/client"
	"spider/crawler_distributed/rpcsupport"
	worker "spider/crawler_distributed/worker/client"
	"spider/engine"
	"spider/scheduler"
	"spider/webs/zongheng/parser"
	"strings"
)

var (
	itemSaverHost = flag.String("itemsaver_host", "", "itemsaver host")
	workerHosts   = flag.String("worker_hosts", "", "worker hosts (comma separated)逗号分隔的一堆host")
)

func main() {
	flag.Parse()

	itemChan, err := itemsaver.ItemSaver(*itemSaverHost) //数据存储,分布式，存储器使用rpc来让远程服务器进行存储
	//itemChan, err := persist.ItemSaver("dating_student") //数据存储
	if err != nil {
		panic(err)
	}

	pool := createClientPool(strings.Split(*workerHosts, ",")) //没验证，略过

	processor := worker.CreateProcessor(pool)

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      50,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
	//os.Mkdir("books", 777) //小说内容存储目录，不看，不存了
	e.Run(engine.Request{
		Url:    "http://book.zongheng.com/store/c0/c0/b0/u0/p1/v9/s9/t0/u0/i1/ALL.html",
		Parser: engine.NewFuncParser(parser.ParseCatalog, config.ParseCatalog),
	})

	//e.Run(engine.Request{
	//	Url:        "http://book.zongheng.com/book/903377.html",
	//	ParserFunc: parser.ParseNovelHomePage,
	//})

	//e.Run(engine.Request{
	//	Url:        "http://test.b.kai12.cn/archives/archives?grade_id=101&class_id=13",
	//	ParserFunc: parser.ParseStudentList,
	//})
	//e.Run(engine.Request{
	//	Url:        "http://test.b.kai12.cn/archives/archives/student?student_id=258&student_edit=1",
	//	ParserFunc: parser.ParseStudentHomePage,
	//})
}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s", h)
		} else {
			log.Printf("error connecting to %s:%v", h, err)
		}
	}

	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}
