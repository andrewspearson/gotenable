package tenableio

import (
	"fmt"
	"github.com/andrewspearson/gotenable"
	"github.com/andrewspearson/gotenable/gotenableutils"
	"net/url"
)

func New(config Config) Request {
	// Configure request with default settings
	request := Request{}
	request.Request = gotenable.RequestDefaults()

	// Configure the base URL
	if config.BaseURL == `` {
		request.BaseURL = baseURL
	} else {
		baseURLParsed, err := url.Parse(config.BaseURL)
		gotenableutils.ErrFatal(err)
		request.BaseURL = fmt.Sprintf("%s://%s", baseURLParsed.Scheme, baseURLParsed.Host)
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
	if config.AccessKey != `` && config.SecretKey != `` {
		if gotenable.ValidKeyHex64.MatchString(config.AccessKey) == true && gotenable.ValidKeyHex64.MatchString(config.SecretKey) == true {
			request.HTTPHeaders.Add(`X-ApiKeys`, fmt.Sprintf("accessKey=%s;secretKey=%s", config.AccessKey, config.SecretKey))
		} else {
			gotenable.Log.Logger.Fatalln(`[FATAL] Invalid API keys.`)
		}
	}

	// Return configured request
	return request
}
