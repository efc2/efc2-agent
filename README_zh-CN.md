# Efc Agent

[![Build Status](https://travis-ci.org/efc2/efc2-agent.svg?branch=master)](https://travis-ci.org/efc2/efc2-agent)
[![Go Report Card](https://goreportcard.com/badge/github.com/efc2/efc2-agent)](https://goreportcard.com/report/github.com/efc2/efc2-agent)

Efc 探针可以收集它所在操作系统的各种指标，然后发送到 [Efc](http://cloud.efc.one) 后端服务，探针由 Go 语言实现。

## 源代码编译

为了从源代码编译 Efc 探针，你需要准备一个 Go 语言环境，版本需要 [>= 1.7](https://golang.org/doc/install)。

```
$ mkdir -p $GOPATH/src/github.com/efc2
$ cd $GOPATH/src/github.com/efc2
$ git clone https://github.com/efc2/efc2-agent
$ cd efc2-agent
$ make build
```

## 使用方法

首先需要设置 license key，你可以在这里找到你的 license key，[https://cloud.efc.one/#/settings](https://cloud.efc.one/#/settings).

```
$ cp efc-agent.conf.example efc-agent.conf
$ vi efc-agent.conf
...
license_key = "*********************"
```

在前台运行探针：

```
$ ./bin/efc-agent
```

更多用法, 见:

```
$ ./bin/efc-agent --help
```

## 相关的资源

Efc 探针深受以下项目的影响：

- [ddagent](https://github.com/datadog/dd-agent)
- [telegraf](https://github.com/influxdata/telegraf)
- [prometheus](https://github.com/prometheus/prometheus)
- [mackerel](https://github.com/mackerelio/mackerel-agent)
