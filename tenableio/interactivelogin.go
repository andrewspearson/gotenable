package tenableio

import (
	"github.com/andrewspearson/gotenable/gotenableutils"
)

// InteractiveLogin takes user input to generate a session token
func (request Request) InteractiveLogin() {
	username := gotenableutils.UserInput(`Username: `, true)
	password := gotenableutils.UserInput(`Password: `, false)
	authCodePrompt := request.Login(username, password, ``)
	for authCodePrompt == true {
		authCode := gotenableutils.UserInput(`Authentication Code: `, true)
		authCodePrompt = request.Login(username, password, authCode)
	}
}
