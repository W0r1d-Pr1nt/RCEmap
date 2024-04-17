package script

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func RemoveHTMLTags(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatal(err)
	}

	return doc.Text()
}
