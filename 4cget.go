package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

var monitorMode bool

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
	defer wg.Done()

	resp, _ := http.Get(url)
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		urldext := strings.Split(url, ".jpg")[0]
		extList := []string{".png", ".webm", ".gif"}
		for _, ext := range extList {
			resp, _ = http.Get(urldext + ext)
			if resp.StatusCode != 404 {
				fileName = strings.Replace(fileName, ".jpg", ext, 1)
				break
			}
		}
	}

	if resp.StatusCode != 404 {
		if _, err := os.Stat(path + "//" + fileName); os.IsNotExist(err) || !monitorMode {

			img, _ := os.Create(path + "//" + fileName)
			defer img.Close()

			b, _ := io.Copy(img, resp.Body)

			var suffixes [5]string
			suffixes[0] = "B"
			suffixes[1] = "KB"
			suffixes[2] = "MB"
			suffixes[3] = "GB"
			suffixes[4] = "TB"

			base := math.Log(float64(b)) / math.Log(1024)
			getSize := math.Pow(1024, base-math.Floor(base))
			getSuffix := suffixes[int(math.Floor(base))]

			fmt.Printf("File downloaded: "+fileName+" - Size: %.2f "+string(getSuffix)+"\n", getSize)
		}
	}

}

func main() {
	var re = regexp.MustCompile("[0-9]+")
	var wg sync.WaitGroup
	var inputUrl string
	var linkImg string
	var nameImg string
	var secondsIteration int

	//Usage validation
	if len(os.Args) <= 1 {
		fmt.Println("[!] USAGE: 4cget https://boards.4channel.org/w/thread/.../...")
		os.Exit(1)
	}

	fmt.Println(len(os.Args))

	if len(os.Args) == 4 && strings.Compare(os.Args[2], "-monitor") == 0 {
		num, err := strconv.Atoi(os.Args[3])
		if err == nil {
			secondsIteration = num
			monitorMode = true
		}
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

	if monitorMode {
		fmt.Println("[*] MONITOR MODE ENABLE [*]" + "\n")
	}

	start := time.Now()
	files := 0

	board := strings.Split(inputUrl, "/")[3]
	thread := strings.Split(inputUrl, "/")[5]

	actualPath, _ := os.Getwd()
	os.MkdirAll(board, os.ModePerm)
	os.MkdirAll(board+"//"+thread, os.ModePerm)
	pathResult := actualPath + "//" + board + "//" + thread

	fmt.Println("Folder created : " + actualPath + "...")
	for {
		resp, _ := http.Get(inputUrl)
		body, _ := ioutil.ReadAll(resp.Body)

		for _, each := range findImages(string(body)) {
			if !strings.Contains(each, "s.4cdn.org") { //This server contains 4chan cosmetic resources
				linkImg = "http:" + strings.Replace(each, "s.jpg", ".jpg", 1)
				nameImg = re.FindAllString(linkImg, -1)[1] + ".jpg"
				wg.Add(1)
				go downloadFile(&wg, linkImg, nameImg, pathResult)
				files++
			}
		}

		wg.Wait()
		if !monitorMode {
			break
		} else {
			for i := secondsIteration; i >= 0; i-- {
				fmt.Printf("Press Ctrl+C to close 4cget" + "\n")
				fmt.Printf("Checking for new files in %v seconds...."+"\n", i)
				time.Sleep(1 * time.Second)
				print("\033[F")
				print("\033[F")
			}
		}

	}
	fmt.Printf("\n"+"✓ DOWNLOAD COMPLETE, %v FILES IN %v "+"\n", files, time.Since(start))
}
