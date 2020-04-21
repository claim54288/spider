package parser

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"spider/engine"
	"strings"
)

var regStudentId = regexp.MustCompile(`student_id=([a-zA-Z0-9]*)&`)

func ParseStudentList(contents *goquery.Document) engine.ParseResult {
	result := engine.ParseResult{}
	contents.Find(".title-links a").Each(func(i int, selection *goquery.Selection) {
		href, _ := selection.Attr("href")
		href = fmt.Sprintf("%s://%s%s", contents.Url.Scheme, contents.Url.Host, href)

		if strings.HasPrefix(href, "javascript") {
			return
		}

		result.Requests = append(result.Requests, engine.Request{
			Url:        href,
			ParserFunc: ParseStudentList,
		})
		result.Items = append(result.Items, engine.Item{
			Url:     href,
			Type:    "student_list",
			Id:      "",
			Payload: selection.Text(),
		})
	})

	contents.Find(".student-lists a").EachWithBreak(func(i int, selection *goquery.Selection) bool {
		href, _ := selection.Attr("href")
		href = fmt.Sprintf("%s://%s%s", contents.Url.Scheme, contents.Url.Host, href)

		if strings.HasPrefix(href, "javascript") || href == "" {
			return false
		}
		result.Requests = append(result.Requests, engine.Request{
			Url:        href,
			ParserFunc: ParseStudentHomePage,
		})
		title, _ := selection.Attr("title")
		//log.Printf("当前title：%s\n", title)
		result.Items = append(result.Items, engine.Item{
			Url:     href,
			Type:    "studentHomePage",
			Id:      getStudentId(regStudentId, href),
			Payload: title,
		})

		return true
	})
	//fmt.Printf("%+v,\n", result.Items)
	return result
}

func getStudentId(reg *regexp.Regexp, str string) string {
	s := reg.FindStringSubmatch(str)
	if len(s) > 1 {
		return s[1]
	}
	return ""
}
