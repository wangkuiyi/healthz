package healthz

import (
	"net"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOK(t *testing.T) {
	assert := assert.New(t)

	assert.NotNil(OK(":12345", 1))

	l, e := net.Listen("tcp", ":0") // OS allocates a free port.
	if e != nil {
		t.Skip("Mocking healthz server cannot listen: ", e)
	}
	go http.Serve(l, nil)

	assert.Nil(OK(l.Addr().String(), 1))
}
