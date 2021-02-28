package pingfederate

import (
	"net"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func TestConfig_Client(t *testing.T) {
	server := httptest.NewUnstartedServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Send response to be tested
		rw.Header().Set("Content-Type", "application/json;charset=utf-8")
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte(`{"resultId":"invalid_credentials","message":"The credentials you provided were not recognized."}`))
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
			want: diag.Diagnostics{
				diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Invalid URL",
					Detail:   "Unable to parse base_url for client: parse \"not a url\": invalid URI for request",
				},
			},
		},
		{
			name:     "handle unresponsive server",
			username: "foo",
			password: "bar",
			baseURL:  "https://localhost:19999",
			want: diag.Diagnostics{
				diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Connection Error",
					Detail:   "Unable to connect to PingFederate: Unknown host/port",
				},
			},
		},
		{
			name:     "unauthenticated",
			username: "foo",
			password: "bar",
			baseURL:  server.URL,
			want: diag.Diagnostics{
				diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Connection Error",
					Detail:   "Unable to connect to PingFederate: The credentials you provided were not recognized.",
				},
			},
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
