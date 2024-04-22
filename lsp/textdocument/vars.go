package textdocument

import "github.com/tliron/commonlog"

var (
	logger commonlog.Logger  = commonlog.GetLogger("")
	State  map[string]string = map[string]string{}
)
