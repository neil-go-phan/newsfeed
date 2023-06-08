package crawl

import "github.com/mmcdole/gofeed"

func ParseRSS(url string) (*gofeed.Feed, error){
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err != nil {
		return nil, err
	} 
	return feed, nil
}