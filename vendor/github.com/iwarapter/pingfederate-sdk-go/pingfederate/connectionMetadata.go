package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
)

type ConnectionMetadataService service

//Export - Export a connection's SAML metadata that can be given to a partner.
//RequestType: POST
//Input: input *ExportInput
func (s *ConnectionMetadataService) Export(input *ExportInput) (result *string, resp *http.Response, err error) {
	path := "/connectionMetadata/export"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("POST", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//Convert - Convert a partner's SAML metadata into a JSON representation.
//RequestType: POST
//Input: input *ConvertInput
func (s *ConnectionMetadataService) Convert(input *ConvertInput) (result *ConvertMetadataResponse, resp *http.Response, err error) {
	path := "/connectionMetadata/convert"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("POST", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}
