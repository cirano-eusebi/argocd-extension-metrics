package server

import (
	"net/http"

	"go.uber.org/zap"
)

// mimirRoundTripper is a custom http.RoundTripper that adds the X-Scope-OrgID header to requests.
type mimirRoundTripper struct {
	logger       *zap.SugaredLogger
	tenant       string
	roundTripper http.RoundTripper
}

// RoundTrip implements the http.RoundTripper interface and adds the X-Scope-OrgID header to the request.
func (m *mimirRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	m.logger.Infof("Adding X-Scope-OrgID %s header to request: %s", m.tenant, request.URL.String())
	clonedRequest := new(http.Request)
	*clonedRequest = *request
	clonedRequest.Header = make(http.Header)
	for k, s := range request.Header {
		clonedRequest.Header[k] = s
	}
	clonedRequest.Header.Set("X-Scope-OrgID", m.tenant)
	request = clonedRequest
	return m.roundTripper.RoundTrip(request)
}
