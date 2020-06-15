package pingfederate

import (
	"crypto/tls"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"net/http"
	"net/url"
	"os"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate"
)

type Config struct {
	Username string
	Password string
	Context  string
	BaseURL  string
}

// Client configures and returns a fully initialized PAClient
func (c *Config) Client() (interface{}, diag.Diagnostics) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	url, _ := url.Parse(c.BaseURL)
	client := pingfederate.NewClient(c.Username, c.Password, url, c.Context, nil)
	if os.Getenv("TF_LOG") == "DEBUG" || os.Getenv("TF_LOG") == "TRACE" {
		client.LogDebug = true
	}
	return client, nil
}
