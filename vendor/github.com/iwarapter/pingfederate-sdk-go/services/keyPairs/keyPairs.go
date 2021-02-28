package keyPairs

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client/metadata"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/config"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/request"
)

const (
	// ServiceName - The name of service.
	ServiceName = "KeyPairs"
)

type KeyPairsService struct {
	*client.PfClient
}

// New creates a new instance of the KeyPairsService client.
func New(cfg *config.Config) *KeyPairsService {

	return &KeyPairsService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a KeyPairs operation
func (c *KeyPairsService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetKeyAlgorithms - Get list of the key algorithms supported for key pair generation.
//RequestType: GET
//Input:
func (s *KeyPairsService) GetKeyAlgorithms() (output *models.KeyAlgorithms, resp *http.Response, err error) {
	return s.GetKeyAlgorithmsWithContext(context.Background())
}

//GetKeyAlgorithmsWithContext - Get list of the key algorithms supported for key pair generation.
//RequestType: GET
//Input: ctx context.Context,
func (s *KeyPairsService) GetKeyAlgorithmsWithContext(ctx context.Context) (output *models.KeyAlgorithms, resp *http.Response, err error) {
	path := "/keyPairs/keyAlgorithms"
	op := &request.Operation{
		Name:       "GetKeyAlgorithms",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.KeyAlgorithms{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}
