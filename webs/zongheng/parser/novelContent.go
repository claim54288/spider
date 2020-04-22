package parser

import (
	"github.com/PuerkitoBio/goquery"
	"spider/crawler_distributed/config"
	"spider/engine"
	"spider/model"
	"strconv"
	"strings"
)

func parseNovelContent(contents *goquery.Document, novel model.Novel) engine.ParseResult {
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
			Url:    href,
			Parser: NewNovelContentParser(novel),
		})
	}
	return result
}

type NovelContentParser struct {
	novel model.Novel
}

func (n *NovelContentParser) Parse(content *goquery.Document) engine.ParseResult {
	return parseNovelContent(content, n.novel)
}

func (n *NovelContentParser) Serialize() (name string, args interface{}) {
	return config.ParseNovelContent, n.novel
}

func NewNovelContentParser(n model.Novel) *NovelContentParser {
	return &NovelContentParser{novel: n}
}
