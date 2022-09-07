package gotenable

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
	"net/url"
)

func HTTPTransportConfig(proxy string, cert string, insecureSkipVerify bool) *http.Transport {
	// Create a new HTTP transport
	transport := &http.Transport{}

	// Set proxy
	if proxy != `` {
		proxyURL, err := url.Parse(proxy)
		if err != nil {
			Log.Logger.Fatalln(`[FATAL]`, err)
		}
		transport.Proxy = http.ProxyURL(proxyURL)
	}

	// Set cert
	if cert != `` {
		caCert, err := ioutil.ReadFile(cert)
		if err != nil {
			Log.Logger.Fatalln(`[FATAL]`, err)
		}
		caCertPool, _ := x509.SystemCertPool()
		caCertPool.AppendCertsFromPEM(caCert)
		transport.TLSClientConfig = &tls.Config{RootCAs: caCertPool}
	}

	// Set TLS verification
	if insecureSkipVerify == true {
		transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: insecureSkipVerify}
		Log.Logger.Println(`[WARN] TLS verification has been disabled!`)
	}
	return transport
}
