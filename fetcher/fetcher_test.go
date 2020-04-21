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
	req, err := http.NewRequest("GET", "http://test.b.kai12.cn/archives/archives?grade_id=101&class_id=13", nil)
	if err != nil {
		log.Printf("wrong status code:%d", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36")

	req.AddCookie(&http.Cookie{
		Name:   "2bdoc",
		Value:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBsaWNhdGlvbiI6InRvYiIsImNoaWxkX3VzZXJfaWQiOjAsInNjaG9vbF9jb2RlIjoiazEyc2Nob29sXzAzX3Rlc3QiLCJ1c2VyX2lkIjo1NTcsImV4cCI6MTU4NzYzMDQ3Nn0.TfKJVOwPecEfp2CRPbIV3EXxpruo0IU53D4oonoNax0",
		Path:   "/",
		Domain: ".kai12.cn",
	})
	req.AddCookie(&http.Cookie{
		Name:   "2cdoc",
		Value:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBsaWNhdGlvbiI6InRvZCIsInNjaG9vbF9jb2RlIjoiazEyc2Nob29sXzAzX3Rlc3QiLCJ1c2VyX2lkIjo5NDYsImV4cCI6MTU4NzYzMDMwNX0.3rFDUxuCi2sM8s71G7HMP_63OUuR4pY3aEEBtNlsJqo",
		Path:   "/",
		Domain: ".kai12.cn",
	})
	req.AddCookie(&http.Cookie{
		Name:   "k12_tob",
		Value:  "e74610bc01ec5b40dc59abbc25188dd6cb911e46",
		Path:   "/",
		Domain: "test.b.kai12.cn",
	})
	req.AddCookie(&http.Cookie{
		Name:   "k12school_03_test_32_102_26_2016_2",
		Value:  "%5B%5D",
		Path:   "/",
		Domain: "test.b.kai12.cn",
	})

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println(resp.Request.URL.String())
	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}
	//document.Find(".nums>span").Filter("i").Each(func(i int, selection *goquery.Selection) {
	fmt.Println(document.Html())
	//})

}
