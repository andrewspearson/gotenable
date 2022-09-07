# GoTenable
[![PkgGoDev](https://pkg.go.dev/badge/github.com/andrewspearson/gotenable)](https://pkg.go.dev/github.com/andrewspearson/gotenable)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/andrewspearson/gotenable)
[![License: MIT](https://img.shields.io/badge/License-MIT-brightgreen.svg)](LICENSE)
[![GoReportCard](https://goreportcard.com/badge/github.com/andrewspearson/gotenable)](https://goreportcard.com/report/github.com/andrewspearson/gotenable)

GoTenable is a Go library for interfacing with Tenable product APIs.
## Getting Started
Tenable.io
```go
package main

import (
	"github.com/andrewspearson/gotenable/gotenableutils"
	"github.com/andrewspearson/gotenable/tenableio"
)

func main() {
	config := tenableio.Config{}
	config.AccessKey = `ACCESS_KEY`
	config.SecretKey = `SECRET_KEY`

	tio := tenableio.New(config)

	resp := tio.Get(`/scans`)

	gotenableutils.PPrintJSON(resp.Body)
}
```
Tenable.sc
```go
package main

import (
	"github.com/andrewspearson/gotenable/gotenableutils"
	"github.com/andrewspearson/gotenable/tenablesc"
)

func main() {
	config := tenablesc.Config{}
	config.BaseURL = `https://tenablesc.company.local`
	config.AccessKey = `ACCESS_KEY`
	config.SecretKey = `SECRET_KEY`

	tsc := tenablesc.New(config)

	resp := tsc.Get(`/scan`)

	gotenableutils.PPrintJSON(resp.Body)
}
```
Downloads
```go
package main

import (
	"github.com/andrewspearson/gotenable/downloads"
	"github.com/andrewspearson/gotenable/gotenableutils"
)

func main() {
	config := downloads.Config{}
	config.BearerToken = `BEARER_TOKEN`

	dl := downloads.New(config)

	resp := dl.Get(`/pages`)

	gotenableutils.PPrintJSON(resp.Body)
}
```
## Documentation
See [DOCUMENTATION.md](DOCUMENTATION.md) for complete details.
