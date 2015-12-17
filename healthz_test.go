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

	l, e := net.Listen("tcp", ":0") // OS allocates a free port.
	if e != nil {
		t.Skip("Mocking healthz server cannot listen: ", e)
	}
	go func() {
		time.Sleep(2 * time.Second) // Artificial delay
		http.Serve(l, nil)
	}()

	assert.NotNil(OK(l.Addr().String(), 1*time.Second)) // Less than delay
	assert.Nil(OK(l.Addr().String(), 3*time.Second))    // Longer than delay
}
