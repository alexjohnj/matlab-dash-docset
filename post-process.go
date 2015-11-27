package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
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

	doc.Find(".sticky_header_container").Each(func(i int, s *goquery.Selection) {
		s.Remove()
	})
	// doc.Find("#sidebar").First().Remove()
	buildSectionTOC(doc)
	buildInputTOC(doc)
	buildOutputTOC(doc)

	htmlContent, _ := doc.Html()
	file.Seek(0, 0)
	_, err = file.WriteString(htmlContent)
	if err != nil {
		log.Fatal(err)
	}
	file.Sync()
	return nil
}

func buildInputTOC(doc *goquery.Document) {
	doc.Find(".input_argument_container .argument_name code").Each(func(i int, s *goquery.Selection) {
		inputName := s.Text()
		inputLinkElement := "<a name=\"//apple_ref/cpp/Parameter/"
		inputLinkElement += strings.Replace(url.QueryEscape(inputName), "+", "%20", -1)
		inputLinkElement += "\" class=\"dashAnchor\"></a>"
		s.WrapHtml(inputLinkElement)
	})
}

func buildOutputTOC(doc *goquery.Document) {
	doc.Find(".output_argument_container .argument_name code").Each(func(i int, s *goquery.Selection) {
		outputName := s.Text()
		outputLinkElement := "<a name=\"//apple_ref/cpp/Value/"
		outputLinkElement += strings.Replace(url.QueryEscape(outputName), "+", "%20", -1)
		outputLinkElement += "\" class=\"dashAnchor\"></a>"
		s.WrapHtml(outputLinkElement)
	})
}

func buildSectionTOC(doc *goquery.Document) {
	doc.Find("#doc_center_content h2").Each(func(i int, s *goquery.Selection) {
		sectionName := s.Text()
		sectionElement := "<a name=\"//apple_ref/cpp/Section/"
		sectionElement += strings.Replace(url.QueryEscape(sectionName), "+", "%20", -1)
		sectionElement += "\" class=\"dashAnchor\"></a>"
		s.WrapHtml(sectionElement)
	})
}
