package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/PuerkitoBio/goquery"
)

type ReadItem struct {
	gorm.Model
	URL   string
	Title string
}

func addUlr(w http.ResponseWriter, r *http.Request) {

	url := r.FormValue("url")

	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var title string

	// Find the review items
	doc.Find("title").Each(func(i int, s *goquery.Selection) {
		if title == "" {
			title = s.Text()
		}
	})

	fmt.Println(title)

	addDBItem(url, title)

	return
}

func addDBItem(url string, title string) {
	db, err := gorm.Open("sqlite3", "item.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&ReadItem{})

	// Create
	db.Create(&ReadItem{URL: url, Title: title})
}

func main() {
	fmt.Println("Start server")

	http.HandleFunc("/add_url", addUlr)

	http.ListenAndServe(":9001", nil)
}
