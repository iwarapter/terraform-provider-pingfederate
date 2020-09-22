package config

import (
	"crypto/tls"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate"
)

type Config struct {
	Username          *string
	Password          *string
	LogDebug          *bool
	MaskAuthorization *bool
	Endpoint          *string

	// The HTTP client to use when sending requests. Defaults to
	// `http.DefaultClient`.
	HTTPClient *http.Client
}

func NewConfig() *Config {
	//TODO this is epically shit dont leave it here, parse a ca bundle etc
	/* #nosec */
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	return &Config{
		MaskAuthorization: pingfederate.Bool(true),
		LogDebug:          pingfederate.Bool(false),
		HTTPClient:        http.DefaultClient,
	}
}

func (c *Config) WithPassword(password string) *Config {
	c.Password = pingfederate.String(password)
	return c
}

func (c *Config) WithUsername(username string) *Config {
	c.Username = pingfederate.String(username)
	return c
}

func (c *Config) WithEndpoint(endpoint string) *Config {
	c.Endpoint = pingfederate.String(endpoint)
	return c
}

func (c *Config) WithDebug(debug bool) *Config {
	c.LogDebug = pingfederate.Bool(debug)
	return c
}

func (c *Config) WithMaskAuthorization(debug bool) *Config {
	c.MaskAuthorization = pingfederate.Bool(debug)
	return c
}
