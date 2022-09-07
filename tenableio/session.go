package tenableio

import (
	"fmt"
	"github.com/andrewspearson/gotenable"
	"github.com/andrewspearson/gotenable/gotenableutils"
)

func (request Request) Login(username string, password string, authCode string) bool {
	// Session request
	var resp Response
	if validUsername.MatchString(username) == true {
		if authCode == `` {
			apiSessionReq := apiSessionReq{}
			apiSessionReq.Username = username
			apiSessionReq.Password = password
			resp = request.Post(`/session`, gotenableutils.StructToJSON(apiSessionReq))
		} else if gotenable.ValidTOTP.MatchString(authCode) == true {
			apiSessionTOTPReq := apiSessionTOTPReq{}
			apiSessionTOTPReq.Username = username
			apiSessionTOTPReq.Password = password
			apiSessionTOTPReq.TwoFactorToken = authCode
			resp = request.Post(`/session`, gotenableutils.StructToJSON(apiSessionTOTPReq))
		} else {
			gotenable.Log.Logger.Fatalln(`[FATAL] Invalid authentication code.`)
		}
	} else {
		gotenable.Log.Logger.Fatalln(`[FATAL] Invalid username.`)
	}

	// Session response
	apiSessionResp := apiSessionResp{}
	gotenableutils.JSONToStruct(resp.Body, &apiSessionResp)
	if resp.Resp.StatusCode == 200 {
		if apiSessionResp.TwoFactor == true {
			return true
		} else {
			request.HTTPHeaders.Add(`x-cookie`, fmt.Sprintf("token=%s", apiSessionResp.Token))
		}
	} else {
		if apiSessionResp.Error != `` {
			gotenable.Log.Logger.Fatalln("[FATAL]", apiSessionResp.Error)
		} else {
			gotenable.Log.Logger.Fatalln(`[FATAL] Authentication failure.`)
		}
	}
	return false
}

func (request Request) Logout() {
	if request.HTTPHeaders.Get(`x-cookie`) != `` {
		request.Delete(`/session`)
	} else {
		gotenable.Log.Logger.Println(`[WARN] Session token does not exist. Skipping logout function.`)
	}
}
