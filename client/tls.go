package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
)

func NewTLSConfig() *tls.Config {

	certpool := x509.NewCertPool()
	pemCerts, err := ioutil.ReadFile("certs/ca.crt")
	if err == nil {
		certpool.AppendCertsFromPEM(pemCerts)
	}

	return &tls.Config{
		RootCAs: certpool,
	}
}
