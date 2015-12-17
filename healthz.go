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
	client := &http.Client{Timeout: timeout}
	resp, e := client.Get("http://" + addr + "/healthz")
	if e != nil {
		return e
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
