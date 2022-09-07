package downloads

import (
	"fmt"
	"github.com/andrewspearson/gotenable"
	"github.com/andrewspearson/gotenable/gotenableutils"
	"github.com/hashicorp/go-retryablehttp"
	"io/ioutil"
	"net/http"
)

func (request Request) Get(path string) Response {
	req, err := retryablehttp.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", request.BaseURL, path), nil)
	gotenableutils.ErrFatal(err)
	return request.Do(req)
}

func (request Request) Do(req *retryablehttp.Request) (response Response) {
	// Set headers
	req.Header = request.HTTPHeaders

	// Log entire HTTP request and response for debugging
	request.HTTPClient.RequestLogHook = func(logger retryablehttp.Logger, request *http.Request, i int) {
		gotenableutils.RequestDump(request, true, `DEBUG`)
	}
	request.HTTPClient.ResponseLogHook = func(logger retryablehttp.Logger, response *http.Response) {
		gotenableutils.ResponseDump(response, true, `DEBUG`)
	}

	// Execute the request
	resp, err := request.HTTPClient.Do(req)
	gotenableutils.ErrFatal(err)
	//defer resp.Body.Close()

	// Return data
	response.Resp = resp
	response.Body, err = ioutil.ReadAll(resp.Body)
	gotenableutils.ErrFatal(err)

	// Close response body
	err = resp.Body.Close()
	if err != nil {
		gotenable.Log.Logger.Println(`[WARN] Failed to close response body. Memory leak has occurred.`)
	}

	// Return response
	return response
}
