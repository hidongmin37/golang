package url

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

var baseUrl string = "https://www.saramin.co.kr/zf_user/jobs/list/job-category?cat_kewd=84&loc_mcd=101000&exp_cd=2&exp_max=2&search_optional_item=y&search_done=y&panel_count=y&preview=y&isAjaxRequest=0&page_count=50&sort=RL&type=job-category&is_param=1&isSearchResultEmpty=1&isSectionHome=0&searchParamCount=4#searchTitle"
var baseOfBaseUrl string = "https://www.saramin.co.kr"

func main() {
	totalPages := getPages()
	for i := 1; i < totalPages+1; i++ {
		getPage(i)
	}

}

func getPage(page int) {
	pageNumber := "&page=" + strconv.Itoa(page)
	pageUrl := baseUrl + pageNumber
	resp, err := http.Get(pageUrl)
	checkErr(err)
	checkCode(resp)
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkErr(err)
	doc.Find("h2.job_tit > a").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")
		fmt.Println(link)
	})

}

func getPages() int {
	pages := 0
	resp, err := http.Get(baseUrl)
	checkErr(err)
	checkCode(resp)

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkErr(err)
	doc.Find(".PageBox").Each(func(i int, s *goquery.Selection) {
		pages = s.Children().Length()
	})
	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(resp *http.Response) {
	if resp.StatusCode != 200 {
		log.Fatalln("Request failed with status code: " + strconv.Itoa(resp.StatusCode))
	}
}
