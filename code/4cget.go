package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

const version = "1.7" // Current version

var monitorMode bool

// SiteInfo holds the URL pattern, regex for image extraction, and an ID.
type SiteInfo struct {
	ID    string
	URL   string
	ImgRE *regexp.Regexp
}

// Initialize the site info map with URL patterns and corresponding regex.
var siteInfoMap = map[string]SiteInfo{
	"4chan": {
		ID:    "4chan",
		URL:   "https://boards.4chan.org",
		ImgRE: regexp.MustCompile(`<a[^>]+href="(//i\.4cdn\.org[^"]+)"`),
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
			url = strings.Replace(url, "//i.4cdn.org", "https://i.4cdn.org", 1)
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

func downloadFile(wg *sync.WaitGroup, url string, fileName string, path string, client *http.Client) {
	defer wg.Done()

	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("[!] Error downloading file:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 429 {
		fmt.Println("[!] Received HTTP 429 Too Many Requests. You are being rate-limited.")
		fmt.Println("[!] Consider using the --sleep flag to add delays between downloads.")
		return
	}

	if resp.StatusCode != 404 && resp.StatusCode == 200 {
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
	} else {
		fmt.Printf("[!] Received HTTP %d for %s\n", resp.StatusCode, url)
	}
}

// checkForUpdates checks the latest release from GitHub and compares it with the current version.
func checkForUpdates() (latestVersion string, updateAvailable bool) {
	apiURL := "https://api.github.com/repos/SegoCode/4cget/releases/latest"
	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("[!] Error checking for updates:", err)
		return "", false
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Printf("[!] GitHub API returned status code %d\n", resp.StatusCode)
		return "", false
	}

	var release struct {
		TagName string `json:"tag_name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		fmt.Println("[!] Error decoding GitHub API response:", err)
		return "", false
	}

	latestVersion = strings.TrimPrefix(release.TagName, "v")
	if latestVersion != version {
		return latestVersion, true
	}
	return latestVersion, false
}

// displayHelp shows the help message with explanations and examples.
func displayHelp() {
	fmt.Println(`
4cget - A tool to download images from 4chan threads.

Usage:
  4cget [options] <thread_url>

Options:
  --help                 Display this help message.
  --monitor <seconds>    Enable monitor mode with interval in seconds.
                         The program will check for new images every specified interval.
  --sleep <seconds>      Sleep duration in seconds between downloads.
                         Useful to avoid getting rate-limited by the server.
  --proxy <proxy_url>    Proxy URL (e.g., http://proxyserver:port).
  --proxyuser <user>     Proxy username for authentication.
  --proxypass <pass>     Proxy password for authentication.

Examples:

  Basic usage:
    4cget https://boards.4chan.org/w/thread/123456

  Enable monitor mode with a 60-second interval:
    4cget --monitor 60 https://boards.4chan.org/w/thread/123456

  Use a proxy with authentication:
    4cget --proxy http://proxyserver:port --proxyuser username --proxypass password https://boards.4chan.org/w/thread/123456

  Add delay between downloads to prevent rate-limiting:
    4cget --sleep 2 https://boards.4chan.org/w/thread/123456

Note:
  - Ensure that all flags are prefixed with '--'.
  - The thread URL must be a valid URL from a supported site.
  - Use '--sleep' to add delays between downloads to avoid getting rate-limited (HTTP 429 errors).
`)
}

func main() {
	var wg sync.WaitGroup
	var inputUrl string
	var thread string
	var siteID string

	// Define command-line flags
	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	helpFlag := fs.Bool("help", false, "Display this help message")
	monitorIntervalFlag := fs.Int("monitor", 0, "Enable monitor mode with interval in seconds")
	sleepFlag := fs.Int("sleep", 0, "Sleep duration in seconds between downloads")
	proxyFlag := fs.String("proxy", "", "Proxy URL (e.g., http://proxyserver:port)")
	proxyUserFlag := fs.String("proxyuser", "", "Proxy username")
	proxyPassFlag := fs.String("proxypass", "", "Proxy password")

	// Manually parse flags and positional arguments
	var args []string
	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		if arg == "--" {
			// All remaining args are positional
			args = append(args, os.Args[i+1:]...)
			break
		}
		if strings.HasPrefix(arg, "--") {
			// Flag
			fs.Parse(os.Args[i:])
			break
		}
		if strings.HasPrefix(arg, "-") && !strings.HasPrefix(arg, "--") {
			fmt.Printf("Invalid flag: %s. Flags must start with '--'.\n", arg)
			os.Exit(1)
		}
		// Positional argument
		args = append(args, arg)
	}

	// After parsing flags, any remaining arguments are positional
	args = append(args, fs.Args()...)

	// If --help is provided, display help message and exit
	if *helpFlag {
		displayHelp()
		return
	}

	// Input URL validation
	if len(args) < 1 {
		fmt.Println("[!] USAGE: 4cget [options] <thread_url>")
		fmt.Println("Use '--help' to see available options.")
		os.Exit(1)
	}
	inputUrl = args[0]

	monitorMode = (*monitorIntervalFlag > 0)
	secondsIteration := *monitorIntervalFlag
	sleepDuration := *sleepFlag
	proxyURL := *proxyFlag

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

	// Check for updates before starting the download
	latestVersion, updateAvailable := checkForUpdates()
	if updateAvailable {
		fmt.Printf("[*] UPDATE AVAILABLE %s [*]\n\n", latestVersion)
	}

	fmt.Println("[*] DOWNLOAD STARTED (" + inputUrl + ") [*]\n")
	if monitorMode {
		fmt.Println("[*] MONITOR MODE ENABLED [*]\n")
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

	fmt.Println("Folder created : " + actualPath + "...\n")

	// Setup HTTP client with optional proxy and authentication
	client := &http.Client{}
	if proxyURL != "" {
		proxyParsed, err := url.Parse(proxyURL)
		if err != nil {
			fmt.Println("[!] Invalid proxy URL:", err)
			os.Exit(1)
		}
		if *proxyUserFlag != "" {
			proxyParsed.User = url.UserPassword(*proxyUserFlag, *proxyPassFlag)
		}
		client.Transport = &http.Transport{Proxy: http.ProxyURL(proxyParsed)}
	}

	for { // Main loop for monitorMode
		resp, err := client.Get(inputUrl)
		if err != nil {
			fmt.Println("[!] Error fetching URL:", err)
			os.Exit(1)
		}
		body, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		imageURLs := findImages(string(body), siteID)
		for _, each := range imageURLs {
			parts := strings.Split(each, "/")
			nameImg := parts[len(parts)-1]
			wg.Add(1)
			go downloadFile(&wg, each, nameImg, pathResult, client)
			files++

			// Sleep between starting downloads if sleepDuration > 0
			if sleepDuration > 0 {
				time.Sleep(time.Duration(sleepDuration) * time.Second)
			}
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
