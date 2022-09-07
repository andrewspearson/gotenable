package downloads

import (
	"fmt"
	"github.com/andrewspearson/gotenable"
	"strings"
)

func New(config Config) Request {
	// Configure request with default settings
	request := Request{}
	request.Request = gotenable.RequestDefaults()

	// Configure the base URL
	if config.BaseURL == `` {
		request.BaseURL = baseURL
	} else {
		request.BaseURL = strings.TrimSuffix(config.BaseURL, `/`)
	}

	// Configure the HTTP client's Transport settings
	if config.Proxy != `` || config.Cert != `` || config.InsecureSkipVerify == true {
		request.HTTPClient.HTTPClient.Transport = gotenable.HTTPTransportConfig(config.Proxy, config.Cert, config.InsecureSkipVerify)
	}

	// Add standard headers
	request.HTTPHeaders.Add(`Accept`, `application/json`)
	if config.Vendor == `` {
		config.Vendor = `Undefined`
	}
	if config.Product == `` {
		config.Product = `Undefined`
	}
	if config.Version == `` {
		config.Version = `Undefined`
	}
	request.HTTPHeaders.Add(`User-Agent`, fmt.Sprintf("Integration/1.0 (%s; %s; Build/%s)", config.Vendor, config.Product, config.Version))

	// Add authentication headers
	if config.BearerToken != `` {
		if gotenable.ValidKeyHex20.MatchString(config.BearerToken) == true {
			request.HTTPHeaders.Add(`Authorization`, fmt.Sprintf("Bearer %s", config.BearerToken))
		} else {
			gotenable.Log.Logger.Fatalln(`[FATAL] Invalid bearer token.`)
		}
	}

	// Return configured request
	return request
}
