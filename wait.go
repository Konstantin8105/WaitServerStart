package wait

import (
	"net/http"
	"time"
)

// ByAddress - wait of starting server
func ByAddress(address string) <-chan struct{} {
	ch := make(chan struct{})

	go func() {
		for {
			resp, err := http.Get(address)
			if err != nil {
				time.Sleep(time.Millisecond)
				continue
			}
			_ = resp.Body.Close()
			break
		}
		ch <- struct{}{}
		close(ch)
	}()

	return ch
}
