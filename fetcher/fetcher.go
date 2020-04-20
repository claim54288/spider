package fetcher

import (
	"bufio"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"io"
	"log"
	"net/http"
	"time"
)

var limiter = time.NewTicker(200 * time.Millisecond)

//页面信息拉取
func Fetch(url string) (*goquery.Document, error) {
	//<-limiter.C
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("wrong status code:%d", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36")

	req.AddCookie(&http.Cookie{
		Name:   "AST",
		Value:  "15873666657876ecb2312f7",
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
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code:%d", resp.StatusCode)
	}

	content, err := goquery.NewDocumentFromReader(resp.Body)
	content.Url = resp.Request.URL
	return content, err

}

//识别是什么编码的网页
func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		log.Printf("Fetcher error:%v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
