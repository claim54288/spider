package parser

import (
	"github.com/PuerkitoBio/goquery"
	"spider/engine"
)

func ParseCityList(contexts *goquery.Document) engine.ParseResult {
	result := engine.ParseResult{}

	contexts.Find(".city-list dd>a").EachWithBreak(func(i int, selection *goquery.Selection) bool {
		if i == 10 {
			return false
		}
		href, _ := selection.Attr("href")
		result.Requests = append(result.Requests, engine.Request{
			Url:        href,
			ParserFunc: ParseCity,
		})
		result.Items = append(result.Items, selection.Text())
		return true
	})
	return result
}
