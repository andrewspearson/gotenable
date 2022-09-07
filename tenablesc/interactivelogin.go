package tenablesc

import (
	"github.com/andrewspearson/gotenable/gotenableutils"
)

// InteractiveLogin takes user input to generate a session token
func (request Request) InteractiveLogin() {
	username := gotenableutils.UserInput(`Username: `, true)
	password := gotenableutils.UserInput(`Password: `, false)
	request.Login(username, password)
}
