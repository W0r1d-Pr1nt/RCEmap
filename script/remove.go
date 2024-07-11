package script

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// 使用 goquery 库的 Selection.Text() 方法获取去除了 HTML 标签的纯文本

func RemoveHTMLTags(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatal(err)
	}

	return doc.Text()
}
