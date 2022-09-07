package gotenableutils

import (
	"github.com/andrewspearson/gotenable"
	"net/http"
	"net/http/httputil"
)

// ErrFatal is a convenience function for error handling
func ErrFatal(err error) {
	if err != nil {
		gotenable.Log.Logger.Fatalln(`[FATAL]`, err)
	}
}

// RequestDump logs the entire HTTP request for debugging
func RequestDump(req *http.Request, body bool, logLevel string) {
	reqDump, err := httputil.DumpRequestOut(req, body)
	if err != nil {
		gotenable.Log.Logger.Println(`[WARN] Failed to dump HTTP request. This could result in incomplete logging.`)
	}
	gotenable.Log.Logger.Printf("[%s] HTTP request dump:\n%s", logLevel, string(reqDump))
}

// ResponseDump logs entire HTTP response for debugging
func ResponseDump(resp *http.Response, body bool, logLevel string) {
	respDump, err := httputil.DumpResponse(resp, body)
	if err != nil {
		gotenable.Log.Logger.Println(`[WARN] Failed to dump HTTP response. This could result in incomplete logging.`)
	}
	gotenable.Log.Logger.Printf("[%s] HTTP response dump:\n%s", logLevel, string(respDump))
}
