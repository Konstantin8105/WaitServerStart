# WaitServerStart


[![Coverage Status](https://coveralls.io/repos/github/Konstantin8105/WaitServerStart/badge.svg?branch=master)](https://coveralls.io/github/Konstantin8105/WaitServerStart?branch=master)
[![Build Status](https://travis-ci.org/Konstantin8105/WaitServerStart.svg?branch=master)](https://travis-ci.org/Konstantin8105/WaitServerStart)
[![Go Report Card](https://goreportcard.com/badge/github.com/Konstantin8105/WaitServerStart)](https://goreportcard.com/report/github.com/Konstantin8105/WaitServerStart)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/Konstantin8105/WaitServerStart/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/Konstantin8105/WaitServerStart?status.svg)](https://godoc.org/github.com/Konstantin8105/WaitServerStart)

Simple waiter of server start

```golang
func main() {
	// Start a server
	srv, address := startHttpServer()

	// Wait starting of server
	<-wait.ByAddress(address)

	// Testing
	resp, err := http.Get(address)
	if err != nil {
		fmt.Println("RESPONSE")
		panic(err)
	}
	fmt.Println("Server is run...")
	_ = resp.Body.Close()
	if err := srv.Shutdown(context.Background()); err != nil {
		panic(err)
	}

	// Output: Server is run...
}
```
