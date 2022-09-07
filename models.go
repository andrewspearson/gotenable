package gotenable

import (
	"github.com/hashicorp/go-retryablehttp"
	"github.com/hashicorp/logutils"
	"log"
	"net/http"
	"net/http/cookiejar"
	"regexp"
	"time"
)

// RegExes for input validation
var ValidTOTP = regexp.MustCompile(`^[0-9]{6}$`)
var ValidKeyHex20 = regexp.MustCompile(`^[a-f0-9]{20}$`)
var ValidKeyHex32 = regexp.MustCompile(`^[a-f0-9]{32}$`)
var ValidKeyHex64 = regexp.MustCompile(`^[a-f0-9]{64}$`)

// gotenable models for global settings
type LeveledLog struct {
	Logger      *log.Logger
	LevelFilter *logutils.LevelFilter
}

// Client models that will be inherited
type Config struct {
	BaseURL            string
	Proxy              string
	Cert               string
	InsecureSkipVerify bool
}

type Request struct {
	HTTPClient  *retryablehttp.Client
	HTTPHeaders http.Header
	BaseURL     string
}

type Response struct {
	Resp *http.Response
	Body []byte
}

// RequestDefaults sets default values on new Request objects
func RequestDefaults() Request {
	Request := Request{}
	Request.HTTPClient = retryablehttp.NewClient()
	Request.HTTPClient.Logger = Log.Logger
	Request.HTTPClient.HTTPClient.Timeout = 10 * time.Second
	jar, err := cookiejar.New(&cookiejar.Options{}) // https://pkg.go.dev/golang.org/x/net/publicsuffix
	if err != nil {
		Log.Logger.Fatalln(`[FATAL]`, err)
	}
	Request.HTTPClient.HTTPClient.Jar = jar
	Request.HTTPHeaders = http.Header{}
	return Request
}
