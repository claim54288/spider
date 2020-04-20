package parser

import (
	"github.com/PuerkitoBio/goquery"
	"spider/engine"
	"spider/model"
	"strconv"
	"strings"
)

func ParseNovelContent(contents *goquery.Document, novel model.Novel) engine.ParseResult {
	title := contents.Find(".title_txtbox").First().Text()
	novel.Title = title
	novel.Contents, _ = contents.Find(".content").Html()

	result := engine.ParseResult{}
	result.Items = append(result.Items, engine.Item{
		Url:     contents.Url.String(),
		Type:    "zongheng",
		Id:      strconv.Itoa(novel.BookId),
		Payload: novel,
	})
	href, _ := contents.Find(".chap_btnbox>a").Last().Attr("href")
	if strings.HasPrefix(href, "http") {
		result.Requests = append(result.Requests, engine.Request{
			Url: href,
			ParserFunc: func(document *goquery.Document) engine.ParseResult {
				return ParseNovelContent(document, novel)
			},
		})
	}
	return result
}
