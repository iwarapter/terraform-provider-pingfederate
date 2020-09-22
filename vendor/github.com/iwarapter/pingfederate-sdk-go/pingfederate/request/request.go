package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"reflect"
	"regexp"
	"runtime"
	"time"

	"github.com/google/uuid"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client/metadata"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/config"
)

const logReqMsg = `DEBUG: Request %s/%s: %s
---[ REQUEST ]--------------------------------------
%s
-----------------------------------------------------`
const logRespMsg = `DEBUG: Response %s/%s: %s
---[ RESPONSE ]--------------------------------------
%s
-----------------------------------------------------`

// A Request is the service request to be made.
type Request struct {
	Config     config.Config
	ClientInfo metadata.ClientInfo

	AttemptTime  time.Time
	Time         time.Time
	Operation    *Operation
	HTTPRequest  *http.Request
	HTTPResponse *http.Response
	Body         io.ReadSeeker
	BodyStart    int64 // offset from beginning of Body that the request body starts
	Params       interface{}
	Error        error
	Data         interface{}
	RequestID    string
}

// An Operation is the service API operation to be made.
type Operation struct {
	Name        string
	HTTPMethod  string
	HTTPPath    string
	QueryParams map[string]string
}

// New returns a new Request pointer for the service API
// operation and parameters.
//
// Params is any value of input parameters to be the request payload.
// Data is pointer value to an object which the request's response
// payload will be deserialized to.
func New(cfg config.Config, clientInfo metadata.ClientInfo, operation *Operation, params interface{}, data interface{}) *Request {

	method := operation.HTTPMethod

	var buf io.ReadWriter
	if params != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(params)
		if err != nil {
			return nil
		}
	}

	httpReq, _ := http.NewRequest(method, "", buf)

	httpReq.SetBasicAuth(*cfg.Username, *cfg.Password)
	httpReq.Header.Add("X-Xsrf-Header", "pingfederate")
	httpReq.Header.Add("User-Agent", fmt.Sprintf("%s/%s (%s; %s; %s)", pingfederate.SDKName, pingfederate.SDKVersion, runtime.Version(), runtime.GOOS, runtime.GOARCH))

	if params != nil {
		httpReq.Header.Set("Content-Type", "application/json")
	}

	var err error
	httpReq.URL, err = url.Parse(clientInfo.Endpoint + operation.HTTPPath)
	if err != nil {
		httpReq.URL = &url.URL{}
		err = fmt.Errorf("invalid endpoint uri %s", err)
	}

	q := httpReq.URL.Query()
	for k, v := range operation.QueryParams {
		if v != "" {
			q.Set(k, v)
		}
	}
	httpReq.URL.RawQuery = q.Encode()

	r := &Request{
		Config:     cfg,
		ClientInfo: clientInfo,

		AttemptTime: time.Now(),
		Time:        time.Now(),
		Operation:   operation,
		HTTPRequest: httpReq,
		Body:        nil,
		Params:      params,
		Error:       err,
		Data:        data,
		RequestID:   uuid.New().String(),
	}
	r.Body = bytes.NewReader([]byte{})

	return r
}

// Send will send the request, returning error if errors are encountered.
//
// Send will sign the request prior to sending. All Send Handlers will
// be executed in the order they were set.
//
// Canceling a request is non-deterministic. If a request has been canceled,
// then the transport will choose, randomly, one of the state channels during
// reads or getting the connection.
//
// readLoop() and getConn(req *Request, cm connectMethod)
// https://github.com/golang/go/blob/master/src/net/http/transport.go
//
// Send will not close the request.Request's body.
func (r *Request) Send() error {
	//for {
	if *r.Config.LogDebug {
		requestDump, err := httputil.DumpRequest(r.HTTPRequest, true)
		requestDumpStr := string(requestDump)
		if *r.Config.MaskAuthorization {
			r, _ := regexp.Compile(`Authorization: (.*)`)
			requestDumpStr = r.ReplaceAllString(requestDumpStr, "Authorization: ********")
		}

		if err != nil {
			fmt.Println(err)
		}
		log.Printf(logReqMsg, r.ClientInfo.ServiceName, r.Operation.Name, r.RequestID, requestDumpStr)
	}
	r.AttemptTime = time.Now()

	resp, err := r.Config.HTTPClient.Do(r.HTTPRequest)
	if err != nil {
		r.Error = err
		return err
	}
	r.HTTPResponse = resp

	if *r.Config.LogDebug {
		requestDump, err := httputil.DumpResponse(r.HTTPResponse, true)
		if err != nil {
			fmt.Println(err)
		}
		log.Printf(logRespMsg, r.ClientInfo.ServiceName, r.Operation.Name, r.RequestID, string(requestDump))
	}

	r.CheckResponse()
	if r.Error != nil {
		return r.Error
	}
	//r.Handlers.ValidateResponse.Run(r)
	//if r.Error != nil {
	//	r.Handlers.UnmarshalError.Run(r)
	//	err := r.Error
	//
	//	r.Handlers.Retry.Run(r)
	//	r.Handlers.AfterRetry.Run(r)
	//	if r.Error != nil {
	//		debugLogReqError(r, "Validate Response", false, err)
	//		return r.Error
	//	}
	//	debugLogReqError(r, "Validate Response", true, err)
	//	continue
	//}

	//r.Handlers.Unmarshal.Run(r)
	if r.DataFilled() {
		//v := reflect.Indirect(reflect.ValueOf(r.Data))
		defer r.HTTPResponse.Body.Close()

		if err := json.NewDecoder(r.HTTPResponse.Body).Decode(&r.Data); err != nil {
			if err == io.EOF {
				err = nil // ignore EOF errors caused by empty response body
			} else {
				r.Error = fmt.Errorf("failed to unmarshall response %s", err)
				return err
			}
		}
	}

	//	break
	//}

	return nil
}

// DataFilled returns true if the request's data for response deserialization
// target has been set and is a valid. False is returned if data is not
// set, or is invalid.
func (r *Request) DataFilled() bool {
	return r.Data != nil && reflect.ValueOf(r.Data).Elem().IsValid()
}

// CheckResponse checks the API response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range.
// API error responses are expected to have either no response
// body, or a JSON response body that maps to ErrorResponse. Any other
// response body will be silently ignored.
func (r *Request) CheckResponse() {
	if c := r.HTTPResponse.StatusCode; 200 <= c && c <= 299 {
		return
	}
	r.Data = nil
	errorResponse := models.ApiResult{}
	data, err := ioutil.ReadAll(r.HTTPResponse.Body)
	if err == nil && data != nil {
		err = json.Unmarshal(data, &errorResponse)
		if err != nil {
			r.Error = fmt.Errorf("unable to parse error response: %s", string(data))
			return
		}
	}

	r.Error = &PingFederateError{
		ApiResult: errorResponse,
	}
}

// PingFederateError occurs when PingFederate returns a non 2XX response
type PingFederateError struct {
	ApiResult models.ApiResult
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
