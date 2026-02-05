package url

import (
	"fmt"
	"net/http"
)

type result struct {
	url    string
	status string
}

func main() {
	c := make(chan result)

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
		go hitURL2(url, c)
	}
	for i := 0; i < len(urls); i++ {
		res := <-c
		fmt.Println(res.url, res.status)
	}

}

// <- with parameter는 보내는 것 만 된다.
func hitURL2(url string, c chan<- result) {
	fmt.Println("Checking URL", url)
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- result{url: url, status: status}
}
