# mixcloudclient

A simple CLI for interacting with Mixcloud

<p align="left">
<img src="https://img.shields.io/github/go-mod/go-version/dreddick-home/mixcloudclient">
<img src="https://img.shields.io/github/v/release/dreddick-home/mixcloudclient">
<img src="https://github.com/dreddick-home/mixcloudclient/workflows/CICD/badge.svg">
<img src="https://img.shields.io/badge/License-Apache%202.0-blue.svg">
</p>


## Description

This is a utility for interacting with [Mixcloud](https://www.mixcloud.com).

The project was started to find an easier way to search Mixcloud for content.

### Search

The search function uses go routines to query the Mixcloud API quickly using a search term. Client-side filters can be applied to exclude or include items from the results.

## Usage

Use the following syntax to run mixcloudclient commands from your terminal window:

```console
$ mixcloudclient [command] [flags]
```

### Commands

#### Search

```
Flags:
  -e, --excludes strings   Must exclude term, multiple items accepted.
  -h, --help               help for search
  -i, --includes strings   Must include term, multiple items accepted.
  -m, --max int32          Max results (in 100s). Default 20.
  -t, --term string        Search Term
  -w, --workers int32      The max number of concurrent workers. Defaults to number of cores of system. (default 8)
```


![Nick Warren Search](https://raw.githubusercontent.com/dreddick-home/mixcloudclient/master/img/mixcloudclient_usage1.gif)

## Docker

```console
$ docker run dreddick/mixcloudclient:v0.0.2 search -t 'nick warren' -i 1995
```


## Install

### Build and Install the Binaries from Source

#### Prerequisite Tools

* Git
* Go 


#### Fetch from GitHub

```console
$ git clone https://github.com/dreddick-home/mixcloudclient.git
$ cd mixcloudclient
$ go build -o /usr/local/bin/mixcloudclient
```


### Releases

See https://github.com/dreddick-home/mixcloudclient/releases
