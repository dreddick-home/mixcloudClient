# mixcloudclient

A simple CLI for interacting with Mixcloud

## Description

This is a utility for interacting with [Mixcloud](https://www.mixcloud.com).

The project was started to find an easier way to search Mixcloud for content.

### Search

The search function uses go routines to query the Mixcloud API quickly using a search term. Client-side filters can be applied to exclude or include items from the results.

## Usage


### search

```
Flags:
  -e, --excludes strings   Must exclude term, multiple items accepted.
  -h, --help               help for search
  -i, --includes strings   Must include term, multiple items accepted.
  -m, --max int32          Max results (in multiples of 100). Default 5. (default 5)
  -t, --term string        Search Term
  -w, --workers int32      The max number of concurrent workers. Defaults to number of cores of system. (default 8)
```


![Nick Warren Search](https://raw.githubusercontent.com/dreddick-home/mixcloudclient/master/img/mixcloudclient_usage1.gif)


## Install

### Using go

```console
$ go get -u github.com/dreddick-home/mixcloudclient
```

Install in a custom location

```console
$ git clone https://github.com/dreddick-home/mixcloudclient.git
$ cd mixcloudclient
$ go build -o /usr/local/bin/mixcloudclient
```


### Releases

See https://github.com/dreddick-home/mixcloudclient/releases
