package wait_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	freePort "github.com/Konstantin8105/FreePort"
	wait "github.com/Konstantin8105/WaitServerStart"
)

func TestByAddress(t *testing.T) {
	port, err := freePort.Get()
	if err != nil {
		t.Error(err)
	}

	address := fmt.Sprintf("http://127.0.0.1:%d", port)

	ch := wait.ByAddress(address)

	srv := startHttpServer(port)

	<-ch

	if err := srv.Shutdown(context.Background()); err != nil {
		t.Error(err)
	}
}

func startHttpServer(port int) *http.Server {
	srv := &http.Server{Addr: fmt.Sprintf(":%d", port)}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, "hello world\n")
	})

	go func() {
		time.Sleep(time.Millisecond)
		_ = srv.ListenAndServe()
	}()

	return srv
}
