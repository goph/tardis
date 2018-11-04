# Tardis

[![Build Status](https://img.shields.io/travis/com/goph/tardis.svg?style=flat-square)](https://travis-ci.com/goph/tardis)
[![Go Report Card](https://goreportcard.com/badge/github.com/goph/tardis?style=flat-square)](https://goreportcard.com/report/github.com/goph/tardis)
[![GolangCI](https://golangci.com/badges/github.com/goph/tardis.svg)](https://golangci.com/r/github.com/goph/tardis)
[![GoDoc](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/goph/tardis)

**Time machine for your application to control time.**

## Best practices

One should not create a hard dependency on this library in code. Instead custom, private interfaces
should be created which is compatible with this library or can wrap it with a simple adapter.

For example:

```go
package main

import(
	"time"
	
	"github.com/goph/tardis"
)

type myClock interface {
	Now() time.Time
}

type myService struct {
	clock myClock
}

func newMyService(clock myClock) *myService {
	return &myService{
		clock: clock,
	}
}

func main() {
	service := newMyService(tardis.SystemClock)
}
```


## Development

Install the dependencies:

```bash
$ make vendor # or dep ensure
```

When all coding and testing is done, please run the test suite:

```bash
$ make check
```


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
