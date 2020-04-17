package scheduler

import (
	"spider/engine"
)

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s *QueuedScheduler) ConfigureMasterWorkerChan(chan engine.Request) {
	panic("implement me")
}

//这边这种写法是为了通过调度器来控制发给哪个处理线程，而不是一堆处理线程在那边抢，方便进行负载均衡等(虽然这边并没有实现)
func (s *QueuedScheduler) Run() {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)

	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request //通过var声明不初始化的chan是nil，不会报错或者阻塞 case <-的时候不会进这个分支
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}

			select {
			case r := <-s.requestChan:
				//把 r 发送到 处理器
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				//同时有请求和处理器存在的时候尝试进行这个操作，但是并不能保证进行，所以在成功进行之后才进行任务的删除
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
