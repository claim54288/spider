package main

import (
	"spider/crawler_distributed/rpcsupport"
	"spider/engine"
	"spider/model"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"

	//start ItemSaverServer
	go serveRpc(host, "test1")
	time.Sleep(time.Second)

	//start ItemSaverClient
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	//Call save
	item := engine.Item{
		Url:  "http://book.zongheng.com/book/849617.html",
		Type: "zongheng",
		Id:   "849617",
		Payload: model.Novel{
			BookId:             849617,
			Author:             "关中老人",
			Name:               "一脉承腔",
			LetterNum:          15326236,
			RecommendedTotally: 663,
			Clicked:            20,
			RecommendedWeekly:  230,
			Chapter:            "",
			TotalChapterNum:    0,
			Title:              "",
			Contents:           "",
		},
	}

	result := ""
	err = client.Call("ItemSaverService.Save", item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result:%s;err:%s", result, err)
	}
}
