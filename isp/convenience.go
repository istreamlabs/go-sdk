package isp

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"

	link "github.com/tent/http-link-go"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

// HighLevelClient wraps an APIClient and http.Client, as well as providing
// custom high-level functionality like pagination.
type HighLevelClient struct {
	*APIClient
	*http.Client
}

// AllPages is a high-level utility to automatically paginate and return the
// set of collected response items for paginated API calls.
func (c *HighLevelClient) AllPages(req interface{}) (interface{}, *http.Response, GenericOpenAPIError) {
	// Every request object has its own `Execute` method since they all return
	// different structs. We use reflection to call because of that.
	out := reflect.ValueOf(req).MethodByName("Execute").Call([]reflect.Value{})
	items := out[0]
	resp := out[1].Interface().(*http.Response)
	err := out[2].Interface().(GenericOpenAPIError)

	if err.Error() != "" {
		return nil, resp, err
	}

processResponse:
	// Check for a `next` link relation header for pagination. If present, we
	// must follow it, append the data, and repeat until no more `next` link
	// is available (signaling the end of the list).
	for _, header := range resp.Header["Link"] {
		links, err := link.Parse(header)
		if err != nil {
			return nil, resp, GenericOpenAPIError{
				error: err.Error(),
			}
		}

		for _, parsed := range links {
			if parsed.Rel == "next" {
				// Link URIs may be relative, so resolve them.
				loc, _ := url.Parse(parsed.URI)
				loc = resp.Request.URL.ResolveReference(loc)

				// Response gets set to the next response in the series.
				resp, err = c.Get(parsed.URI)
				if err != nil {
					return nil, resp, GenericOpenAPIError{
						error: err.Error(),
					}
				}

				data, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					return nil, resp, GenericOpenAPIError{
						error: err.Error(),
					}
				}

				// Since we don't know the return type, use reflection to create a new
				// instance of it to store the current response's data, which we'll
				// decode and append to the existing items list.
				additional := reflect.MakeSlice(items.Elem().Type(), 0, 0)
				addIface := additional.Interface()
				err = c.decode(&addIface, data, resp.Header.Get("Content-Type"))
				if err != nil {
					return nil, resp, GenericOpenAPIError{
						body:  data,
						error: err.Error(),
					}
				}

				// Append the new items and then pop back up top to process the
				// headers in this latest response.
				items = reflect.AppendSlice(items, additional)
				goto processResponse
			}
		}
	}

	return items.Interface(), resp, err
}

// ListSourcesAll returns a list of all source summaries. This is a convenience
// function which adds type information to `AllPages`.
func (c *HighLevelClient) ListSourcesAll(req ApiListSourcesRequest) ([]Summary, *http.Response, GenericOpenAPIError) {
	items, resp, err := c.AllPages(req)
	return items.([]Summary), resp, err
}

// ListChannelsAll returns a list of all channel summaries. This is a
// convenience function which adds type information to `AllPages`.
func (c *HighLevelClient) ListChannelsAll(req ApiListChannelsRequest) ([]Summary2, *http.Response, GenericOpenAPIError) {
	items, resp, err := c.AllPages(req)
	return items.([]Summary2), resp, err
}

// NewWithClientCredentials creates a new authenticated client using the OAuth
// 2.0 client credentials flow, caching and refreshing the access token as
// needed whenever requests to the API are made.
func NewWithClientCredentials(clientID, clientSecret, organization string) *HighLevelClient {
	auth := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     "https://istreamplanet.auth0.com/oauth/token",
		EndpointParams: url.Values{
			"audience": []string{"https://platform.dtc.istreamplanet.net/" + organization},
		},
	}

	config := NewConfiguration()
	config.HTTPClient = auth.Client(oauth2.NoContext)

	client := NewAPIClient(config)
	return &HighLevelClient{
		APIClient: client,
		Client:    config.HTTPClient,
	}
}
