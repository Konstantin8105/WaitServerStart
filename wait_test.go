package wait_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	wait "github.com/Konstantin8105/WaitServerStart"
)

func TestByAddress(t *testing.T) {
	srv, address := startHttpServer()
	ch := wait.ByAddress(address)

	<-ch

	if err := srv.Shutdown(context.Background()); err != nil {
		t.Error(err)
	}
}

func startHttpServer() (*http.Server, string) {
	port := 8090
	address := fmt.Sprintf("http://127.0.0.1:%d", port)

	srv := &http.Server{Addr: fmt.Sprintf(":%d", port)}

	go func() {
		time.Sleep(5 * time.Millisecond)
		_ = srv.ListenAndServe()
	}()

	return srv, address
}

func ExampleByAddress() {
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
