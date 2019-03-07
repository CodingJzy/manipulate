package maniphttp

import (
	"crypto/tls"
	"net/http"

	"go.aporeto.io/manipulate"
)

// ExtractCredentials extracts the username and password from the given manipulator.
// Note: the given manipulator must be an HTTP Manipulator or it will return an error.
func ExtractCredentials(manipulator manipulate.Manipulator) (string, string) {

	m, ok := manipulator.(*httpManipulator)
	if !ok {
		panic("You can only pass a HTTP Manipulator to ExtractCredentials")
	}

	m.renewLock.Lock()
	u, p := m.username, m.password
	m.renewLock.Unlock()

	return u, p
}

// ExtractEndpoint extracts the endpoint url from the given manipulator.
// Note: the given manipulator must be an HTTP Manipulator or it will return an error.
func ExtractEndpoint(manipulator manipulate.Manipulator) string {

	m, ok := manipulator.(*httpManipulator)
	if !ok {
		panic("You can only pass a HTTP Manipulator to ExtractEndpoint")
	}

	return m.url
}

// ExtractNamespace extracts the default namespace from the given manipulator.
// Note: the given manipulator must be an HTTP Manipulator or it will return an error.
func ExtractNamespace(manipulator manipulate.Manipulator) string {

	m, ok := manipulator.(*httpManipulator)
	if !ok {
		panic("You can only pass a HTTP Manipulator to ExtractNamespace")
	}

	return m.namespace
}

// ExtractTLSConfig returns a copy of the tls config from the given manipulator.
// Note: the given manipulator must be an HTTP Manipulator or it will return an error.
func ExtractTLSConfig(manipulator manipulate.Manipulator) *tls.Config {

	m, ok := manipulator.(*httpManipulator)
	if !ok {
		panic("You can only pass a HTTP Manipulator to ExtractTLSConfig")
	}

	return &tls.Config{
		RootCAs:            m.tlsConfig.RootCAs,
		Certificates:       m.tlsConfig.Certificates,
		InsecureSkipVerify: m.tlsConfig.InsecureSkipVerify,
	}
}

// SetGlobalHeaders sets the given headers to all requests that will be sent.
func SetGlobalHeaders(manipulator manipulate.Manipulator, headers http.Header) {

	m, ok := manipulator.(*httpManipulator)
	if !ok {
		panic("You can only pass a HTTP Manipulator to SetGlobalHeaders")
	}

	m.globalHeaders = headers
}
