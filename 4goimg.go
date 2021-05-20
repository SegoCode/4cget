package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"sync"
)

func findImages(html string) []string {
	imgs := imgRE.FindAllStringSubmatch(html, -1)
	out := make([]string, len(imgs))
	for i := range out {
		out[i] = imgs[i][1]
	}
	return out
}

func downloadFile(wg *sync.WaitGroup, url string, fileName string) {

	resp, _ := http.Get(url)
	if resp.StatusCode == 404 {
		url = strings.Replace(url, ".jpg", ".png", 1)
		fileName = strings.Replace(fileName, ".jpg", ".png", 1)
		resp, _ = http.Get(url)
		if resp.StatusCode == 404 {
			url = strings.Replace(url, ".pgn", ".gif", 1)
			fileName = strings.Replace(fileName, ".png", ".gif", 1)
			resp, _ = http.Get(url)
		}
	}
	defer resp.Body.Close()

	if resp.StatusCode != 404 {
		img, _ := os.Create(fileName)
		defer img.Close()

		b, _ := io.Copy(img, resp.Body)
		fmt.Println("File downloaded: "+fileName+" - Size (Bytes):", b)
	}
	wg.Done()
}


func main() {
	var imgRE = regexp.MustCompile(`<img[^>]+\bsrc=["']([^"']+)["']`)
	var re = regexp.MustCompile("[0-9]+")
	var wg sync.WaitGroup
	var inputUrl string
	var linkImg string
	var nameImg string

	//Usage validation
	if len(os.Args) <= 1 {
		fmt.Println("[!] USAGE: 4goimg https://boards.4channel.org/w/thread/.../...")
		os.Exit(1)
	}

	//input url validation
	inputUrl = os.Args[1]
	_, errParse := url.ParseRequestURI(inputUrl)
	if errParse != nil {
		fmt.Println("[!] URL NOT VALID (Example: https://boards.4channel.org/w/thread/.../...)")
		os.Exit(1)
	}

	fmt.Println(`
  _  _                              
 | || |                           
 | || |_ __ _  ___  _ _ __ ___   __ _ 
 |__   _/ _' |/ _ \| | '_ ' _ \ / _' |
    | || (_| | (_) | | | | | | | (_| |
    |_| \__. |\___/|_|_| |_| |_|\__. |
         __/ |                   __/ |
        |___/                   |___/ 
		
               [ github.com/SegoCode ]` + "\n")

	fmt.Println("[*] DOWNLOAD STARTED (" + inputUrl + ")\n")

	resp, _ := http.Get(inputUrl)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("[!] CONNECTION ERROR")
	}

	for _, each := range findImages(string(body)) {
		linkImg = "http:" + strings.Replace(each, "s.jpg", ".jpg", 1)
		nameImg = re.FindAllString(linkImg, -1)[1] + ".jpg"
		wg.Add(1)
		go downloadFile(&wg, linkImg, nameImg)
	}

	wg.Wait()
	fmt.Println("\n" + "[*] DOWNLOAD COMPLETE [*]")

}
