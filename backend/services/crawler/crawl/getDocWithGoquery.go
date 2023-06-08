package crawl

import (
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
)

func GetDocWithGoquery(url string) (*goquery.Document, error) {
	doc := new(goquery.Document)
	log.Println("get doc with goquery url", url)

	resp, err := sendRequest(url)
	if err != nil {
		return doc, err
	}

	defer resp.Body.Close()

	doc, err = goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("can not create goquery document, err:%s\n", err)
		return doc, err
	}
	return doc, nil
}

func sendRequest(url string) (*http.Response, error){
	client := http.Client{
		Timeout: 30 * time.Minute,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Errorln("can not create when crawl HTTP:", err)
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.97 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		log.Errorln("can not do http request:", err)
		return nil, err
	}

	return resp, nil
}