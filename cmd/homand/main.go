package main

import (
	"context"
	"github.com/anantadwi13/cli-whm/internal/homand"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	osSign = make(chan os.Signal, 1)
)

func main() {
	signal.Notify(osSign, syscall.SIGINT, syscall.SIGTERM)
	e := echo.New()
	e.HideBanner = true
	hc := homand.NewHealthChecker()
	server := homand.NewServer(hc)

	go func() {
		homand.RegisterHandlers(e, server)
		err := e.Start(":80")
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	// todo add ssl synchronizer (certman -> haproxy)

	select {
	case <-osSign:
		err := e.Shutdown(context.Background())
		if err != nil {
			panic(err)
		}
	}
}
