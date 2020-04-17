package main

import (
	"spider/engine"
	"spider/scheduler"
	"spider/zongheng/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 50,
	}
	e.Run(engine.Request{
		Url:        "http://book.zongheng.com/store/c1/c0/b0/u0/p1/v9/s9/t0/u0/i1/ALL.html",
		ParserFunc: parser.ParseCatalog,
	})
}
