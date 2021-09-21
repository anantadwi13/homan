package haproxy

import (
	"github.com/anantadwi13/cli-whm/internal/homan/external/api/haproxy/client"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"net/http"
)

//NewHaproxyClient proxyHost => "localhost:8080" haproxyHost => "haproxy:5555"
func NewHaproxyClient(proxyHost, haproxyHost string) (*client.HAProxyDataPlaneAPI, runtime.ClientAuthInfoWriter) {
	auth := httptransport.BasicAuth("admin", "mypassword")
	clientTransport := httptransport.New(proxyHost, client.DefaultBasePath, client.DefaultSchemes)
	roundTripper := newCustomRoundTripper(clientTransport.Transport, haproxyHost)
	clientTransport.Transport = roundTripper
	return client.New(clientTransport, nil), auth
}

type customRoundTripper struct {
	defaultTransport http.RoundTripper
	haproxyHost      string
}

func newCustomRoundTripper(defaultTransport http.RoundTripper, haproxyHost string) http.RoundTripper {
	return &customRoundTripper{defaultTransport, haproxyHost}
}

func (t *customRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	request.Header.Set(http.CanonicalHeaderKey("x-target-host"), "http://"+t.haproxyHost+"/")
	response, err := t.defaultTransport.RoundTrip(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
