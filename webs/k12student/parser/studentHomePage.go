package parser

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"spider/engine"
)

func ParseStudentHomePage(contents *goquery.Document) engine.ParseResult {
	result := engine.ParseResult{}
	last := contents.Find(".registration-card a").Last()
	href, _ := last.Attr("href")
	if href == "" {
		return result
	}
	href = fmt.Sprintf("%s://%s%s", contents.Url.Scheme, contents.Url.Host, href)
	result.Requests = append(result.Requests, engine.Request{
		Url:        href,
		ParserFunc: ParseStudentInfo,
	})
	result.Items = append(result.Items, engine.Item{
		Url:     href,
		Type:    "student_list",
		Id:      "",
		Payload: last.Text(),
	})

	return result
}
