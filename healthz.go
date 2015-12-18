package healthz

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func init() {
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})
}

func OK(addr string, timeout time.Duration) error {
	start := time.Now()
	for {
		left := timeout - time.Since(start)
		if left < 0 {
			break
		}

		client := &http.Client{Timeout: left}
		resp, e := client.Get("http://" + addr + "/healthz")
		if e != nil {
			continue
		}
		defer resp.Body.Close()

		body, e := ioutil.ReadAll(resp.Body)
		if e != nil {
			return e
		}

		if string(body) != "OK" {
			return fmt.Errorf("/healthz returned %v", string(body))
		}
		return nil
	}
	return fmt.Errorf("Timeout checking %s", addr)
}
