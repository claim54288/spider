package parser

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"spider/fetcher"
	"testing"
)

func TestParseNovelHomePage(t *testing.T) {
	content, err := fetcher.Fetch("http://book.zongheng.com/book/482604.html")
	if err != nil {
		panic(err)
	}
	content.Find(".nums>span>i").Each(func(i int, selection *goquery.Selection) {
		fmt.Printf("%d, %s", i, selection.Text())
	})

}
