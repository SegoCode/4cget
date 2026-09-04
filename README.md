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
[![License: MIT](https://img.shields.io/badge/License-MIT-blue?style=flat-square)](https://github.com/SegoCode/4cget/blob/main/LICENSE)
[![GitHub downloads](https://img.shields.io/github/downloads/SegoCode/4cget/total?style=flat-square&label=downloads)](https://github.com/SegoCode/4cget/releases)
[![Bitcoin BTC](https://img.shields.io/badge/buy_me_a_coffee-BTC-F7931A?style=flat-square&logo=bitcoin&logoColor=white)](https://github.com/SegoCode/SegoCode/discussions/2)

Easy to use, simply and fast 4chan thread media downloader. Simple, easy and functional.

> [!NOTE]
> I archived this repository in July 2025, when Cloudflare was deployed on `boards.4chan.org` and `i.4cdn.org` returning a challenge message instead of the thread. I didn’t have time to investigate how Cloudflare challenge works... but now I’ve seen that it works reliably with dynamic IPs.... Manually browsing 4chan regularly before running the tool can improve your IP reputation enough to avoid triggering the challenge. After testing it several times, **so I unarchived the project and resuming development.**

## Features

- Portable, single executable
- Customizable monitor mode and intervals
- Configurable output root
- Standard library only

## Quick Start & Information

4cget organizes downloaded files by board and thread. By default, it creates this structure in the current directory:

```shell
current-directory/
└── board/
    └── thread/
        └── files
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

#### Select the Output Folder

Use `--o <folder>` to set the root of the download structure. 4cget accepts relative and absolute paths, creates missing directories, and keeps the `board/thread` hierarchy below the selected root.

```shell
4cget https://boards.4chan.org/w/thread/... --o ./downloads
4cget --o /media/4chan https://boards.4chan.org/w/thread/...
```

`--o` can appear before or after the thread URL and can be combined with `--sleep` or `--monitor`:

```shell
4cget --sleep 2 https://boards.4chan.org/w/thread/... --o ./downloads
4cget --o ./downloads --monitor 60 https://boards.4chan.org/w/thread/...
```

The first command creates:

```shell
downloads/
└── w/
    └── thread-id/
        └── files
```

#### Display Help Message

Use the `--help` flag to display the help message with all available options:

```shell
4cget --help
```

#### Display Version

Use `--version` to print the installed version without making network requests:

```shell
4cget --version
```

Release binaries print their release tag, for example `4cget 1.7`. Local builds use `4cget 0.0` unless the build injects a version through Go linker flags.

> [!NOTE]
> All flags use the `--` prefix. For example, use `--o` instead of `-o`. `--monitor` and `--sleep` cannot be combined.


## Download

Latest binaries:

- Windows x64 https://github.com/SegoCode/4cget/releases/latest/download/4cget-windows-amd64.exe
- macOS Apple Silicon https://github.com/SegoCode/4cget/releases/latest/download/4cget-darwin-arm64
- Linux x64 https://github.com/SegoCode/4cget/releases/latest/download/4cget-linux-amd64

All platforms: https://github.com/SegoCode/4cget/releases/latest

---
<p align="center"><a href="https://github.com/SegoCode/4cget/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=SegoCode/4cget" />
</a></p>
