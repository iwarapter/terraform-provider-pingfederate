package framework

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TestConfig_Client(t *testing.T) {
	server := httptest.NewUnstartedServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Send response to be tested
		rw.Header().Set("Content-Type", "application/json;charset=utf-8")
		rw.WriteHeader(http.StatusUnauthorized)
		_, _ = rw.Write([]byte(`{"resultId":"invalid_credentials","message":"The credentials you provided were not recognized."}`))
	}))
	l, _ := net.Listen("tcp", ":0")
	server.Listener = l //for CI tests as host.docker.internal is window/macosx
	server.StartTLS()
	// Close the server when test finishes
	defer server.Close()

	tests := []struct {
		name     string
		username string
		password string
		baseURL  string
		want     diag.Diagnostics
	}{
		{
			name:     "handle malformed urls",
			username: "foo",
			password: "bar",
			baseURL:  "not a url",
			want:     diag.Diagnostics{diag.NewErrorDiagnostic("Invalid URL", "Unable to parse base_url for client: parse \"not a url\": invalid URI for request")},
		},
		{
			name:     "handle unresponsive server",
			username: "foo",
			password: "bar",
			baseURL:  "https://localhost:19999",
			want:     diag.Diagnostics{diag.NewErrorDiagnostic("Connection Error", "Unable to connect to PingFederate: Unknown host/port")},
		},
		{
			name:     "unauthenticated",
			username: "foo",
			password: "bar",
			baseURL:  server.URL,
			want:     diag.Diagnostics{diag.NewErrorDiagnostic("Connection Error", "Unable to connect to PingFederate: The credentials you provided were not recognized.")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &pfConfig{
				Username: tt.username,
				Password: tt.password,
				BaseURL:  tt.baseURL,
			}
			_, diags := c.Client()
			if !reflect.DeepEqual(diags, tt.want) {
				t.Errorf("Client() diags = %v, want %v", diags, tt.want)
			}
		})
	}
}

func TestIsVersion(t *testing.T) {
	cli := pfClient{}
	tests := []struct {
		version  string
		expected bool
	}{
		{"10.0", false},
		{"10.1", true},
		{"10.2", false},
		{"10.3", false},
		{"11.0", false},
		{"11.1", false},
		{"11.2", false},
		{"11.3", false},
		{"12.0", false},
		{"12.1", false},
		{"12.2", false},
		{"12.3", false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("we handle %s", tt.version), func(t *testing.T) {
			cli.apiVersion = tt.version
			var err error
			cli.major, cli.minor, err = parseVersion(tt.version)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, cli.IsVersion(10, 1))
		})
	}
}

func TestIsVersionLessThan(t *testing.T) {
	cli := pfClient{}
	tests := []struct {
		version  string
		expected bool
	}{
		{"10.0", true},
		{"10.1", true},
		{"10.2", true},
		{"10.3", true},
		{"11.0", true},
		{"11.1", true},
		{"11.2", false},
		{"11.3", false},
		{"12.0", false},
		{"12.1", false},
		{"12.2", false},
		{"12.3", false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("we handle %s", tt.version), func(t *testing.T) {
			cli.apiVersion = tt.version
			var err error
			cli.major, cli.minor, err = parseVersion(tt.version)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, cli.IsVersionLessEqThan(11, 1))
		})
	}
}

func TestIsVersionGreaterThan(t *testing.T) {
	cli := pfClient{}
	tests := []struct {
		version  string
		expected bool
	}{
		{"10.0", false},
		{"10.1", false},
		{"10.2", false},
		{"10.3", false},
		{"11.0", false},
		{"11.1", true},
		{"11.2", true},
		{"11.3", true},
		{"12.0", true},
		{"12.1", true},
		{"12.2", true},
		{"12.3", true},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("we handle %s", tt.version), func(t *testing.T) {
			cli.apiVersion = tt.version
			var err error
			cli.major, cli.minor, err = parseVersion(tt.version)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, cli.IsVersionGreaterEqThan(11, 1))
		})
	}
}
