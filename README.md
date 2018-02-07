# Efc Agent

[![Build Status](https://travis-ci.org/efc2/efc2-agent.svg?branch=master)](https://travis-ci.org/efc2/efc2-agent)
[![Go Report Card](https://goreportcard.com/badge/github.com/efc2/efc2-agent)](https://goreportcard.com/report/github.com/efc2/efc2-agent)
[![codecov](https://codecov.io/gh/efc2/efc2-agent/branch/master/graph/badge.svg)](https://codecov.io/gh/efc2/efc2-agent)

[中文版 README](README_zh-CN.md)

Efc Agent is written in Go for collecting metrics from the system it's
running on, or from other services, and sending them to [Efc](https://cloud.efc.one).

## Building from source

To build Efc Agent from the source code yourself you need to have a working Go environment with [version 1.7+](https://golang.org/doc/install).

```
$ mkdir -p $GOPATH/src/github.com/efc2
$ cd $GOPATH/src/github.com/efc2
$ git clone https://github.com/efc2/efc2-agent
$ cd efc2-agent
$ make build
```

## Usage

First you need to set a license key, which can be found at [https://cloud.efc.one/#/settings](https://cloud.efc.one/#/settings).

```
$ cp efc-agent.conf.example efc-agent.conf
$ vi efc-agent.conf
...
license_key = "*********************"
```

Run the agent in foreground:

```
$ ./bin/efc-agent
```

For more options, see:

```
$ ./bin/efc-agent --help
```

## Related works

I have been influenced by the following great works:

- [ddagent](https://github.com/datadog/dd-agent)
- [telegraf](https://github.com/influxdata/telegraf)
- [prometheus](https://github.com/prometheus/prometheus)
- [mackerel](https://github.com/mackerelio/mackerel-agent)
