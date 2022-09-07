package downloads

import (
	"github.com/andrewspearson/gotenable"
)

const baseURL = `https://www.tenable.com/downloads/api/v2`

// Client models
type Config struct {
	gotenable.Config
	Vendor      string
	Product     string
	Version     string
	BearerToken string
}

type Request struct {
	gotenable.Request
}

type Response struct {
	gotenable.Response
}
