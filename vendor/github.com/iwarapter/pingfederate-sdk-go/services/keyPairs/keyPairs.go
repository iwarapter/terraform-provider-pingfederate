package keyPairs

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type KeyPairsService struct {
	Client *client.PfClient
}

// New creates a new instance of the KeyPairsService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *KeyPairsService {

	return &KeyPairsService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetKeyAlgorithms - Get list of the key algorithms supported for key pair generation.
//RequestType: GET
//Input:
func (s *KeyPairsService) GetKeyAlgorithms() (result *models.KeyAlgorithms, resp *http.Response, err error) {
	path := "/keyPairs/keyAlgorithms"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}
