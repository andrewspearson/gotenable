package tenablesc

import (
	"github.com/andrewspearson/gotenable"
	"github.com/andrewspearson/gotenable/gotenableutils"
	"strconv"
)

func (request Request) Login(username string, password string) {
	// Token request
	apiTokenReq := apiTokenReq{}
	apiTokenReq.Username = username
	apiTokenReq.Password = password
	resp := request.Post(`/token`, gotenableutils.StructToJSON(apiTokenReq))

	// Token response
	if resp.Resp.StatusCode == 200 {
		apiTokenResp := apiTokenResp{}
		gotenableutils.JSONToStruct(resp.Body, &apiTokenResp)
		request.HTTPHeaders.Add(`X-SecurityCenter`, strconv.Itoa(apiTokenResp.Response.Token))
	} else {
		gotenable.Log.Logger.Fatalln("[FATAL] Authentication failure.")
	}
}

func (request Request) Logout() {
	if request.HTTPHeaders.Get(`cookie`) != `` {
		request.Delete(`/token`)
	} else {
		gotenable.Log.Logger.Println(`[WARN] Session token does not exist. Skipping logout function.`)
	}
}
