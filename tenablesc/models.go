package tenablesc

import (
	"github.com/andrewspearson/gotenable"
)

// Client models
type Config struct {
	gotenable.Config
	Vendor    string
	Product   string
	Version   string
	AccessKey string
	SecretKey string
}

type Request struct {
	gotenable.Request
}

type Response struct {
	gotenable.Response
}

// API models
type apiTokenReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type apiTokenResp struct {
	Response struct {
		Token int `json:"token"`
	} `json:"response"`
}
