# GoTenable Documentation
## Package: gotenable
The gotenable package contains the following:

**Leveled Logging**  
GoTenable uses Hashicorp's [logutils](https://github.com/hashicorp/logutils) to provide leveled logging.
Valid log levels are: `DEBUG`, `INFO`, `WARN`, `ERR`, and `FATAL`. The default log level is `WARN`. `DEBUG` will log
full HTTP requests and responses to the console. Use caution when debugging because API keys and/or username/password
will be captured and logged.
```go
gotenable.Log.LevelFilter.MinLevel = "DEBUG"
```
The GoTenable leveled logger can be used in applications that have imported GoTenable.
Simply start your log message with `[DEBUG]`, `[INFO]`, `[WARN]`, `[ERR]`, or `[FATAL]` text.
```go
gotenable.Log.Logger.Println(`[WARN] Log message goes here.`)
```
\
**Retryable HTTP Client**  
GoTenable uses Hashicorp's [go-retryablehttp](https://github.com/hashicorp/go-retryablehttp) HTTP client.

## Package: tenableio
The tenableio package contains the following:

**Client configuration and creation**  
The client may be configured with a variety of settings. All of these settings have default values so none of them need
to be specified in your code.

All settings specified in client configuration.
```go
config := tenableio.Config{}
config.BaseURL = `https://cloud.tenable.com`
config.Proxy = `http://proxy.company.local:8080`
config.Cert = `~/Desktop/cacert.pem`
//config.InsecureSkipVerify = true
config.Vendor = `Foo company`
config.Product = `Foo app`
config.Version = `1.0.0`
config.AccessKey = `ACCESS_KEY`
config.SecretKey = `SECRET_KEY`

tio := tenableio.New(config)
```

No settings specified in client configuration.
```go
config := tenableio.Config{}

tio := tenableio.New(config)
```
\
**HTTP timeout**  
GoTenable sets the client's HTTP timeout to 10 seconds by default but this is configurable.
```go
tio.HTTPClient.HTTPClient.Timeout = 30 * time.Second
```
\
**Authentication**  
In addition to API key authentication, GoTenable also provides session based username/password/two-factor
authentication and session deletion. This is useful when developing programs intended to be run by end users that might
not have API keys. Please note, these functions leverage unsupported API endpoints which Tenable may change without
notice.
```go
tio.Login(`USERNAME`, `PASSWORD`, ``)
defer tio.Logout()
```
GoTenable also provides a function for interactive logins.
```go
tio.InteractiveLogin()
defer tio.Logout()
```
The InteractiveLogin() function will prompt the user for Username, Password, and a two-factor Authentication Code if
applicable. Passwords will not echo to the console. The user will only be prompted for an Authentication Code if
two-factor authentication is enabled.
```text
Username: apearson@tenable.com
Password: 
Authentication Code: 123456
```

\
**HTTP Methods**  
The client provides functions for HTTP GET, POST, PUT, PATCH, and DELETE. Additional functions will not likely be
implemented simply due to a lack of time to develop and maintain. However, the HTTP functions provided by GoTenable
will enable you to easily make any API calls listed on
[developer.tenable.com](https://developer.tenable.com/reference/navigate).

GET
```go
resp := tio.Get(`/scans`)
```
POST
```go
type createScanReqModel struct {
	UUID     string            `json:"uuid"`
	Settings struct {
		Name        string `json:"name"`
		Enabled     bool   `json:"enabled"`
		TextTargets string `json:"text_targets"`
	}                          `json:"settings"`
}

createScanReq := createScanReqModel{}
createScanReq.UUID = `bbd4f805-3966-d464-b2d1-0079eb89d69708c3a05ec2812bcf`
createScanReq.Settings.Name = `gotenable-documentation`
createScanReq.Settings.Enabled = false
createScanReq.Settings.TextTargets = `192.168.0.0/24`

resp = tio.Post(`/scans`, gotenableutils.StructToJSON(createScanReq))
```

## Package: tenablesc
The tenablesc package contains the following:

**Client configuration and creation**  
The client may be configured with a variety of settings. All of these settings, except for `BaseURL`, have default
values, so you do not need to specify them in your code.

All settings specified in client configuration.
```go
config := tenablesc.Config{}
config.BaseURL = `https://tenablesc.company.local`
config.Proxy = `http://proxy.company.local:8080`
config.Cert = `~/Desktop/cacert.pem`
//config.InsecureSkipVerify = true
config.Vendor = `Foo company`
config.Product = `Foo app`
config.Version = `1.0.0`
config.AccessKey = `ACCESS_KEY`
config.SecretKey = `SECRET_KEY`

