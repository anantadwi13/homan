package homand

import (
	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestHealthChecker(t *testing.T) {
	hc := NewHealthChecker()

	http.HandleFunc("/ok", func(rw http.ResponseWriter, req *http.Request) {
		_, _ = rw.Write([]byte("ok"))
	})
	http.HandleFunc("/ok2", func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(299)
		_, _ = rw.Write([]byte("ok"))
	})
	http.HandleFunc("/error", func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(300)
		_, _ = rw.Write([]byte("error"))
	})
	http.HandleFunc("/error2", func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(500)
		_, _ = rw.Write([]byte("error"))
	})
	httpServer := http.Server{Addr: "localhost:32132", Handler: http.DefaultServeMux}
	go func() {
		err := httpServer.ListenAndServe()
		if err != http.ErrServerClosed {
			assert.Nil(t, err)
		}
	}()
	defer httpServer.Shutdown(context.TODO())

	isAvailable, err := hc.IsAvailable(context.TODO(), HealthCheckHTTP, "http://localhost:32132/ok")
	assert.Nil(t, err)
	assert.True(t, isAvailable)
	isAvailable, err = hc.IsAvailable(context.TODO(), HealthCheckHTTP, "http://localhost:32132/ok2")
	assert.Nil(t, err)
	assert.True(t, isAvailable)
	isAvailable, err = hc.IsAvailable(context.TODO(), HealthCheckHTTP, "http://localhost:32132/error")
	assert.Nil(t, err)
	assert.False(t, isAvailable)
	isAvailable, err = hc.IsAvailable(context.TODO(), HealthCheckHTTP, "http://localhost:32132/error2")
	assert.Nil(t, err)
	assert.False(t, isAvailable)
	isAvailable, err = hc.IsAvailable(context.TODO(), HealthCheckTCP, "localhost:32132")
	assert.Nil(t, err)
	assert.True(t, isAvailable)
	isAvailable, err = hc.IsAvailable(context.TODO(), HealthCheckTCP, "localhost:32133")
	assert.Nil(t, err)
	assert.False(t, isAvailable)
}
