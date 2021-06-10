package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"03_bing/metadata"

	"github.com/PuerkitoBio/goquery"
)

// At last check, this isn't functioning properly.  Still playing to figure out how to get to work
// as bing and libraries have probably changed since BHG was released.

func handler(i int, s *goquery.Selection) {
	url, ok := s.Find("a").Attr("href")
	log.Printf(url)
	if !ok {
		return
	}

	fmt.Printf("%d: %s\n", i, url)
	res, err := http.Get(url)
	if err != nil {
		return
	}

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	defer res.Body.Close()

	r, err := zip.NewReader(bytes.NewReader(buf), int64(len(buf)))
	if err != nil {
		return
	}

	cp, ap, err := metadata.NewProperties(r)
	if err != nil {
		return
	}

	log.Printf(
		"%25s %25s - %s %s\n",
		cp.Creator,
		cp.LastModifiedBy,
		ap.Application,
		ap.GetMajorVersion())
}

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("Missing required argument. Usage: main.go <domain> <ext>")
	}
	domain := os.Args[1]
	filetype := os.Args[2]

	q := fmt.Sprintf(
		"site:%s && filetype:%s && instreamset:(url title):%s",
		domain,
		filetype,
		filetype)

	search := fmt.Sprintf("http://www.bing.com/search?q=%s", url.QueryEscape(q))
	res, err := http.Get(search)
	if err != nil {
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Panicln(err)
	}
	defer res.Body.Close()
	s := "html body.b_sbText div#b_content main ol#b_results li.b_algo div.b_title h2"

	doc.Find(s).Each(handler)
}
