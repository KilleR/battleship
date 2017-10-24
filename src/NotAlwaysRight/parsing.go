package main

import "regexp"

func urlToPostId(url string) string {
	rex := regexp.MustCompile(`/([0-9]+)/`)

	matches := rex.FindStringSubmatch(url)

	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}