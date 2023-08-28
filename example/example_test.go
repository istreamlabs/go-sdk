package main

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/istreamlabs/go-sdk/isp"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
)

// TestSerializeEmptyMap tests that an empty map is serialized as an empty
// object in JSON and sent over the wire. This is necessary for some API
// calls as the empty object has semantic meaning to the backend, so it is
// not safe to omit it.
// See also `templates/model_simple.mustache` which was modified from the
// original OpenAPI Generator Go template to facilitate this new behavior.
func TestSerializeEmptyMap(t *testing.T) {
	ch := isp.Channel{
		Publishing: &isp.ChannelPublishing{
			Publications: []isp.ChannelPublishingPublicationsInner{
				{
					Startover: &isp.ChannelPublishingPublicationsInnerStartover{
						FirstProgramStart: &map[string]interface{}{},
					},
				},
			},
		},
	}

	b, err := json.Marshal(ch)
	assert.NoError(t, err)

	assert.Contains(t, string(b), `"first_program_start":{}`)
}

type CustomTransport struct {
	Base http.RoundTripper
}

func (t *CustomTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return nil, nil
}

type CustomTransportWithClose struct {
	Base http.RoundTripper
}

func (t *CustomTransportWithClose) RoundTrip(req *http.Request) (*http.Response, error) {
	return nil, nil
}

func (t *CustomTransportWithClose) CloseIdleConnections() {
	// Do nothing
}

func TestCloseableTransport(t *testing.T) {
	custom1 := &CustomTransport{http.DefaultTransport}
	assert.Panics(t, func() {
		// Creation should fail because the custom transport does not implement
		// the CloseIdleConnections method.
		isp.NewCloseableTransport(custom1)
	})

	custom2 := &oauth2.Transport{Base: http.DefaultTransport}
	assert.Panics(t, func() {
		closeable2 := isp.NewCloseableTransport(custom2)

		// Screw things up after creation by setting a custom transport without
		// support for CloseIdleConnections as the oauth2.Transport's base.
		custom2.Base = custom1
		closeable2.CloseIdleConnections()
	})

	custom3 := &CustomTransportWithClose{http.DefaultTransport}
	closeable3 := isp.NewCloseableTransport(custom3)
	closeable3.CloseIdleConnections()
}
