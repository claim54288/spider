package fetcher

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestFetch(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://book.zongheng.com/store/c1/c0/b0/u0/p1/v9/s9/t0/u0/i1/ALL.html", nil)
	if err != nil {
		log.Printf("wrong status code:%d", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36")
	req.AddCookie(&http.Cookie{
		Name:   "AST",
		Value:  "158715039955566c0f0e1b0",
		Path:   "/",
		Domain: ".zongheng.com",
		MaxAge: 2000,
	})
	req.AddCookie(&http.Cookie{
		Name:   "v_user",
		Value:  "%7Chttp%3A%2F%2Fbook.zongheng.com%2Fstore%2Fc1%2Fc0%2Fb0%2Fu0%2Fp1%2Fv9%2Fs9%2Ft0%2Fu0%2Fi1%2FALL.html%7C20021342",
		Path:   "/",
		Domain: ".zongheng.com",
	})
	req.AddCookie(&http.Cookie{
		Name:   "ver",
		Value:  "2018",
		Path:   "/",
		Domain: ".zongheng.com",
	})
	req.AddCookie(&http.Cookie{
		Name:   "PassportCaptchaId",
		Value:  "b8efb2b37f9fd11d72b3a6e903c98fa3",
		Path:   "/",
		Domain: ".zongheng.com",
	})
	req.AddCookie(&http.Cookie{
		Name:   "ZHID",
		Value:  "F0F278E1774045C56A2ED70A002A21E6",
		Path:   "/",
		Domain: ".zongheng.com",
	})

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println(resp.Request.URL.String())
	//document, err := goquery.NewDocumentFromReader(resp.Body)
	//if err != nil {
	//	panic(err)
	//}
	//document.Find(".nums>span").Filter("i").Each(func(i int, selection *goquery.Selection) {
	//fmt.Println(document.Html())
	//})

}
