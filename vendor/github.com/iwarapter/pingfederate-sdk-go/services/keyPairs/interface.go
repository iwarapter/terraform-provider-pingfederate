package keyPairs

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type KeyPairsAPI interface {
	GetKeyAlgorithms() (output *models.KeyAlgorithms, resp *http.Response, err error)
}
