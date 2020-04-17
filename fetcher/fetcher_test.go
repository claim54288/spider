package fetcher

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"testing"
)

func TestFetch(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://book.zongheng.com/book/909369.html", nil)
	if err != nil {
		log.Printf("wrong status code:%d", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}
	document.Find(".nums>span").Filter("i").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Html())
	})

}