tsc := tenableio.New(config)
```
Only required settings specified in client configuration.
```go
config := tenablesc.Config{}
config.BaseURL = `https://tenablesc.company.local`

tsc := tenablesc.New(config)
```
\
**HTTP timeout**  
GoTenable sets the client's HTTP timeout to 10 seconds by default but this is configurable.
```go
tio.HTTPClient.HTTPClient.Timeout = 30 * time.Second
```
\
**Authentication**  
In addition to API key authentication, GoTenable also provides session based username/password authentication and
session deletion. This is useful when developing programs intended to be run by end users that might not have API keys.
```go
tsc.Login(`USERNAME`, `PASSWORD`)
defer tsc.Logout()
```
GoTenable also provides a function for interactive logins.
```go
tsc.InteractiveLogin()
defer tsc.Logout()
```
The InteractiveLogin() function will prompt the user for Username and Password. Passwords will not echo to the console.
```text
Username: apearson
Password: 
```
\
**HTTP Methods**  
The client provides functions for HTTP GET, POST, PUT, PATCH, and DELETE. Additional functions will not likely be
implemented simply due to a lack of time to develop and maintain. However, the HTTP functions provided by GoTenable
will enable you to easily make any API calls listed on
[docs.tenable.com/tenablesc/api](https://docs.tenable.com/tenablesc/api/).

GET
```go
resp := tsc.Get(`/scan`)
```
POST
```go
type createScanReqModel struct {
	Type       string   `json:"type"`
	Name       string   `json:"name"`
	Policy     struct {
		ID int      `json:"id"`
	}                   `json:"policy"`
	Zone       struct {
		ID int      `json:"id"`
	}                   `json:"zone"`
	Repository struct {
		ID int      `json:"id"`
	}                   `json:"repository"`
	IPList     string   `json:"ipList"`
}

createScanReq := createScanReqModel{}
createScanReq.Name = `gotenable-documentation`
createScanReq.Policy.ID = 1000002
createScanReq.Zone.ID = 1
createScanReq.Repository.ID = 1
createScanReq.IPList = `192.168.0.0/24`

resp := tsc.Post(`/scan`, gotenableutils.StructToJSON(createScanReq))
```

## Package: downloads
The downloads package contains the following:

**Client configuration and creation**  
The client may be configured with a variety of settings. All of these settings, except for `BearerToken`, have default
values, so you do not need to specify them in your code.

All settings specified in client configuration.
```go
config := downloads.Config{}
config.BaseURL = `https://www.tenable.com/downloads/api/v2`
config.Proxy = `http://proxy.company.local:8080`
config.Cert = `~/Desktop/cacert.pem`
//config.InsecureSkipVerify = true
config.Vendor = `Foo company`
config.Product = `Foo app`
config.Version = `1.0.0`
config.BearerToken = `BEARER_TOKEN`

dl := downloads.New(config)
```
Only required settings specified in client configuration.
```go
config := downloads.Config{}
config.BearerToken = `BEARER_TOKEN`

dl := downloads.New(config)
```
\
**HTTP timeout**  
GoTenable sets the client's HTTP timeout to 10 seconds by default but this is configurable.
```go
dl.HTTPClient.HTTPClient.Timeout = 30 * time.Second
```
\
**HTTP Methods**  
The client provides a function for HTTP GET. Additional functions will not likely be implemented simply due to a lack of
time to develop and maintain. However, the HTTP functions provided by GoTenable will enable you to easily make any API
calls listed on [developer.tenable.com](https://developer.tenable.com/reference/navigate).

GET
```go
resp := dl.Get(`/pages`)
```

## Package: gotenableutils
The gotenableutils package provides functions to assist API application development work.
All of the following functions may be used in packages that have imported gotenable.

**PPrintJSON**  
Pretty prints JSON responses.
```go
gotenableutils.PPrintJSON(resp.Body)
```
\
**StructToJSON**  
Transforms struct data into JSON data (JSON marshal).
```go
body := gotenableutils.StructToJSON(createScanReq)
```
\
**JSONToStruct**  
Transforms JSON data into struct data (JSON unmarshal).
```go
createScanResp := createScanRespModel{}
gotenableutils.JSONToStruct(resp.Body, &createScanResp)
```
\
**ErrFatal**  
Shortcut for writing the tedious `if err != nil...`.
```go
policyReq.PolicyTemplate.ID, err = strconv.Atoi(policyTemplateID)
gotenableutils.ErrFatal(err)
```
\
**UserInput**  
Prompts and receives user input. User input does not have to echo.
```go
username := gotenableutils.UserInput(`Username: `, true)
password := gotenableutils.UserInput(`Password: `, false)
```