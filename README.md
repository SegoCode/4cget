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
- Configurable proxy
- Customizable monitor mode and intervals
- No dependences, no go mod

## Quick Start & Information

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

#### Use a Proxy Server

If you need to route your requests through a proxy server:

```shell
4cget https://boards.4channel.org/w/thread/... --proxy http://proxyserver:port
```

#### Proxy Authentication

If your proxy server requires authentication:

```shell
4cget https://boards.4channel.org/w/thread/... --proxy http://proxyserver:port --proxyuser username --proxypass password
```

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
