package getdata

import (
	"encoding/base64"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ExampleScrape get telegarm v2list page data
func ExampleScrape(count string, cors bool) (string, bool) {
	// Request the HTML page.
	var c int
	var err error
	c, err = strconv.Atoi(count)
	var url string
	url = "https://t.me/s/V2List"
	if cors {
		url = strings.Join([]string{"https://cors.zme.ink", url}, "/")
	}
	// fmt.Println(url)
	res, err := http.Get(url)
	if err != nil {
		return "bad", false
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		// log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		return res.Status, false
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	root := doc.Find("body.widget_frame_base > main.tgme_main > div.tgme_container > section.tgme_channel_history > div.tgme_widget_message_wrap")
	length := root.Length()
	findthis := root.Eq(length - c).Find("div.tgme_widget_message_text").Text()
	return findthis, true
}

// MakeList use split to make a array for string
func MakeList(d string) []string {
	x := []string{}
	l := strings.Split(d, "vmess://")
	for _, item := range l {
		var l int
		l = len(item)
		if l > 0 {
			var strHaiCoder string
			var newstr string
			var v string

			decodeBytes, err := base64.StdEncoding.DecodeString(item)
			if err != nil {
				return x
			}
			strHaiCoder = `"ps" :"翻墙党fanqiangdang.com","" :`
			reg := regexp.MustCompile(strHaiCoder)
			newstr = reg.ReplaceAllString(string(decodeBytes), `"ps" :`)
			var strtobyte []byte = []byte(newstr)
			v = strings.Join([]string{"vmess:", base64.StdEncoding.EncodeToString(strtobyte)}, "//")
			x = append(x, v)
		}
	}
	return x
}

// MakeData is a make Array to BASE64 string function
func MakeData(d []string) string {
	var data string = strings.Join(d[:], "\n")
	var strtobyte []byte = []byte(data)
	return base64.StdEncoding.EncodeToString(strtobyte)
}

// Start this
func Start(n string, w bool) string {
	var d []string
	var dd string = ""
	data, status := ExampleScrape(n, w)
	if status {
		d = MakeList(data)
		dd = MakeData(d)
	}
	return dd
}
