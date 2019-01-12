package pingfederate

import (
	"crypto/tls"
	"net/http"
	"net/url"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate"
)

type Config struct {
	Username string
	Password string
	Context  string
	BaseURL  string
}

// Client configures and returns a fully initialized PAClient
func (c *Config) Client() (interface{}, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	url, _ := url.Parse(c.BaseURL)
	client := pingfederate.NewClient(c.Username, c.Password, url, c.Context, nil)

	return client, nil
}
