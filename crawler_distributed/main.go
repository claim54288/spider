package main

import (
	"spider/crawler_distributed/config"
	"spider/crawler_distributed/persist/client"
	"spider/engine"
	"spider/scheduler"
	"spider/webs/zongheng/parser"
)

func main() {
	itemChan, err := client.ItemSaver(config.ItemSaverPort) //数据存储,分布式，存储器使用rpc来让远程服务器进行存储
	//itemChan, err := persist.ItemSaver("dating_student") //数据存储
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 50,
		ItemChan:    itemChan,
	}
	//os.Mkdir("books", 777) //小说内容存储目录，不看，不存了
	e.Run(engine.Request{
		Url:        "http://book.zongheng.com/store/c0/c0/b0/u0/p1/v9/s9/t0/u0/i1/ALL.html",
		ParserFunc: parser.ParseCatalog,
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
