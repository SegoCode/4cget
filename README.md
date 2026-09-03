# 4cget

<img  src="https://raw.githubusercontent.com/SegoCode/4cget/main/media/demo1.3.gif">

<p align="center">
  <a href="#about">About</a> •
  <a href="#features">Features</a> •
  <a href="#quick-start--information">Quick Start & Information</a> •
  <a href="#download">Download</a> 
</p>

## About
[![Top language](https://img.shields.io/github/languages/top/{username}/{reponame}?style=flat-square)](https://github.com/{username}/{reponame})
[![Repository size](https://img.shields.io/github/repo-size/{username}/{reponame}?style=flat-square&label=repo%20size)](https://github.com/{username}/{reponame})
[![Commit activity per year](https://img.shields.io/github/commit-activity/y/{username}/{reponame}?style=flat-square&label=commits)](https://github.com/{username}/{reponame}/graphs/commit-activity)
[![GitHub downloads](https://img.shields.io/github/downloads/{username}/{reponame}/total?style=flat-square&label=downloads)](https://github.com/{username}/{reponame}/releases)
[![Licencia: PolyForm Noncommercial + GNU AGPL-3.0](https://img.shields.io/badge/License-PolyForm%20Noncommercial%20%2B%20GNU%20AGPL--3.0-blue?style=flat-square)](https://github.com/{username}/{reponame}/blob/main/LICENSE)
[![Bitcoin BTC](https://img.shields.io/badge/buy_me_a_coffee-BTC-F7931A?style=flat-square&logo=bitcoin&logoColor=white)](https://github.com/SegoCode/SegoCode/discussions/2)

Easy to use, simply and fast 4chan thread media downloader. Simple, easy and functional.

> [!NOTE]
> We archived this repository on 2025-07-05 after Cloudflare raised bot scoring on `boards.4chan.org`. Cloudflare terminates the public HTML origin and scores each GET from IP reputation, ASN, TLS fingerprint, and User-Agent. A score above the challenge threshold returns interstitial HTML or 403 instead of the thread document. A score below it returns the post list. Clients fetch media from `i.4cdn.org`, a separate CDN with a lower bot threshold.
>
> The score is computed per request. The same IP can fail and pass on consecutive fetches. We retested live threads against the board origin and the image CDN. Those fetches completed. We unarchived because the HTML interface is intact; access tracks Cloudflare's edge score.

## Features

- Portable, single executable
- Customizable monitor mode and intervals
- No dependences, no go mod

## Quick Start & Information

4cget downloads the files organized by boards and threads.

```shell
root
  └───board
      └───thread
            └───files
```

run from source code (Golang installation required).

```shell
git clone https://github.com/SegoCode/4cget
cd 4cget\code
go run 4cget.go https://boards.4channel.org/w/thread/...
```
Or better [donwload a binary](https://github.com/SegoCode/4cget/releases).

### Available Parameters

`4cget` provides various parameters to customize its behavior. Below are detailed examples and explanations for each available option:

#### Basic Usage

Download all images from a thread:

```shell
4cget https://boards.4channel.org/w/thread/...
```

#### Enable Monitor Mode

Use the `--monitor` flag to enable monitor mode, which checks for new files every specified number of seconds:

```shell
4cget https://boards.4channel.org/w/thread/... --monitor 10
```

*In this example, `4cget` will check every 10 seconds for new images.*

####  Add Delay Between Downloads

Use the `--sleep` flag to add a delay between downloads (useful to avoid rate-limiting):

```shell
4cget https://boards.4channel.org/w/thread/... --sleep 2
```

*This adds a 2-second delay between each download.*

#### Display Help Message

Use the `--help` flag to display the help message with all available options:

```shell
4cget --help
```

> [!NOTE]
> All flags must be prefixed with `--`. For example, use `--monitor` instead of `-monitor`.


## Download

https://github.com/SegoCode/4cget/releases/

---
<p align="center"><a href="https://github.com/SegoCode/4cget/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=SegoCode/4cget" />
</a></p>
