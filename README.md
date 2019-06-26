ts
===
timestamp convert & compare tool

[![GoDoc](https://godoc.org/github.com/liujianping/ts?status.svg)](https://godoc.org/github.com/liujianping/ts) [![Go Report Card](https://goreportcard.com/badge/github.com/liujianping/ts)](https://goreportcard.com/report/github.com/liujianping/ts) [![Build Status](https://travis-ci.org/liujianping/ts.svg?branch=master)](https://travis-ci.org/liujianping/ts) [![Version](https://img.shields.io/github/tag/liujianping/ts.svg)](https://github.com/liujianping/ts/releases) [![Coverage Status](https://coveralls.io/repos/github/liujianping/ts/badge.svg?branch=master)](https://coveralls.io/github/liujianping/ts?branch=master)

## Install

### Shell Install support Linux & MacOS

````bash
# binary will be $(go env GOPATH)/bin/ts
$: curl -sfL https://raw.githubusercontent.com/liujianping/ts/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# In alpine linux (as it does not come with curl by default)
$: wget -O - -q https://raw.githubusercontent.com/liujianping/ts/master/install.sh | sh -s 

```` 

### Brew Install only MacOS
````bash
$: brew tap liujianping/tap && brew install ts
````

### Source Install
````bash
$: git clone git@github.com:liujianping/ts.git
$: cd ts
$: go install -mod vendor
````

## Quick Start

````bash
$: ts -h
timestamp convert & compare tool

Usage:
  ts [flags]

Examples:

	(now timestamp)	$: ts
	(now add)		$: ts --add 1d
	(now sub)		$: ts --sub 1d
	(convert)		$: ts "2019/06/24 23:30:10"
	(pipe)			$: echo "2019/06/24 23:30:10" | ts
	(format)		$: ts -f "2019/06/25 23:30:10"
	(before)		$: ts -b "2019/06/25 23:30:10" ; echo $?
	(after)			$: ts -a "2019/06/25 23:30:10" ; echo $?
	(timezone)		$: ts -f "2019/06/25 23:30:10" -z "Asia/Shanghai"


Flags:
      --add duration      add duration
  -a, --after string      after compare
  -b, --before string     before compare
  -f, --format string     time format
  -h, --help              help for ts
      --sub duration      sub duration
  -z, --timezone string   time zone
````

support system format:
- ANSIC       = "Mon Jan _2 15:04:05 2006"
- UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
- RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
- RFC822      = "02 Jan 06 15:04 MST"
- RFC822Z     = "02 Jan 06 15:04 -0700" - RFC822 with numeric zone
- RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
- RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
- RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" - RFC1123 with numeric zone
- RFC3339     = "2006-01-02T15:04:05Z07:00"
- RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
- Kitchen     = "3:04PM"
- Stamp      = "Jan _2 15:04:05"
- StampMilli = "Jan _2 15:04:05.000"
- StampMicro = "Jan _2 15:04:05.000000"
- StampNano  = "Jan _2 15:04:05.000000000"