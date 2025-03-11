package common

import (
	"log"
	"net/http"
	"strings"
)


func IsImage(url string) bool {
	// Create the request
	req, _ := http.NewRequest("HEAD", url, nil)

	// Set the User-Agent header
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Check the Content-Type header
	contentType := resp.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "image/") {
		return true
	} else {
		return false
	}
}

func CheckURL(string) {

}
