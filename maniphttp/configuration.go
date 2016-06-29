package maniphttp

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"net/http"

	"golang.org/x/crypto/pkcs12"
)

// TLSConfiguration holds various TLS configuration details
type TLSConfiguration struct {
	PkcsPath     string
	Password     string
	CAPath       string
	SkipInsecure bool
}

// NewTLSConfiguration returns a new TLSConfiguration
func NewTLSConfiguration(pkcs, password, ca string, skip bool) *TLSConfiguration {

	return &TLSConfiguration{
		PkcsPath:     pkcs,
		Password:     password,
		CAPath:       ca,
		SkipInsecure: skip,
	}
}

func (t *TLSConfiguration) makeHTTPClient() (*http.Client, error) {

	if t.PkcsPath == "" {
		return &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: t.SkipInsecure,
				},
			},
		}, nil
	}

	data, err := ioutil.ReadFile(t.PkcsPath)
	if err != nil {
		return nil, err
	}

	blocks, err := pkcs12.ToPEM(data, t.Password)
	if err != nil {
		return nil, err
	}

	var pemData []byte
	for _, b := range blocks {
		pemData = append(pemData, pem.EncodeToMemory(b)...)
	}

	cert, err := tls.X509KeyPair(pemData, pemData)
	if err != nil {
		return nil, err
	}

	caCert, err := ioutil.ReadFile(t.CAPath)

	if err != nil {
		return nil, err
	}

	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(caCert)

	TLSConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: t.SkipInsecure,
		RootCAs:            pool,
	}

	TLSConfig.BuildNameToCertificate()

	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: TLSConfig,
		},
	}, nil
}
