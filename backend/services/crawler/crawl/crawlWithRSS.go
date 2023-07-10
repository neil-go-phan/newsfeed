package crawl

import (
	"context"
	"crawler/entities"
	"fmt"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/mmcdole/gofeed"
	log "github.com/sirupsen/logrus"
)

const RSS_XPATH_QUERY = "link[type='application/rss+xml']"

func TestGetRSSFeed(crawler entities.Crawler) (*gofeed.Feed, string, error) {
	feed, err := ParseRSS(crawler.SourceLink)
	if err == nil {
		return feed, "", nil
	}

	link, ok := GetRSSLink(crawler.SourceLink)
	if !ok {
		return nil, "", fmt.Errorf("not found rss link")
	}
	log.Printf("export rss link success: %s", link)
	feed, err = ParseRSS(link)
	if err != nil {
		return nil, "", err
	}
	return feed, link, nil
}

func GetRSSFeed(feedLink string) (*gofeed.Feed, error) {
	feed, err := ParseRSS(feedLink)
	if err == nil {
		return feed, nil
	}

	link, ok := GetRSSLink(feedLink)
	if !ok {
		return nil, fmt.Errorf("not found rss link")
	}
	log.Printf("export rss link success: %s", link)
	feed, err = ParseRSS(link)
	if err != nil {
		return nil, err
	}
	return feed, nil
}

func ParseRSS(url string) (*gofeed.Feed, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err != nil {
		return nil, err
	}
	return feed, nil
}

func GetRSSLink(url string) (string, bool) {
	log.Printf("%s is not a feed link, try goquery to export feed link...", url)
	link, ok := GetRSSWithGoquery(url)
	if ok {
		return link, true
	}
	log.Println("can not export with goquery, try chromedp...")
	link, ok = GetRSSWithChromedp(url)
	return link, ok
}

func GetRSSWithGoquery(url string) (string, bool) {
	doc, err := GetDocWithGoquery(url)
	if err != nil {
		log.Error("error occrus while trying to get html by goquery ", err)
		return "", false
	}

	link, ok := ExportRSSLinkGoquery(doc)
	if !ok {
		return "", false
	}
	return link, true
}

func ExportRSSLinkGoquery(doc *goquery.Document) (string, bool) {
	link, ok := doc.Find(RSS_XPATH_QUERY).First().Attr("href")
	if !ok {
		return "", ok
	}
	return link, ok
}

func GetRSSWithChromedp(url string) (string, bool) {
	ctx, cancel, err := GetDocWithChromedp(url)
	defer cancel()
	if err != nil {
		log.Error(err)
		return "", false
	}
	log.Println("chromedp navigate to url success")
	link, ok := ExportRSSLinkChromedp(ctx)
	if !ok {
		return "", false
	}
	return link, true
}

func ExportRSSLinkChromedp(ctx context.Context) (string, bool) {
	rssNodes, err := getRSSNodes(ctx)
	if err != nil {
		log.Error("error occrus while getRSSLink with chromedp ", err)
		return "", false
	}
	log.Println("detect rss node ", len(rssNodes))
	// most of time it is only one node
	// sometime it was two or three. But we will only get the first node
	// cause this is a test. if the user doesn't like the result then they can manually enter the rss link
	link, ok := exportRSSLinkFromNode(rssNodes[0])
	return link, ok
}

func getRSSNodes(ctx context.Context) ([]*cdp.Node, error) {
	var rssNodes []*cdp.Node
	findNodeCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	err := chromedp.Run(findNodeCtx, chromedp.Nodes(RSS_XPATH_QUERY, &rssNodes))
	if err != nil {
		return rssNodes, err
	}
	return rssNodes, nil
}

func exportRSSLinkFromNode(node *cdp.Node) (string, bool) {
	link, ok := node.Attribute("href")
	return link, ok
}
