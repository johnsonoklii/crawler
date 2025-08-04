package main

import (
	"fmt"
	"github.com/johnsonoklii/crawler/collect"
	"github.com/johnsonoklii/crawler/log"
	"go.uber.org/zap/zapcore"
	"time"
)

func main() {
	plugin, c := log.NewFilePlugin("./log.txt", zapcore.InfoLevel)
	defer c.Close()
	logger := log.NewLogger(plugin)
	logger.Info("log init end")

	url := "https://google.com"

	//proxyUrls := []string{"http://127.0.0.1:8888", "http://127.0.0.1:8889"}
	//Proxy, err := proxy.RoundRobinProxySwitcher(proxyUrls...)
	//if err != nil {
	//	fmt.Println("RoundRobinProxySwitcher failed")
	//}

	var f collect.Fetcher = collect.BrowserFetch{
		Timeout: 5000 * time.Millisecond,
		//Proxy:   Proxy,
	}
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
