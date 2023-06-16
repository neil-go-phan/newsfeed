package helpers

import "crawler/entities"

func CheckEmptyArticles(articles []entities.Article) bool {
	var count int
	for _, article := range articles {
		if article.Title == "" && article.Description == "" && article.Link == "" {
			count++
		}
	}
	return count == len(articles)
}