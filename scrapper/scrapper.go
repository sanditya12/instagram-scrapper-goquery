package scrapper

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

//GetDateTime return the datetime in ISO 8601 time
func GetDateTime(link string) {
	res, err := http.Get(link)
	fmt.Println("requesting ", link)
	checkErr(err)
	checkCode(res)
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	uploadTime := doc.Find("body")
	fmt.Println(uploadTime.Html())
	// uploadTime.Each(func(i int, s *goquery.Selection) {
	// 	fmt.Println(s.Html())
	// })
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
}
