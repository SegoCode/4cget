package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"sync"
)

func findImages(html string) []string {
	imgRE := regexp.MustCompile(`<img[^>]+\bsrc=["']([^"']+)["']`)
	imgs := imgRE.FindAllStringSubmatch(html, -1)
	out := make([]string, len(imgs))
	for i := range out {
		out[i] = imgs[i][1]
	}
	return out
}

func downloadFile(wg *sync.WaitGroup, url string, fileName string, path string) {

	//i know, just work
	resp, _ := http.Get(url)
	if resp.StatusCode == 404 {
		url = strings.Replace(url, ".jpg", ".png", 1)
		fileName = strings.Replace(fileName, ".jpg", ".png", 1)
		resp, _ = http.Get(url)
		if resp.StatusCode == 404 {
			url = strings.Replace(url, ".png", ".webm", 1)
			fileName = strings.Replace(fileName, ".png", ".webm", 1)
			resp, _ = http.Get(url)
			if resp.StatusCode == 404 {
				url = strings.Replace(url, ".webm", ".gif", 1)
				fileName = strings.Replace(fileName, ".webm", ".gif", 1)
				resp, _ = http.Get(url)
			}
		}
	}
	defer resp.Body.Close()

	if resp.StatusCode != 404 {

		img, _ := os.Create(path + "//" + fileName)
		defer img.Close()

		b, _ := io.Copy(img, resp.Body)
		fmt.Println("File downloaded: "+fileName+" - Size (Bytes):", b)
	}
	wg.Done()
}

func main() {
	var re = regexp.MustCompile("[0-9]+")
	var wg sync.WaitGroup
	var inputUrl string
	var linkImg string
	var nameImg string

	//Usage validation
	if len(os.Args) <= 1 {
		fmt.Println("[!] USAGE: 4cget https://boards.4channel.org/w/thread/.../...")
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
░░██╗██╗░█████╗░░██████╗░███████╗████████╗
░██╔╝██║██╔══██╗██╔════╝░██╔════╝╚══██╔══╝
██╔╝░██║██║░░╚═╝██║░░██╗░█████╗░░░░░██║░░░
███████║██║░░██╗██║░░╚██╗██╔══╝░░░░░██║░░░
╚════██║╚█████╔╝╚██████╔╝███████╗░░░██║░░░
░░░░░╚═╝░╚════╝░░╚═════╝░╚══════╝░░░╚═╝░░░
                    [ github.com/SegoCode ]` + "\n")

	fmt.Println("[*] DOWNLOAD STARTED (" + inputUrl + ") [*] \n")

	resp, _ := http.Get(inputUrl)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("[!] CONNECTION ERROR")
		os.Exit(1)
	}

	board := strings.Split(inputUrl, "/")[3]
	thread := strings.Split(inputUrl, "/")[5]

	actualPath, _ := os.Getwd()
	os.MkdirAll(board, os.ModePerm)
	os.MkdirAll(board+"//"+thread, os.ModePerm)
	pathResult := actualPath + "//" + board + "//" + thread

	for _, each := range findImages(string(body)) {
		linkImg = "http:" + strings.Replace(each, "s.jpg", ".jpg", 1)
		nameImg = re.FindAllString(linkImg, -1)[1] + ".jpg"
		wg.Add(1)
		go downloadFile(&wg, linkImg, nameImg, pathResult)
	}

	wg.Wait()
	fmt.Println("\n" + "[*] DOWNLOAD COMPLETE [*]")

}
