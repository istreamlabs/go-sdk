package main

import (
	"encoding/json"
	"testing"

	"github.com/istreamlabs/go-sdk/isp"
	"github.com/stretchr/testify/assert"
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
