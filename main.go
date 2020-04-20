package main

import (
	"os"
	"spider/engine"
	"spider/persist"
	"spider/scheduler"
	"spider/zongheng/parser"
)

func main() {
	itemChan, err := persist.ItemSaver("dating_novel") //数据存储
	//itemChan, err := persist.ItemSaver("dating_student") //数据存储
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 50,
		ItemChan:    itemChan,
	}
	os.Mkdir("books", 777)
	e.Run(engine.Request{
		Url:        "http://book.zongheng.com/store/c0/c0/b0/u0/p1/v9/s9/t0/u0/i1/ALL.html",
		ParserFunc: parser.ParseCatalog,
	})

	//e.Run(engine.Request{
	//	Url:        "http://book.zongheng.com/book/903377.html",
	//	ParserFunc: parser.ParseNovelHomePage,
	//})
}
