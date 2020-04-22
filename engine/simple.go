package engine

import (
	"log"
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

		parseResult, err := Worker(r)
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
