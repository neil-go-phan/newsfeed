package crawl

import (
	"context"
	"crawler/entities"
	"crawler/helpers"
	"fmt"
	"sync"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	log "github.com/sirupsen/logrus"
)

func CrawlWithChromedp(crawler entities.Crawler) ([]entities.Article, error) {
	var articles []entities.Article

	ctx, cancel, err := GetDocWithChromedp(crawler.SourceLink)
	defer cancel()
	if err != nil {
		return articles, nil
	}

	articles, err = chromedpCrawlArticles(ctx, crawler)
	if err != nil {
		return articles, nil
	}

	return articles, nil
}

func chromedpCrawlArticles(ctx context.Context, crawler entities.Crawler) ([]entities.Article, error) {
	articles := make([]entities.Article, 0)
	articleNodes, err := getArticleList(ctx, crawler)
	if err != nil {
		return articles, err
	}
	
	for _, node := range articleNodes {
		articles = append(articles, chromedpCrawlOneArticle(ctx, crawler, node))
	}
	return articles, nil
}

func getArticleList(ctx context.Context, crawler entities.Crawler) ([]*cdp.Node, error) {
	var articleNodes []*cdp.Node

	err := chromedp.Run(ctx, chromedp.Nodes(fmt.Sprintf(`//*[@class='%s']`, crawler.ArticleDiv), &articleNodes))

	if err != nil {
		return articleNodes, err
	}

	return articleNodes, nil
}

func chromedpCrawlOneArticle(ctx context.Context, crawler entities.Crawler, node *cdp.Node) (entities.Article) {
	var article entities.Article
	// title
	crawlErr := make(chan error, 3)
	var wg sync.WaitGroup

	wg.Add(1)
	go func(article *entities.Article) {
		crawlErr <- chromedpCrawlTitle(article, ctx, crawler, node)
		wg.Done()
	}(&article)

	// Description

	wg.Add(1)
	go func(article *entities.Article) {
		crawlErr <- chromedpCrawlDescription(article, ctx, crawler, node)
		defer wg.Done()
	}(&article)

	wg.Add(1)
	go func(article *entities.Article) {
		crawlErr <- chromedpCrawlAuthors(article, ctx, crawler, node)
		defer wg.Done()
	}(&article)

	// link
	// there are only one node
	wg.Add(1)
	go func(article *entities.Article) {
		defer wg.Done()
		crawlErr <- chromedpCrawlLink(article, ctx, crawler, node)
	}(&article)

	done := make(chan bool)
	go func(done chan bool) {
		for err := range crawlErr {
			if err != nil {
				log.Printf("error occurs while crawl: %s", <-crawlErr)
			}
		}
		done <- true
	}(done)

	wg.Wait()
	close(crawlErr)
	<-done
	close(done)
	return article
}

func chromedpCrawlTitle(article *entities.Article, ctx context.Context, crawler entities.Crawler, node *cdp.Node) error {
	titleCtx, cancelTitle := context.WithTimeout(ctx, 1*time.Second)
	defer cancelTitle()
	titleQuery := fmt.Sprintf(`%s//*[@class='%s']`, node.FullXPath(), crawler.ArticleTitle)
	err := chromedp.Run(titleCtx,
		chromedp.Text(titleQuery, &article.Title))
	if err != nil {
		return err
	}
	return nil
}

func chromedpCrawlAuthors(article *entities.Article, ctx context.Context, crawler entities.Crawler, node *cdp.Node) error {
	taskCtx, cancel:= context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	query := fmt.Sprintf(`%s//*[@class='%s']`, node.FullXPath(), crawler.ArticleAuthors)
	err := chromedp.Run(taskCtx,
		chromedp.Text(query, &article.Authors))
	if err != nil {
		return err
	}
	return nil
}

func chromedpCrawlDescription(article *entities.Article, ctx context.Context, crawler entities.Crawler, node *cdp.Node) error {
	descriptionCtx, cancelDescription := context.WithTimeout(ctx, 1*time.Second)
	defer cancelDescription()
	descriptionQuery := fmt.Sprintf(`%s//*[@class='%s']`, node.FullXPath(), crawler.ArticleDescription)
	err := chromedp.Run(descriptionCtx, chromedp.Text(descriptionQuery, &article.Description))
	if err != nil {
		return err
	}
	return nil
}

func chromedpCrawlLink(article *entities.Article, ctx context.Context, crawler entities.Crawler, node *cdp.Node) error {
	linkCtx, cancelLink := context.WithTimeout(ctx, 1*time.Second)
	defer cancelLink()
	var linkNodes []*cdp.Node
	linkNodesQuery := fmt.Sprintf(`%s//*[@class='%s']`, node.FullXPath(), crawler.ArticleLink)
	err := chromedp.Run(linkCtx, chromedp.Nodes(linkNodesQuery, &linkNodes))
	if err != nil {
		return err
	}

	for _, linkNode := range linkNodes {
		link, ok := linkNode.Attribute("href")
		if ok {
			article.Link = helpers.FormatLink(link, crawler.ArticleLink)
		} else {
			var linkChilds []*cdp.Node
			err = chromedp.Run(linkCtx, chromedp.Nodes(fmt.Sprintf(`%s//*`, linkNode.FullXPath()), &linkChilds))
			if err != nil {
				continue
			}
			for _, child := range linkChilds {
				linkChild, ok := child.Attribute("href")
				if ok {
					article.Link = helpers.FormatLink(linkChild, crawler.ArticleLink)
					break
				}
			}
		}
	}
	if err != nil {
		return err
	}
	return nil
}