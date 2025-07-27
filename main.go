package main

import (
	"fmt"
	"github.com/johnsonoklii/crawler/collect"
)

func main() {
	url := "https://book.douban.com/subject/37339619/"
	var f collect.Fetcher = collect.BrowserFetch{}
	body, err := f.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))

	//doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	//if err != nil {
	//	fmt.Println("new document failed:%v", err)
	//}
	//
	//doc.Find("div ul li a[target=_blank]").Each(
	//	func(i int, s *goquery.Selection) {
	//		title := s.Text()
	//		fmt.Printf("%d:%s\n", i, title)
	//	})
}
