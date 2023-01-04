package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"

	ihttp "github.com/yahfiilham/student-api/pkg/http"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, syscall.SIGTERM)

	go func() {
		oscall := <-ch
		log.Warn().Msgf("system call: %+v", oscall)
		cancel()
	}()

	err := run(ctx)
	if err != nil {
		log.Warn().Msgf("server run ")
	}
}

func run(ctx context.Context) error {
	h := ihttp.NewHandler()
	ihttp.Routes(h)

	srv := http.Server{
		Addr: fmt.Sprintf(":%d", 8080),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Msgf("listen: %+s\n", err)
		}
	}()

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// graceful shutdown
	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatal().Err(err).Msgf("server shutdown failed")
	}

	if err == http.ErrServerClosed {
		err = nil
	}

	return nil
}