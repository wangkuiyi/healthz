package healthz

import (
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestOK(t *testing.T) {
	assert := assert.New(t)

	assert.NotNil(OK(":12345", time.Second))

	if addr, e := stealUnusedPort(); e != nil {
		t.Skip("Cannot steal an unused local port: ", e)
	} else {
		go func() {
			time.Sleep(2 * time.Second) // Artificial delay
			if e := http.ListenAndServe(addr, nil); e != nil {
				// NOTE: there might be a chance that during the
				// period after the invocation of stealUnusedPort and
				// that of http.ListenAndServe, some other process on
				// the system acquires the stolen port.
				t.Skip("Cannot listen and serve on ", addr, " : ", e)
			}
		}()

		assert.NotNil(OK(addr, 1*time.Second)) // Less than delay
		assert.Nil(OK(addr, 3*time.Second))    // Longer than delay
	}
}

func stealUnusedPort() (string, error) {
	l, e := net.Listen("tcp", ":0") // OS allocates a free port.
	if e != nil {
		return "", e
	}
	addr := l.Addr().String()
	l.Close()
	return addr, nil
}
