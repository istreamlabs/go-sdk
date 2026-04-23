package isp

import (
	"net/http"

	"golang.org/x/oauth2"
)

type CloseableTransport struct {
	http.RoundTripper
}

type idleConnectionsCloser interface {
	CloseIdleConnections()
}

func (t *CloseableTransport) CloseIdleConnections() {
	base := t.RoundTripper
	if o, ok := base.(*oauth2.Transport); ok {
		// oath2.Transport doesn't implement what we need, but we know it stores
		// the underlying RoundTripper in a field called Base and can access it
		// here. It's likely an `http.Transport` and should implement the
		// connection closer.
		base = o.Base
	}

	if base == nil {
		base = http.DefaultTransport
	}

	if closer, ok := base.(idleConnectionsCloser); ok {
		closer.CloseIdleConnections()
	} else {
		panic("HTTP transport does not implement IdleConnectionsCloser")
	}
}

func NewCloseableTransport(t http.RoundTripper) *CloseableTransport {
	// Attempt to fail fast if we get an invalid transport. Still have to check
	// later as e.g. `oauth2.Transport.Base` could change at runtime.
	switch t.(type) {
	case *oauth2.Transport, idleConnectionsCloser:
		// ok
	default:
		panic("HTTP transport does not implement IdleConnectionsCloser")
	}
	return &CloseableTransport{t}
}
