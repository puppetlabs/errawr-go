package impl

import (
	"github.com/puppetlabs/errawr-go"
)

type HTTPErrorMetadata struct {
	ErrorStatus int `json:"error_status"`
}

func (hem HTTPErrorMetadata) Status() int {
	return hem.ErrorStatus
}

type ErrorMetadata struct {
	HTTPErrorMetadata *HTTPErrorMetadata `json:"http"`
}

func (em ErrorMetadata) HTTP() (errawr.HTTPMetadata, bool) {
	return em.HTTPErrorMetadata, em.HTTPErrorMetadata != nil
}
