package isp

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"reflect"

	link "github.com/tent/http-link-go"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

// RelNext is the `next` link relationship, used for pagination.
const RelNext = "next"

var (
	// ContextDisablePaging disable automatic pagination handling for a request.
	ContextDisablePaging = contextKey("paging")
)

// GetLink returns the URI of the first HTTP Link header with the requested
// relationship value, nil otherwise.
func GetLink(resp *http.Response, rel string) *url.URL {
	// Multiple Link headers may be present.
	for _, header := range resp.Header["Link"] {
		// Each Link header may have multiple links in it.
		if links, err := link.Parse(header); err == nil {
			for _, parsed := range links {
				if parsed.Rel == rel {
					// Link URIs may be relative, so resolve them.
					loc, _ := url.Parse(parsed.URI)
					return resp.Request.URL.ResolveReference(loc)
				}
			}
		}
	}

	return nil
}

// getAllPages fetches all pages in a paginated response and appends the
// results of each call to the parsed slice. Returns the total combined slice
// and the *last* response. If any request in the chain has an error, then
// processing is stopped and the error is returned.
func getAllPages(client *APIClient, parsed interface{}, resp *http.Response) (interface{}, *http.Response, GenericOpenAPIError) {
	var err error
	items := reflect.ValueOf(parsed)

	// Check for a `next` link relation header for pagination. If present, we
	// must follow it, append the data, and repeat until no more `next` link
	// is available (signaling the end of the list).
	for {
		if uri := GetLink(resp, RelNext); uri != nil {
			// Response gets set to the next response in the series.
			resp, err = client.cfg.HTTPClient.Get(uri.String())
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
			additional := reflect.New(items.Type())
			addIface := additional.Interface()
			err = client.decode(&addIface, data, resp.Header.Get("Content-Type"))
			if err != nil {
				return nil, resp, GenericOpenAPIError{
					body:  data,
					error: err.Error(),
				}
			}

			// Append the new items and then pop back up top to process the
			// headers in this latest response.
			items = reflect.AppendSlice(items, additional.Elem())
			continue
		}
		break
	}

	return items.Interface(), resp, GenericOpenAPIError{}
}

// HighLevelClient wraps an APIClient and http.Client, as well as providing
// custom high-level functionality.
type HighLevelClient struct {
	*APIClient
	*http.Client
}

// Decode a response body based on the content type.
func (c *HighLevelClient) Decode(v interface{}, body []byte, contentType string) error {
	return c.decode(v, body, contentType)
}

// GetModel performs an HTTP GET on the given URI and loads the model from the
// response if successful.
func (c *HighLevelClient) GetModel(uri string, model interface{}) (*http.Response, error) {
	req, _ := http.NewRequest(http.MethodGet, uri, nil)
	return c.DoModel(req, model)
}

// DoModel takes an HTTP request like http.Client.Do and makes the request. If
// successful, it also loads the model from the response.
func (c *HighLevelClient) DoModel(req *http.Request, model interface{}) (*http.Response, error) {
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unable to get model: status code %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = c.decode(model, body, resp.Header.Get("Content-Type"))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Do the request, logging the request/response if debugging is enabled.
func (c *HighLevelClient) Do(req *http.Request) (*http.Response, error) {
	if req.Header == nil {
		req.Header = http.Header{}
	}
	for k, v := range c.cfg.DefaultHeader {
		if req.Header.Get(k) == "" {
			req.Header.Set(k, v)
		}
	}

	if c.cfg.Debug {
		dump, err := httputil.DumpRequestOut(req, true)
		if err != nil {
			return nil, err
		}
		log.Printf("\n%s\n", string(dump))
	}

	res, err := c.Client.Do(req)

	if c.cfg.Debug {
		dump, err := httputil.DumpResponse(res, true)
		if err != nil {
			return res, err
		}
		log.Printf("\n%s\n", string(dump))
	}

	return res, err
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
