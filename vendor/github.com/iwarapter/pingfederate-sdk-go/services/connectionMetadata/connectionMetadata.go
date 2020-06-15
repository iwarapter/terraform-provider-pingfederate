package connectionMetadata

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type ConnectionMetadataService struct {
	Client *client.PfClient
}

// New creates a new instance of the ConnectionMetadataService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *ConnectionMetadataService {

	return &ConnectionMetadataService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//Export - Export a connection's SAML metadata that can be given to a partner.
//RequestType: POST
//Input: input *ExportInput
func (s *ConnectionMetadataService) Export(input *ExportInput) (result *string, resp *http.Response, err error) {
	path := "/connectionMetadata/export"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("POST", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//Convert - Convert a partner's SAML metadata into a JSON representation.
//RequestType: POST
//Input: input *ConvertInput
func (s *ConnectionMetadataService) Convert(input *ConvertInput) (result *models.ConvertMetadataResponse, resp *http.Response, err error) {
	path := "/connectionMetadata/convert"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("POST", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

type ConvertInput struct {
	Body models.ConvertMetadataRequest
}

type ExportInput struct {
	Body models.ExportMetadataRequest
}
