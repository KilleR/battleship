package main

import (
	"net/http"
	"errors"
	"strconv"
	"io/ioutil"
	"golang.org/x/net/html"
	"strings"
	"fmt"
)



func getPage(url string) ([]byte, error) {
	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.New("Page returned status "+strconv.Itoa(res.StatusCode)+" "+res.Status)
	}

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	res.Body.Close()

	return content, nil
}

func parsePage(url string) ([]story, error) {
	var (
		stories []story
		currentStory story
		currentParagraph storyParagraph
		storyLocation string // track where we are in the story, so we can get the text content
		formatting string // track current formatting
		nextPrev string // track if we're looking at a next or prev button
	)

	fmt.Println("Fetching:",url)

	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	z := html.NewTokenizer(res.Body)

	depth := 0
	sectionDepth := 0

	parseLoop:
	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			break parseLoop
		case tt == html.StartTagToken:
			t := z.Token()

			// if we're already searching a story, keep digging
			if depth > 0 {
				depth++
				if sectionDepth > 0 {
					sectionDepth++
				}
				switch t.Data {
				case "h1":
					// check to see if we've got a story header
					for _, a := range t.Attr {
						if a.Key == "class" {
							// got class, check if it's a story header
							if a.Val == "storytitle" {
								storyLocation = "title"
								sectionDepth++
							}
							break
						}
					}
				case "div":
					for _, a := range t.Attr {
						if a.Key == "class" {
							// got class, check if it's a story header
							if a.Val == "post_header" {
								storyLocation = "header"
								sectionDepth++
							}
							if a.Val == "storycontent" {
								storyLocation = "content"
								sectionDepth++
							}
							if a.Val == "nextprevtitles" {
								storyLocation = "nextprev"
								sectionDepth++
							}
						}

						if a.Key == "style" {
							if a.Val == "float: left;" {
								if storyLocation == "nextprev" {
									nextPrev = "prev"
								}
							}
							if a.Val == "float: right;" {
								if storyLocation == "nextprev" {
									nextPrev = "next"
								}
							}
						}
					}
				case "p":
					currentParagraph = storyParagraph{}
				case "a":
					if storyLocation == "title" {
						// if we're in the title, look for the title link
						for _, a := range t.Attr {
							if a.Key == "href" {
								currentStory.Location = a.Val
								break
							}
						}
					}
					if storyLocation == "nextprev" {
						// if we're in the next/prev section, look for the links
						for _, a := range t.Attr {
							if a.Key == "href" {
								if nextPrev == "prev" {
									currentStory.Prev = urlToPostId(a.Val)
								}
								if nextPrev == "next" {
									currentStory.Next = urlToPostId(a.Val)
								}
								break
							}
						}
					}
				case "i":
					formatting = "i"
				case "b":
					formatting = "b"
				}
			} else {
				// only look for new post data if we're not already digging
				isDiv := t.Data == "div"
				if isDiv {
					// check to see if it is a post
					for _, a := range t.Attr {
						if a.Key == "class" {
							// got class, check if it's a post
							if a.Val == "post" {
								depth++ // start digging
								currentStory = story{}
							}
						}

						if a.Key == "id" && currentStory.Id == "" {
							currentStory.Id = a.Val
						}
					}
				}
			}
		case tt == html.EndTagToken:
			if depth > 0 {
				depth--
				if sectionDepth > 0 {
					sectionDepth--
					if sectionDepth == 0 {
						//fmt.Println("END OF SECTION")
						storyLocation = ""
					}
				}
				if depth == 0 {
					//fmt.Println("END OF POST")
					stories = append(stories, currentStory)
				}

				t := z.Token()
				if formatting == t.Data {
					formatting = ""
				}

				switch t.Data{
					case "p":
						currentStory.Body = append(currentStory.Body, currentParagraph)
				}
			}
		case tt == html.TextToken:
			if depth > 0 {
				data := strings.Trim(z.Token().Data, " ,|")
				//fmt.Println(depth, "Text token in (",storyLocation,"):", data)
				switch storyLocation {
				case "title" :
					currentStory.Title += data
				case "header":
					data = strings.Trim(data, "\n\t")
					if data != "" {
						currentStory.Tags = append(currentStory.Tags, data)
					}
				case "content":
					currentParagraph.Content = append(currentParagraph.Content, storyText{formatting, data})
				}
			}
		}
	}

	return stories, nil
}