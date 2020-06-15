package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"runtime"
)

const logReqMsg = `DEBUG: Request %s Details:
---[ REQUEST ]--------------------------------------
%s
-----------------------------------------------------`
const logRespMsg = `DEBUG: Response %s Details:
---[ RESPONSE ]--------------------------------------
%s
-----------------------------------------------------`

type PfClient struct {
	Username   string
	Password   string
	BaseURL    *url.URL
	Context    string
	LogDebug   bool
	httpClient *http.Client
}

func NewClient(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *PfClient {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	c := &PfClient{httpClient: httpClient}
	c.Username = username
	c.Password = password
	c.BaseURL = baseUrl
	c.Context = context

	return c
}

func (c *PfClient) NewRequest(method string, path *url.URL, body interface{}) (*http.Request, error) {
	u := c.BaseURL.ResolveReference(path)
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.SetBasicAuth(c.Username, c.Password)
	req.Header.Add("X-Xsrf-Header", "pingfederate")
	req.Header.Add("User-Agent", fmt.Sprintf("%s/%s (%s; %s; %s)", SDKName, SDKVersion, runtime.Version(), runtime.GOOS, runtime.GOARCH))
	return req, nil
}

func (c *PfClient) Do(req *http.Request, v interface{}) (*http.Response, error) {
	if c.LogDebug {
		requestDump, err := httputil.DumpRequest(req, true)
		if err != nil {
			fmt.Println(err)
		}
		log.Printf(logReqMsg, "pingaccess-sdk-go", string(requestDump))
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if c.LogDebug {
		responseDump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			fmt.Println(err)
		}
		log.Printf(logRespMsg, "pingaccess-sdk-go", string(responseDump))
	}
	err = CheckResponse(resp)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
		} else {
			decErr := json.NewDecoder(resp.Body).Decode(v)
			if decErr == io.EOF {
				decErr = nil // ignore EOF errors caused by empty response body
			}
			if decErr != nil {
				err = decErr
			}
		}
	}
	return resp, err
}

// CheckResponse checks the API response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range.
// API error responses are expected to have either no response
// body, or a JSON response body that maps to ErrorResponse. Any other
// response body will be silently ignored.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	errorResponse := ApiResult{}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		err = json.Unmarshal(data, &errorResponse)
		if err != nil {
			return fmt.Errorf("unable to parse error response: %s", string(data))
		}
	}

	return &PingFederateError{
		ApiResult: errorResponse,
	}
}

// PingFederateError occurs when PingFederate returns a non 2XX response
type PingFederateError struct {
	ApiResult ApiResult
}

func (r *PingFederateError) Error() (message string) {
	if r.ApiResult.Message != nil {
		message = *r.ApiResult.Message
	}
	if r.ApiResult.ValidationErrors != nil && len(*r.ApiResult.ValidationErrors) > 0 {
		for _, v := range *r.ApiResult.ValidationErrors {
			if v.Message != nil {
				message = fmt.Sprintf("%s\n%s", message, *v.Message)
			}
		}
	}
	return
}

//ApiResult - Details on the result of the operation.
type ApiResult struct {
	DeveloperMessage *string             `json:"developerMessage,omitempty"`
	Message          *string             `json:"message,omitempty"`
	ResultId         *string             `json:"resultId,omitempty"`
	ValidationErrors *[]*ValidationError `json:"validationErrors,omitempty"`
}

//ValidationError - A data input validation error.
type ValidationError struct {
	DeveloperMessage *string `json:"developerMessage,omitempty"`
	ErrorId          *string `json:"errorId,omitempty"`
	FieldPath        *string `json:"fieldPath,omitempty"`
	Message          *string `json:"message,omitempty"`
}
