package crawl

import (
	"crawler/entities"
	"crawler/helpers"
	
	"github.com/PuerkitoBio/goquery"
)

func CrawlWithGoQuery(crawler entities.Crawler) ([]entities.Article, error) {
	var articles []entities.Article

	doc, err := GetDocWithGoquery(crawler.SourceLink)
	if err != nil {
		return articles, nil
	}

	crawledArticles := goqueryCrawl(doc, crawler)
	articles = append(articles, crawledArticles...)

	return articles, nil
}

func goqueryCrawl(doc *goquery.Document, crawler entities.Crawler) ([]entities.Article) {
	var articles []entities.Article
	doc.Find(helpers.FormatClassName(crawler.ArticleDiv)).Each(func(i int, s *goquery.Selection) {
		articles = append(articles, goqueryCrawlOneArticle(crawler, s)) 
	})

	return articles
}

func goqueryCrawlOneArticle(crawler entities.Crawler, s *goquery.Selection) entities.Article {
	var article entities.Article
	article.Title = s.Find(helpers.FormatClassName(crawler.ArticleTitle)).Text()
	article.Description = s.Find(helpers.FormatClassName(crawler.ArticleDescription)).Text()
	article.Link = goqueryCrawlLink(crawler, s)
	article.Authors = s.Find(helpers.FormatClassName(crawler.ArticleAuthors)).Text()
	return article
}

func goqueryCrawlLink(crawler entities.Crawler, s *goquery.Selection) string {
	var link = ""
	href, ok := s.Find(helpers.FormatClassName(crawler.ArticleLink)).Attr("href")
	if ok {
		link = helpers.FormatLink(href, crawler.SourceLink)
	} else {
		s.Find(helpers.FormatClassName(crawler.ArticleLink)).Each(func(i int, s *goquery.Selection) {
			href, ok := s.Find("a").Attr("href")
			if ok {
				link = helpers.FormatLink(href, crawler.SourceLink)
			}
		})
	}
	return link
}
