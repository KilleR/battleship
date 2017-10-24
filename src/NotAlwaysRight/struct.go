package main

type story struct {
	Id string
	Title string
	Location string
	Tags []string
	Body []storyParagraph
	Next string
	Prev string
}

type storyParagraph struct {
	Content []storyText
}

type storyText struct {
	Formatting string
	Text string
}