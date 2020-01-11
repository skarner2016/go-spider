package lagou

import (
	"bytes"
	"errors"
	"fmt"
	"go-spider/pkg/common"
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var languages = []string{"php", "go"}

var citys = map[string]string{
	"chengdu": "252",
}

func Start() {
	prefix := "https://www.lagou.com/jobs/list_"
	//url := "PHP/p-city_252?px=default#filterBox"

	suffix := "?px=default#filterBox"

	for _, language := range languages {
		for city, cityNum := range citys {
			// 组装 url
			buffer := bytes.Buffer{}
			buffer.WriteString(prefix)
			buffer.WriteString(strings.ToUpper(language))
			buffer.WriteString("/p-city_")
			buffer.WriteString(cityNum)
			buffer.WriteString(suffix)

			url := buffer.String()

			Handle(url, city, language)

			break
		}
	}

	//Parse()
}

func Handle(url, city, language string) {
	pageNum, err := ParsePage(url)
	if err != nil {
		log.Fatal(err)
	}
	if pageNum < 1 {
		log.Fatal(errors.New("total page error"))
	}
	for i := 1; i <= pageNum; i++ {

		Parse(url, city, language)
		fmt.Println(i)
	}
}

func ParsePage(url string) (int, error) {
	doc, err := common.GetDoc(map[string]string{}, url)
	if err != nil {
		log.Fatal(err)
	}

	totalNum := doc.Find("div.page-number span.totalNum").Text()

	Num, err := strconv.Atoi(totalNum)
	if err != nil {
		return 0, err
	}

	return Num, nil
}

func Parse(url, city, language string) {
	doc, err := common.GetDoc(map[string]string{}, url)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("ul.item_con_list li.con_list_item").Each(func(i int, s *goquery.Selection) {

		position := s.Find("div.list_item_top .position")

		positionName := position.Find(".p_top .position_link h3").Text()
		positionLocation := position.Find(".p_top .position_link span.add em").Text()
		formatTime := position.Find(".p_top .format-time").Text()
		money := position.Find(".p_bot div.li_b_l span.money").Text()

		company := s.Find("div.list_item_top .company")
		companyName := company.Find(".company_name a").Text()
		companyItem := company.Find(".industry").Text()
		companyItem = strings.Replace(companyItem, " ", "", -1)

	})
}
