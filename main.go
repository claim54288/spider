package main

import (
	"os"
	"spider/engine"
	"spider/persist"
	"spider/scheduler"
	"spider/zongheng/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 50,
		ItemChan:    persist.ItemSaver(),
	}
	os.Mkdir("books", 777)
	e.Run(engine.Request{
		Url:        "http://book.zongheng.com/store/c0/c0/b0/u1/p1/v0/s1/t0/u0/i1/ALL.html",
		ParserFunc: parser.ParseCatalog,
	})

	//e.Run(engine.Request{
	//	Url:        "http://book.zongheng.com/book/903377.html",
	//	ParserFunc: parser.ParseNovelHomePage,
	//})
}
