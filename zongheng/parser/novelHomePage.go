package parser

import (
	"github.com/PuerkitoBio/goquery"
	jsoniter "github.com/json-iterator/go"
	"spider/engine"
	"spider/model"
	"strconv"
	"strings"
)

func ParseNovelHomePage(contents *goquery.Document) engine.ParseResult {
	result := engine.ParseResult{}

	contents.Find(".book-info").Each(func(i int, selection *goquery.Selection) {
		href, _ := selection.Find(".read-btn").First().Attr("href")
		if !strings.HasPrefix(href, "http") {
			return
		}
		result.Requests = append(result.Requests, engine.Request{
			Url:        href,
			ParserFunc: engine.NilParser,
		})

		bookIdJson, _ := selection.Find(".book-label .state").Attr("data-sa-d")
		book_id := jsoniter.Get([]byte(bookIdJson), "book_id").ToInt()

		novel := model.Novel{
			BookId: book_id,
			Name:   strings.TrimSpace(selection.Find(".book-name").Text()),
		}
		selection.Find(".nums>span>i").Each(func(i int, s *goquery.Selection) {
			switch i {
			case 0:
				novel.LetterNum = s.Text()
			case 1:
				num, _ := strconv.Atoi(s.Text())
				novel.RecommendedTotally = num
			case 2:
				num, _ := strconv.Atoi(s.Text())
				novel.Clicked = num
			case 3:
				num, _ := strconv.Atoi(s.Text())
				novel.RecommendedWeekly = num
			}
		})
		result.Items = append(result.Items, novel)
	})

	return result
}
