package keyPairs

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type KeyPairsAPI interface {
	GetKeyAlgorithms() (result *models.KeyAlgorithms, resp *http.Response, err error)
}
