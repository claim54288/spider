package parser

import (
	"github.com/PuerkitoBio/goquery"
	"spider/engine"
	"strings"
)

func ParseCatalog(contents *goquery.Document) engine.ParseResult {
	result := engine.ParseResult{}
	contents.Find(".bookname a").Each(func(i int, selection *goquery.Selection) {
		href, _ := selection.Attr("href")
		if !strings.HasPrefix(href, "http") {
			return
		}
		result.Requests = append(result.Requests, engine.Request{
			Url:        href,
			ParserFunc: ParseNovelHomePage,
		})
		result.Items = append(result.Items, engine.Item{
			Url:     "",
			Type:    "bookname",
			Id:      "",
			Payload: selection.Text(),
		})
	})
	return result
}
