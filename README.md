# 4cget

<img  src="https://raw.githubusercontent.com/SegoCode/4cget/main/media/demo1.3.gif">

<p align="center">
  <a href="#about">About</a> •
  <a href="#features">Features</a> •
  <a href="#quick-start--information">Quick Start & Information</a> •
  <a href="#download">Download</a> 
</p>

## About
[![Top language](https://img.shields.io/github/languages/top/SegoCode/4cget?style=flat-square)](https://github.com/SegoCode/4cget)
[![Repository size](https://img.shields.io/github/repo-size/SegoCode/4cget?style=flat-square&label=repo%20size)](https://github.com/SegoCode/4cget)
[![Commit activity per year](https://img.shields.io/github/commit-activity/y/SegoCode/4cget?style=flat-square&label=commits)](https://github.com/SegoCode/4cget/graphs/commit-activity)
[![GitHub downloads](https://img.shields.io/github/downloads/SegoCode/4cget/total?style=flat-square&label=downloads)](https://github.com/SegoCode/4cget/releases)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue?style=flat-square)](https://github.com/SegoCode/4cget/blob/main/LICENSE)
[![Bitcoin BTC](https://img.shields.io/badge/buy_me_a_coffee-BTC-F7931A?style=flat-square&logo=bitcoin&logoColor=white)](https://github.com/SegoCode/SegoCode/discussions/2)

Easy to use, simply and fast 4chan thread media downloader. Simple, easy and functional.

> [!NOTE]
> We archived this repository in July 2025 when Cloudflare tightened bot detection on `boards.4chan.org`, where edge rules evaluate IP reputation, ASN, TLS fingerprinting, and User-Agent per request, returning an interstitial challenge or 403 instead of thread HTML if the threshold is exceeded, while `i.4cdn.org` serves media under much looser limits. Because scoring is dynamic, regular manual browsing on 4chan before running the tool can improve your IP reputation score enough to pass these checks. After verifying that both endpoints complete fetches normally and the underlying HTML interface remains unchanged, **we unarchived the project and resume the development** .

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
