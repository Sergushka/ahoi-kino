package main

import (
	"crypto/tls"
	"fmt"
	"github.com/sergushka/ahoi-kino/api/routers"
	"github.com/sergushka/ahoi-kino/db"
	"net/http"
	"time"
)

func Run() {
	Load()
	db.NewApp()

	//movie, _ := controllers.GetMovieByName("hitman")
	//
	//logger.Println(movie)

	listen()
	listenTls()
}

func listenTls() {
	router := routers.New()
	logger.Printf("Listening tls on %v", PORT)

	tlsConfig := &tls.Config{
		// Causes servers to use Go's default ciphersuite preferences,
		// which are tuned to avoid attacks. Does nothing on clients.
		PreferServerCipherSuites: true,
		// Only use curves which have assembly implementations
		CurvePreferences: []tls.CurveID{
			tls.CurveP256,
			//tls.X25519, // Go 1.8 only
		},
		MinVersion: tls.VersionTLS12,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			//tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305, // Go 1.8 only
			//tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,   // Go 1.8 only
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,

			// Best disabled, as they don't provide Forward Secrecy,
			// but might be necessary for some clients
			// tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			// tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
		}}

	tlsSrv := &http.Server{
		Addr:         fmt.Sprintf(":%d", PORT),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		TLSConfig:    tlsConfig,
		Handler:      router,
	}

	err := tlsSrv.ListenAndServeTLS(CERT_FILE, KEY_FILE)

	if err != nil {
		logger.Fatalf("server failed to start... %v", err)
	}
}

func listen() {
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Connection", "close")
			url := "https://" + req.Host + req.URL.String()
			http.Redirect(w, req, url, http.StatusMovedPermanently)
		}),
	}
	go func() { logger.Fatal(srv.ListenAndServe()) }()
}
