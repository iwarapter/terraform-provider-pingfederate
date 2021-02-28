package keyPairs

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type KeyPairsAPI interface {
	GetKeyAlgorithms() (output *models.KeyAlgorithms, resp *http.Response, err error)
	GetKeyAlgorithmsWithContext(ctx context.Context) (output *models.KeyAlgorithms, resp *http.Response, err error)
}
