# 4cget
<img  src="https://raw.githubusercontent.com/SegoCode/4cget/main/media/demo1.3.gif">

<p align="center">
  <a href="#about">About</a> •
  <a href="#features">Features</a> •
  <a href="#quick-start--information">Quick Start & Information</a> •
  <a href="#download">Download</a> 
</p>

## About
Easy to use, simply and fast 4chan thread media downloader. Simple, easy and functional.

## Features

- Portable, single executable
- Not affected by cloudflare
- No dependences, no go mod

## Quick Start & Information

> [!CAUTION]
> Since 4cget is multithreaded, some CDN may detect it as a ddos attack and subsequent executions may not work as expected. [#4](https://github.com/SegoCode/4cget/issues/4)

<details>
  <summary>Thread lifecycle and download process in concurrent image downloading. Click here to show it.</summary> 
  <p align="center"><img src="https://raw.githubusercontent.com/SegoCode/4cget/main/media/diagram.png"></p>
</details>

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

One parameter, the url of the thread you want to download;
```shell
4cget https://boards.4channel.org/w/thread/...
```
Or `monitor mode` and check for new files every specified seconds;
```shell
4cget https://boards.4channel.org/w/thread/... -monitor 10
```
*In this example 4cget will check every 10 seconds.*


## Download

https://github.com/SegoCode/4cget/releases/

---
<p align="center"><a href="https://github.com/SegoCode/4cget/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=SegoCode/4cget" />
</a></p>
