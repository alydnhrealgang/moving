package cmd

import "github.com/alydnhrealgang/moving/common/utils"

var (
	FlagApiUrl   = "https://192.168.31.49:8443/v1"
	FlagItemType = "article"
	FlagFields   = make([]string, 0)
	FlagProps    = make([]string, 0)
	FlagTags     = make([]string, 0)
	FlagFileName = utils.EmptyString
	FlagFieldMap = make(map[string]string)
	FlagPropMap  = make(map[string]string)
	FlagTagMap   = make(map[string]string)
)
