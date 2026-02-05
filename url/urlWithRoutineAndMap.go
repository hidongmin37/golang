package url

import (
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

func main() {
	results := make(map[string]string)
	c := make(chan requestResult)

	urls := []string{
		"https://www.airbnbd.com/", // FAILED
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}
	for _, url := range urls {
		go hitURL3(url, c)
	}
	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}
	for url, status := range results {
		fmt.Println("URL:", url, "Status:", status)
	}
}

// <- with parameter는 보내는 것 만 된다.
func hitURL3(url string, c chan<- requestResult) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- requestResult{url: url, status: status}
}
