package beansdk

import (
	"fmt"
	"github.com/anjude/bcore/pkg/httputil"
	"github.com/anjude/bean-sdk-go/beansdk/config"
	"github.com/anjude/bean-sdk-go/beansdk/request"
	"runtime"
	"strings"
)

var defaultUserAgent = fmt.Sprintf("BEANSDK/%s (%s; %s) Golang/%s", Version, runtime.GOOS, runtime.GOARCH, strings.Trim(runtime.Version(), "go"))

type Client struct {
	Credential *Credential
	httpClient *httputil.HTTPClient
	config     *config.Config
}

func NewClient(secretID, secretKey string) *Client {
	c := config.GetDefaultConfig()
	credentials := newCredential(secretID, secretKey)
	return &Client{
		Credential: credentials,
		httpClient: httputil.NewHTTPClient(c.BaseUrl, c.Timeout, defaultUserAgent),
	}
}

// Send sends a request to the server and returns the response.
func (c *Client) Send(req request.Request, resp interface{}) error {
	return c.httpClient.Request(req.GetMethod(), req.GetURL(), req.GetHeaders(), req.GetData(), resp)
}
