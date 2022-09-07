package tenableio

import (
	"github.com/andrewspearson/gotenable"
	"regexp"
)

const baseURL = `https://cloud.tenable.com`

var validUsername = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

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
type apiSessionReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type apiSessionTOTPReq struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	TwoFactorToken string `json:"two_factor_token"`
}
type apiSessionResp struct {
	Error              string `json:"error"`
	PasswordResetToken string `json:"password_reset_token"`
	Token              string `json:"token"`
	TwoFactor          bool   `json:"two_factor"`
}
