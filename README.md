# mixcloudclient

A simple CLI for interacting with mixcloud

## Usage

Use the following syntax to run mixcloudclient commands from your terminal window:

```console
$ mixcloudclient [command] [flags]
```

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


![Nick Warren Search](https://raw.githubusercontent.com/dreddick-home/mixcloudclient/master/img/mixcloudclient_usage1.gif)


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