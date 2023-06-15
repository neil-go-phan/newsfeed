package helpers

import (
	"fmt"
	"net/url"
	"strings"
)

func FormatClassName(class string) string {
	var classes string
	hashParts := strings.Split(class, " ")
	for _, s := range hashParts {
		classes = classes + "." + s
	}
	return classes
}

func FormatLink(link string, urlIn string) string {
	// xử lý trường hợp web để link kiểu /hello/halu thay vì example.com/hello/halu
	linkFormated := link
	url, err := url.ParseRequestURI(urlIn)
	if err != nil {
		return ""
	}
	hostname := strings.TrimPrefix(url.Hostname(), "www.")
	if link != "" {
		ok := strings.Contains(linkFormated, hostname)
		if !ok {
			if err != nil {
				return fmt.Sprintf("%s%s", urlIn, linkFormated)
			}
			linkFormated = fmt.Sprintf("https://%s%s", url.Hostname(), linkFormated)
		}
	}

	return linkFormated
}
