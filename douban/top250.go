package douban

import (
	"fmt"
	"go-spider/pkg/common"
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func StartTop250() {
	t1 := time.Now() // get current time
	parse()
	elapsed := time.Since(t1)
	fmt.Println("爬虫结束,总共耗时: ", elapsed)
}

func parse() {
	url := "https://movie.douban.com/top250?start=" + "50" + "&filter="
	response, err := common.SendRequest(map[string]string{}, url)
	if err != nil {
		log.Fatal(err)
	}

	if response.StatusCode != http.StatusOK {
		log.Fatalf("status code err: %d %s", response.StatusCode, response.Status)
	}
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("#content div.grid-16-8 div.article ol li").Each(func(i int, s *goquery.Selection) {
		name := s.Find("div.item .info  span.title:nth-child(1)").Text()
		fmt.Println("name: ", name)

		star := s.Find("div.item div.info div.bd .star .rating_num").Text()
		fmt.Println("star: ", star)

		number := s.Find("div.item div.info div.bd .star span:nth-child(4)").Text()
		fmt.Println("number: ", number)

		desc := s.Find("div.item div.info div.bd .quote span.inq").Text()
		fmt.Println("desc: ", desc)
	})
}
