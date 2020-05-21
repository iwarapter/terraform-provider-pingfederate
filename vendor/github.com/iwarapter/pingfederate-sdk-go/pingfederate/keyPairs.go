package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
)

type KeyPairsService service

//GetKeyAlgorithms - Get list of the key algorithms supported for key pair generation.
//RequestType: GET
//Input:
func (s *KeyPairsService) GetKeyAlgorithms() (result *KeyAlgorithms, resp *http.Response, err error) {
	path := "/keyPairs/keyAlgorithms"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}
