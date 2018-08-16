package impl

import (
	"github.com/puppetlabs/errawr-go"
)

type HTTPErrorMetadata struct {
	ErrorStatus int
}

func (hem HTTPErrorMetadata) Status() int {
	return hem.ErrorStatus
}

type ErrorMetadata struct {
	HTTPErrorMetadata *HTTPErrorMetadata
}

func (em ErrorMetadata) HTTP() (errawr.HTTPMetadata, bool) {
	return em.HTTPErrorMetadata, em.HTTPErrorMetadata != nil
}
