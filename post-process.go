package main

import(
	"github.com/PuerkitoBio/goquery"
	"os"
	"path/filepath"
	"log"
	"net/url"
)

func main() {
	rootDir := "./matlab.docset/Contents/Resources/Documents"

	filepath.Walk(rootDir, walkFunction)
}

func walkFunction(path string, info os.FileInfo, err error) error {
	if filepath.Ext(path) != ".html" {
		return nil
	}
	log.Println(path)
	file, err := os.OpenFile(path, os.O_RDWR, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".sticky_header_container").Each(func(i int, s *goquery.Selection){
		s.Remove()
	});
	// doc.Find("#sidebar").First().Remove()
	buildSectionTOC(doc)

	htmlContent, _ := doc.Html()
	file.Seek(0,0)
	_, err = file.WriteString(htmlContent)
	if err != nil {
		log.Fatal(err)
	}
	file.Sync()
	return nil
}

func buildSectionTOC(doc *goquery.Document) {
	doc.Find("#doc_center_content h2").Each(func(i int, s *goquery.Selection){
		sectionName := s.Text()
		sectionElement := "<a name=\"//apple_ref/cpp/Section/" + url.QueryEscape(sectionName) + "\" class=\"dashAnchor\"></a>"
		s.WrapHtml(sectionElement)
	})
}
