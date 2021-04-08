package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/mmcdole/gofeed"
)

func makeReadme(filename string) error {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("https://blog.kaiiyer.co/rss.xml")
	if err != nil {
		log.Fatalf("error getting feed: %v", err)
	}
	// Get the latest item
	blogItem := feed.Items[0]

	rand.Seed(time.Now().UnixNano())
	whoList := []string{"🦄", "👨‍💻", "🦊", "🦔", "🦡", "🐹", "🐷"}
	who := rand.Intn(len(whoList))
	whatList := []string{"👍", "🎉", "🔥", "💕", "⭐️", "👏", "🙌"}
	what := rand.Intn(len(whatList))
	date := time.Now().Format("2 Jan 2006")

	// Whisk together static and dynamic content until stiff peaks form
	hello := "### Hey! I’m Kai Iyer 👋\n\n --- \n\n ![](https://visitor-badge.laobi.icu/badge?page_id=kaiiyer.visitor-badge)	[![Twiiter Badge](https://img.shields.io/badge/@kaiiyer-blueviolet?style=flat-square&labelColor=1ca0f1&logo=twitter&logoColor=white&link=https://twitter.com/kaiiyer)](https://twitter.com/kaiiyer)	[![Medium Badge](https://img.shields.io/badge/@kaiiyer-black?style=flat-square&labelColor=00000&logo=medium&logoColor=white&link=https://medium.com/@kaiiyer)](https://medium.com/@kaiiyer)	[![Pinterest Badge](https://img.shields.io/badge/@kai_iyer-darkred?style=flat-square&labelColor=red&logo=Pinterest&logoColor=white&link=https://www.pinterest.com/kai_iyer/)](https://www.pinterest.com/kai_iyer/) \n\n <img src='https://media.giphy.com/media/dlMIwDQAxXn1K/giphy.gif' alt='You found me!' width='40%' align='right'/> \n\n#### I'm currently exploring the ~~trends in the Cyber Market~~ Anime World"
	blog := "\n\n✨ I’m an Infosec developer :zap: I love breaking the limits with smart work. \n\n - 🌱 I attend **hackathons** and speak about trending technologies.\n\n - ⭐️ I use: `.py`, `.go`, `.c`, `.js`, `.html`, `.scss`, `.json`, `.yml` \n\n - 💜 Interests: Anime & Cyber \n\n - 😄 I'm Kai Iyer in Japan & NetworkOverclocker in Canada. \n\n - This " + whoList[who] + " says they " + whatList[what] + " my latest blog post: **[" + blogItem.Title + "](" + blogItem.Link + ")**"
	updated := "Last updated by magic on " + date + "." + "\n\n ---"
	data := fmt.Sprintf("%s\n%s\n\n%s\n", hello, blog, updated)

	// Creating new file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Bake at n bytes per second until golden brown
	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}

func main() {

	makeReadme("../README.md")

}
