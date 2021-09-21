package certman

import (
	"context"
	"net/http"
)

//NewCertmanClient proxy => "http://localhost:8080/" certmanHost => "system-certman:5555"
func NewCertmanClient(proxy string, certmanHost string) (*ClientWithResponses, error) {
	return NewClientWithResponses(proxy, WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Add(http.CanonicalHeaderKey("x-target-host"), "http://"+certmanHost+"/")
		return nil
	}))
}
