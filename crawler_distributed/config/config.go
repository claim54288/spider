package config

const (
	//端口
	ItemSaverPort = ":1234" //存储器运行端口
	WorkerPort0   = ":9000" //第一个worker的port

	//elasticsearch
	ElasticIndexNovel = "dating_novel"

	//RPC 调用方法
	ItemSaverRpc    = "ItemSaverService.Save"
	CrawlServiceRpc = "CrawlService.Process"

	//函数序列化名字
	NilParser          = "NilParser"
	ParseNovelHomePage = "ParseNovelHomePage"
	ParseCatalog       = "ParseCatalog"
	ParseNovelContent  = "ParseNovelContent"
)
