package parser

import (
	"github.com/PuerkitoBio/goquery"
	"spider/engine"
)

func ParseCity(contents *goquery.Document) engine.ParseResult {
	result := engine.ParseResult{}
	contents.Find(".list-item .content tr:first-child a").Each(func(i int, s *goquery.Selection) {
		name := s.Text()
		result.Items = append(result.Items, "User "+name)
		href, _ := s.Attr("href")
		result.Requests = append(result.Requests, engine.Request{
			Url: href,
			ParserFunc: func(c *goquery.Document) engine.ParseResult {
				return ParseProfile(c, name)
			},
		})
	})
	return result
}
