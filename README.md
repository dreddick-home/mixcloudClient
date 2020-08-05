# mixcloudclient

A simple CLI for interacting with mixcloud

## Usage


### search

```
  mixcloudclient search [flags]

Flags:
  -e, --excludes strings   Must exclude term, multiple items accepted.
  -h, --help               help for search
  -i, --includes strings   Must include term, multiple items accepted.
  -m, --max int32          Max results (in multiples of 100). Default 5. (default 5)
  -t, --term string        Search Term
  -w, --workers int32      The max number of concurrent workers. Defaults to number of cores of system. (default 8)
```

#### Example

Find mixes using term 'nick warren' and only return those which include 'cream'

```
mixcloudclient search -t 'nick warren' -i cream
```


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