package main

import (
	"fmt"
	"log"
	"encoding/json"
	"os"
	"time"
)

func init() {
	fmt.Println("NAR Init...")
}

func main() {
	start := time.Now()

	var storiesList []story
	//s, err := parsePage("http://notalwaysright.com/manufacturers-suggested-retail-conspiracy/97882/")
	//if err != nil {
	//	fmt.Println("Failed to parse post page:", err)
	//	panic("AAH!")
	//}
	//storiesList = append(storiesList, s...)
	nextStory := "97882"
	for i:=0; i < 5; i++ {
		var s []story
		var err error
		s, err = parsePage("http://notalwaysright.com/-/"+nextStory)
		nextStory = s[0].Next
		if err != nil {
			fmt.Println("Failed to parse post page:", err)
			break
		}
		storiesList = append(storiesList, s...)
	}
	//// repeat until we reach the end
	//if(currentStory.Next != "") {
	//	if len(stories) > 150 {
	//		fmt.Println("Got 5 stories, stopping for now.")
	//	} else {
	//		fmt.Println("Fetching next story:", currentStory.Next)
	//		parsePage(currentStory.Next)
	//	}
	//}

	//if err != nil {
	//	log.Fatalln("Failed to get main page:",err)
	//}

	out, err := json.MarshalIndent(storiesList, "", "  ")
	if err != nil {
		log.Fatalln("Failed to marshal output:", err)
	}

	fmt.Println(string(out))

	for _,v := range storiesList {
		//fmt.Println(v.Tags[len(v.Tags)-1])
		//fmt.Println(v.Tags[len(v.Tags)-2])
		//fmt.Println(v.Tags[len(v.Tags)-3])

		for _,c := range v.Body {
			for _,b := range c.Content {
				fmt.Print(b.Text)
			}
			fmt.Print("\n")
		}

		out, err =json.MarshalIndent(v, "", "  ")
		if err != nil {
			log.Println("Failed to marshal story:", err)
			continue
		}

		file, err := os.OpenFile(v.Id, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Println("Could not open output file:", err)
			continue
		}
		_, err = file.Write(out)
		if err != nil {
			fmt.Println("Could not write to output file:", err)
			continue
		}


	}

	fmt.Println("Done after:", time.Since(start))

}