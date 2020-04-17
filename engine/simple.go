package engine

import (
	"log"
	"spider/fetcher"
)

type SimpleEngine struct{}

func (e *SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, err := worker(r)
		if err != nil {
			continue
		}
		//fmt.Printf("%v", parseResult)
		requests = append(requests, parseResult.Requests...) //页面里要继续爬的路由加进任务队列
		for _, item := range parseResult.Items {
			log.Printf("Got item %+v\n", item)
		}
	}
}

func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher:error fetching url %s:%v", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body), err
}
