package homand

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func TestHealthChecker(t *testing.T) {
	hc := NewHealthChecker()

	host := "127.0.0.1:32100"

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
	httpServer := http.Server{Addr: host, Handler: http.DefaultServeMux}
	go func() {
		err := httpServer.ListenAndServe()
		if err != http.ErrServerClosed {
			assert.Nil(t, err)
		}
	}()
	defer httpServer.Shutdown(context.TODO())

	time.Sleep(2 * time.Second)

	isAvailable, err := hc.IsAvailable(context.TODO(), HealthCheckHTTP, fmt.Sprintf("http://%v/ok", host))
	assert.Nil(t, err)
	assert.True(t, isAvailable)
	isAvailable, err = hc.IsAvailable(context.TODO(), HealthCheckHTTP, fmt.Sprintf("http://%v/ok2", host))
	assert.Nil(t, err)
	assert.True(t, isAvailable)
	isAvailable, err = hc.IsAvailable(context.TODO(), HealthCheckHTTP, fmt.Sprintf("http://%v/error", host))
	assert.Nil(t, err)
	assert.False(t, isAvailable)
	isAvailable, err = hc.IsAvailable(context.TODO(), HealthCheckHTTP, fmt.Sprintf("http://%v/error2", host))
	assert.Nil(t, err)
	assert.False(t, isAvailable)
	isAvailable, err = hc.IsAvailable(context.TODO(), HealthCheckTCP, host)
	assert.Nil(t, err)
	assert.True(t, isAvailable)
	isAvailable, err = hc.IsAvailable(context.TODO(), HealthCheckTCP, "127.0.0.1:32133")
	assert.Nil(t, err)
	assert.False(t, isAvailable)
}
