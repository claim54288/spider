package parser

import (
	"github.com/PuerkitoBio/goquery"
	jsoniter "github.com/json-iterator/go"
	"github.com/shopspring/decimal"
	"spider/engine"
	"spider/model"
	"strconv"
	"strings"
)

func ParseNovelHomePage(contents *goquery.Document) engine.ParseResult {
	result := engine.ParseResult{}

	selection := contents.Find(".book-info").First()
	href, _ := selection.Find(".read-btn").First().Attr("href")
	if !strings.HasPrefix(href, "http") {
		return result
	}

	bookIdJson, _ := selection.Find(".book-label .state").Attr("data-sa-d")
	book_id := jsoniter.Get([]byte(bookIdJson), "book_id").ToInt()

	novel := model.Novel{
		BookId: book_id,
		Name:   strings.TrimSpace(selection.Find(".book-name").Text()),
		Author: contents.Find(".au-name a").Text(),
	}
	selection.Find(".nums>span>i").Each(func(i int, s *goquery.Selection) {
		switch i {
		case 0:
			novel.LetterNum = numHandle(s.Text())
		case 1:
			novel.RecommendedTotally = numHandle(s.Text())
		case 2:
			novel.Clicked = numHandle(s.Text())
		case 3:
			novel.RecommendedWeekly = numHandle(s.Text())
		}
	})
	result.Items = append(result.Items, engine.Item{
		Url:     contents.Url.String(),
		Type:    "zongheng",
		Id:      strconv.Itoa(novel.BookId),
		Payload: novel,
	})

	result.Requests = append(result.Requests, engine.Request{
		Url:    href,
		Parser: NewNovelContentParser(novel),
	})
	return result
}

func numHandle(s string) int64 {
	s = strings.TrimSpace(s)
	if strings.HasSuffix(s, "万") {
		s = strings.TrimSuffix(s, "万")
		f, _ := decimal.NewFromString(s)
		mul := f.Mul(decimal.New(10000, 0))
		return mul.IntPart()
	} else {
		f, _ := decimal.NewFromString(s)
		return f.IntPart()
	}
}
