package view

import (
	"os"
	"spider/engine"
	"spider/frontend/model"
	model2 "spider/model"
	"testing"
)

func TestSearchResultView_Render(t *testing.T) {
	view := CreateSearchResultView("template.html")
	out, err := os.Create("template.test.html")
	page := model.SearchResult{}
	page.Hits = 123
	item := engine.Item{
		Url:  "http://book.zongheng.com/book/849617.html",
		Type: "zongheng",
		Id:   "849617",
		Payload: model2.Novel{
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
	page.Items = []engine.Item{item, item, item, item}
	err = view.Render(out, page)
	if err != nil {
		panic(err)
	}
}
