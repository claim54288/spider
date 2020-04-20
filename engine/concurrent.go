package engine

import (
	goredis "github.com/garyburd/redigo/redis"
	"log"
	"time"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan Item
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run() //调度器启动，进行request和worker的分发

	//启动worker，等待开始抓取和处理数据
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}
	defer redisDB.Close()
	for _, r := range seeds {
		//old := "p1"
		//for i := 2; i < 999; i++ {
		//	target := fmt.Sprintf("p%d", i)
		//	r.Url = strings.ReplaceAll(r.Url, old, target)
		//	old = target
		//	if isDuplicate(r.Url) {
		//		continue
		//	}
		//	e.Scheduler.Submit(r)
		//}

		if isDuplicate(r.Url) {
			continue
		}
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			//返回结果处理
			go func() { e.ItemChan <- item }()
		}

		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

var visitedUrl = make(map[string]bool)

//URL 去重
func isDuplicate(url string) bool {
	if visitedUrl[url] {
		return true
	}
	visitedUrl[url] = true
	return false
}

var redisDB goredis.Conn

func isDuplicateByRedis(url string) bool {
	i, err := goredis.Int(redisDB.Do("GET", url))
	if goredis.ErrNil == err {
		redisDB.Do("SET", url, 1)
		return false
	}
	if err != nil {
		log.Println(err)
	}
	if i == 1 {
		return true
	}
	redisDB.Do("SET", url, 1)
	return false
}

func init() {
	conn, err := goredis.Dial("tcp", "47.102.47.185:6379",
		goredis.DialPassword("claim"),
		goredis.DialDatabase(2),
		goredis.DialConnectTimeout(10*time.Second))
	if err != nil {
		panic(err)
	}
	redisDB = conn
}

func createWorker(in chan Request, out chan ParseResult, s ReadyNotifier) {
	go func() {
		for {
			//告诉调度器，准备好了,把自己准备接受request的chan丢给调度器
			s.WorkerReady(in)
			//准备好了之后就在这边等待任务获得
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
