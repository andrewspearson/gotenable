package gotenable

import (
	"github.com/hashicorp/logutils"
	"log"
	"os"
)

// Create leveled logging
var Log = LeveledLog{}

func init() {
	Log.Logger = &log.Logger{}
	Log.Logger.SetFlags(log.Ldate | log.Ltime | log.LUTC)
	Log.LevelFilter = &logutils.LevelFilter{}
	Log.LevelFilter.Levels = []logutils.LogLevel{"DEBUG", "INFO", "WARN", "ERR", "FATAL"}
	Log.LevelFilter.MinLevel = "WARN"
	Log.LevelFilter.Writer = os.Stderr
	Log.Logger.SetOutput(Log.LevelFilter)
}
