package server

import (
	"crypto/tls"
	"log"
	"net/http"
	"strings"

	"github.com/readytotouch/readytotouch/internal/env"

	"golang.org/x/crypto/acme/autocert"
	"golang.org/x/sync/errgroup"
)

func Run(handler http.Handler) {
	switch env.Environment() {
	case env.Production:
		runProduction(handler)
	case env.Development:
		runDevelopment(handler)
	}
}

// https://stackoverflow.com/questions/37321760/how-to-set-up-lets-encrypt-for-a-go-server-application
// https://stackoverflow.com/a/40494806/17655004
func runProduction(handler http.Handler) {
	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(strings.Split(env.Required("HOSTS"), ",")...),
		Cache:      autocert.DirCache(env.Required("TLS_CERTIFICATES_DIR")),
	}

	server := &http.Server{
		Addr:    ":https",
		Handler: handler,
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
			MinVersion:     tls.VersionTLS12, // improves cert reputation score at https://www.ssllabs.com/ssltest/
		},
	}

	var g errgroup.Group

	g.Go(func() error {
		return http.ListenAndServe(":http", certManager.HTTPHandler(nil))
	})

	g.Go(func() error {
		return server.ListenAndServeTLS("", "") // Key and cert are coming from Let's Encrypt
	})

	log.Fatal(g.Wait())
}

func runDevelopment(handler http.Handler) {
	log.Fatal(http.ListenAndServe(":http", handler))
}
