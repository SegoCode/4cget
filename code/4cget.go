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

// SiteInfo holds the URL pattern, regex for image extraction, and an ID.
type SiteInfo struct {
	ID     string
	URL    string
	ImgRE  *regexp.Regexp
}

// Initialize the site info map with URL patterns and corresponding regex.
var siteInfoMap = map[string]SiteInfo{
	"4chan": {
		ID:    "4chan",
		URL:   "https://boards.4chan.org",
		ImgRE: regexp.MustCompile(`<a[^>]+href="(//is2\.4chan\.org[^"]+)"`),
	},
	"twochen": {
		ID:    "twochen",
		URL:   "https://sturdychan.help/",
		ImgRE: regexp.MustCompile(`(https?://[^/]+/assets/images/src/[a-zA-Z0-9]+\.(?:png|jpg))`),
	},
}

// findImages extracts image URLs from the given HTML based on the site specified.
func findImages(html, siteID string) []string {
	var out []string
	siteInfo, exists := siteInfoMap[siteID]
	if !exists {
		fmt.Printf("No site information found for ID: %s\n", siteID)
		return out
	}

	matches := siteInfo.ImgRE.FindAllStringSubmatch(html, -1)
	for _, match := range matches {
		url := match[1]
		if siteID == siteInfoMap["4chan"].ID {
			url = strings.Replace(url, "//is2.4chan.org", "https://i.4cdn.org", 1)
		}
		out = append(out, url)
	}

	uniqueOut := unique(out) // Clear array of duplicates
	return uniqueOut
}

// unique removes duplicate strings from a slice.
func unique(input []string) []string {
	u := make(map[string]bool)
	var uniqueList []string
	for _, val := range input {
		if _, ok := u[val]; !ok {
			u[val] = true
			uniqueList = append(uniqueList, val)
		}
	}
	return uniqueList
}

func downloadFile(wg *sync.WaitGroup, url string, fileName string, path string) {
	defer wg.Done()

	resp, _ := http.Get(url)
	defer resp.Body.Close()

	if resp.StatusCode != 404 {
		filePath := path + "/" + fileName
		if _, err := os.Stat(filePath); os.IsNotExist(err) || !monitorMode {
			img, err := os.Create(filePath)
			if err != nil {
				fmt.Println("[!] Error creating file:", err)
				return
			}
			defer img.Close()

			b, err := io.Copy(img, resp.Body)
			if err != nil {
				fmt.Println("[!] Error copying response body:", err)
				return
			}

			suffixes := []string{"B", "KB", "MB", "GB", "TB"}

			base := math.Log(float64(b)) / math.Log(1024)
			getSize := math.Pow(1024, base-math.Floor(base))
			getSuffix := suffixes[int(math.Floor(base))]

			fmt.Printf("File downloaded: %s - Size: %.2f %s\n", fileName, getSize, getSuffix)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	var inputUrl string
	var secondsIteration int
	var monitorMode bool
	var thread string
	var siteID string

	// Usage validation
	if len(os.Args) <= 1 {
		fmt.Println("[!] USAGE: 4cget https://boards.4channel.org/w/thread/.../...")
		os.Exit(1)
	}

	if len(os.Args) == 4 && strings.Compare(os.Args[2], "-monitor") == 0 {
		num, err := strconv.Atoi(os.Args[3])
		if err == nil {
			secondsIteration = num
			monitorMode = true
		}
	}

	// Input URL validation
	inputUrl = os.Args[1]
	parsedURL, errParse := url.ParseRequestURI(inputUrl)
	if errParse != nil {
		fmt.Println("[!] URL NOT VALID (Example: https://boards.4channel.org/w/thread/.../...)")
		os.Exit(1)
	}

	for _, site := range siteInfoMap {
		parsedSiteURL, err := url.Parse(site.URL)
		if err != nil {
			fmt.Printf("Error parsing site URL %s: %v\n", site.URL, err)
			continue
		}
		if parsedURL.Host == parsedSiteURL.Host {
			siteID = site.ID
			break
		}
	}

	if siteID == "" {
		fmt.Println("[!] Unsupported site")
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

	fmt.Println("[*] DOWNLOAD STARTED (" + inputUrl + ") [*]\n")
	if monitorMode {
		fmt.Println("[*] MONITOR MODE ENABLE [*]\n")
	}

	start := time.Now()
	files := 0

	// Parse board and thread from URL
	parts := strings.Split(inputUrl, "/")
	board := parts[3]

	// Handle the thread part depending on the site
	if siteID == siteInfoMap["4chan"].ID {
		thread = parts[5]
	} else {
		thread = parts[4]
	}

	// Create necessary directories
	actualPath, _ := os.Getwd()
	os.MkdirAll(fmt.Sprintf("%s/%s", actualPath, board), os.ModePerm)
	os.MkdirAll(fmt.Sprintf("%s/%s/%s", actualPath, board, thread), os.ModePerm)
	pathResult := fmt.Sprintf("%s/%s/%s", actualPath, board, thread)

	fmt.Println("Folder created : " + actualPath + "...")

	for { // Main loop for monitorMode
		resp, _ := http.Get(inputUrl)
		body, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		for _, each := range findImages(string(body), siteID) {
			parts := strings.Split(each, "/")
			nameImg := parts[len(parts)-1]
			wg.Add(1)
			go downloadFile(&wg, each, nameImg, pathResult)
			files++
		}
		wg.Wait()
		if !monitorMode {
			break // Exit main loop
		} else {
			for i := secondsIteration; i >= 0; i-- {
				fmt.Printf("Press Ctrl+C to close 4cget\n")
				fmt.Printf("Checking for new files in %v seconds....\n", i)
				time.Sleep(1 * time.Second)
				print("\033[F\033[F")
			}
		}
	}

	fmt.Printf("\n✓ DOWNLOAD COMPLETE, %v FILES IN %v\n", files, time.Since(start))
}
