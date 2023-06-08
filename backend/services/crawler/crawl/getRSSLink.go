package crawl

import (
	"context"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	log "github.com/sirupsen/logrus"
)

const RSS_XPATH_QUERY = "link[type='application/rss+xml']"

func GetRSSLinkGoqueryDoc(doc *goquery.Document) (string, bool) {
	link, ok := doc.Find(RSS_XPATH_QUERY).First().Attr("href")
	if !ok {
		return "", ok
	}
	return link, ok
}

func GetRSSLinkChromedp(ctx context.Context) (string, bool) {
	rssNodes, err := getRSSNodes(ctx)
	if err != nil {
		log.Error("error occrus while getRSSLink with chromedp ", err)
		return "", false
	}
	// most of time it is only one node
	// sometime it was two or three. But we will only get the first node
	// cause this is a test. if the user doesn't like the result then they can manually enter the rss link
	link, ok := exportRSSLinkFromNode(rssNodes[0])
	return link, ok
}

func getRSSNodes(ctx context.Context) ([]*cdp.Node, error) {
	var rssNodes []*cdp.Node
	err := chromedp.Run(ctx, chromedp.Nodes(RSS_XPATH_QUERY, &rssNodes))
	if err != nil {
		return rssNodes, err
	}
	return rssNodes, nil
}

func exportRSSLinkFromNode(node *cdp.Node) (string, bool){
	link, ok := node.Attribute("href")
	return link, ok
}