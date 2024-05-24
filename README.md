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
<details>
  <summary>Thread lifecycle and download process in a concurrent image downloader application. Click here to show it.</summary> 
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
cd code\4cget
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
