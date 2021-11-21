package urlmanipulations

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"

	"github.com/rs/zerolog"
)

//go:generate mockgen --source=./url_manipulations.go --destination=./url_manipulations_mocks_test.go --package=urlmanipulations_test

type CallerURL interface {
	Call(url *url.URL) error
}

// how to model server with multiple http paths for test - not the best and rather reliable way but it works
func ExampleServer() {
	mux := chi.NewMux()
	mux.Get("/aaaa", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client aaaa")
	})
	mux.Get("/bbbb", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client bbbb")
	})
	// ts := http.Server{Handler: mux, Addr: ":8000"}
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	listenStartChan := make(chan struct{})
	go func() {
		l, err := net.Listen("tcp", ":8000")
		if err != nil {
			// handle error
		}

		listenStartChan <- struct{}{}

		if err := http.Serve(l, mux); err != nil {
			logger.Err(err).Msgf("serve error")
		}
	}()

	ctx := context.Background()
	ctx = logger.WithContext(ctx)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8000/vvvv", nil)

	httpClient := http.Client{}

	<-listenStartChan
	logger.Info().Msgf("server started and we can khow it")
	res, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	greeting, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", greeting)
	// Output: Hello, client
}
