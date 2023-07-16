package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"

	"github.com/joho/godotenv"
	twitterscraper "github.com/n0madic/twitter-scraper"
)

func main() {
	godotenv.Load()

	scraper := twitterscraper.New()
	scraper.LoginOpenAccount()
	profile, err := scraper.GetProfile(os.Getenv("TWITTER_USERNAME"))
	if err != nil {
		panic(err)
	}
	bio := profile.Biography
	re := regexp.MustCompile(`(https?://[^\s]+)`)
	url := re.FindString(bio)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, url, http.StatusMovedPermanently)
	})

	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
