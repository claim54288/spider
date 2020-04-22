package main

import (
	"os"
	"spider/crawler_distributed/config"
	"spider/engine"
	"spider/persist"
	"spider/scheduler"
	"spider/webs/zongheng/parser"
)

func main() {
	itemChan, err := persist.ItemSaver("dating_novel") //数据存储
	//itemChan, err := persist.ItemSaver("dating_student") //数据存储
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      50,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}
	os.Mkdir("books", 777)
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
