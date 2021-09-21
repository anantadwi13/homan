package dns

import (
	"context"
	"net/http"
)

//NewDnsClient  proxy => "http://localhost:8080/" dnsHost => "system-dns:5555"
func NewDnsClient(proxy string, dnsHost string) (*ClientWithResponses, error) {
	return NewClientWithResponses(proxy, WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Add(http.CanonicalHeaderKey("x-target-host"), "http://"+dnsHost+"/")
		return nil
	}))
}
